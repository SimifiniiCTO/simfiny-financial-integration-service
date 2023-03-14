# Makefile for releasing podinfo
#
# The release version is controlled from pkg/version
#--gorm_out=engine=postgres,enums=string,gateway:./protos/service \


TAG?=latest
NAME:=financial-integration-service
DOCKER_REPOSITORY:=feelguuds
DOCKER_IMAGE_NAME:=$(DOCKER_REPOSITORY)/$(NAME)
GIT_COMMIT:=$(shell git describe --dirty --always)
VERSION:=$(shell grep 'VERSION' pkg/version/version.go | awk '{ print $$4 }' | tr -d '"')
EXTRA_RUN_ARGS?=
BUF_VERSION:=1.0.0-rc9
GOCMD=go
GOTEST=$(GOCMD) test
GOVET=$(GOCMD) vet
EXPORT_RESULT ?= false
DC=docker-compose -f compose/docker-compose.yaml -f compose/docker-compose-newrelic.yaml

.PHONY: help
.DEFAULT_GOAL := help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

run:
	go run -ldflags "-s -w -X github.com/SimifiniiCTO/simfinii/pkg/version.REVISION=$(GIT_COMMIT)" cmd/podinfo/* \
	--level=debug --grpc-port=9999 --backend-url=https://httpbin.org/status/401 --backend-url=https://httpbin.org/status/500 \
	--ui-logo=https://raw.githubusercontent.com/SimifiniiCTO/simfinii/gh-pages/cuddle_clap.gif $(EXTRA_RUN_ARGS)

build:
	GIT_COMMIT=$$(git rev-list -1 HEAD) && CGO_ENABLED=0 go build  -ldflags "-s -w -X github.com/SimifiniiCTO/simfinii/pkg/version.REVISION=$(GIT_COMMIT)" -a -o ./bin/podinfo ./cmd/podinfo/*
	GIT_COMMIT=$$(git rev-list -1 HEAD) && CGO_ENABLED=0 go build  -ldflags "-s -w -X github.com/SimifiniiCTO/simfinii/pkg/version.REVISION=$(GIT_COMMIT)" -a -o ./bin/podcli ./cmd/podcli/*

fmt:
	gofmt -l -s -w ./
	goimports -l -w ./

build-charts:
	helm lint charts/*
	helm package charts/*

build-container:
	docker build -t $(DOCKER_IMAGE_NAME):$(VERSION) .

build-xx:
	docker buildx build \
	--platform=linux/amd64 \
	-t $(DOCKER_IMAGE_NAME):$(VERSION) \
	--load \
	-f Dockerfile.xx .

build-base:
	docker build -f Dockerfile.base -t $(DOCKER_REPOSITORY)/financial-integration-service-base:latest .

push-base: build-base
	docker push $(DOCKER_REPOSITORY)/financial-integration-service-base:latest

test-container:
	@docker rm -f podinfo || true
	@docker run -dp 9898:9898 --name=podinfo $(DOCKER_IMAGE_NAME):$(VERSION)
	@docker ps
	@TOKEN=$$(curl -sd 'test' localhost:9898/token | jq -r .token) && \
	curl -sH "Authorization: Bearer $${TOKEN}" localhost:9898/token/validate | grep test

.PHONY: integration-test
integration-test: kube-deploy
	./test/e2e.sh

push-container:
	docker tag $(DOCKER_IMAGE_NAME):$(VERSION) $(DOCKER_IMAGE_NAME):latest
	docker push $(DOCKER_IMAGE_NAME):$(VERSION)
	docker push $(DOCKER_IMAGE_NAME):latest
	docker tag $(DOCKER_IMAGE_NAME):$(VERSION) quay.io/$(DOCKER_IMAGE_NAME):$(VERSION)
	docker tag $(DOCKER_IMAGE_NAME):$(VERSION) quay.io/$(DOCKER_IMAGE_NAME):latest
	docker push quay.io/$(DOCKER_IMAGE_NAME):$(VERSION)
	docker push quay.io/$(DOCKER_IMAGE_NAME):latest

version-set:
	@next="$(TAG)" && \
	current="$(VERSION)" && \
	sed -i '' "s/$$current/$$next/g" pkg/version/version.go && \
	sed -i '' "s/tag: $$current/tag: $$next/g" charts/podinfo/values.yaml && \
	sed -i '' "s/tag: $$current/tag: $$next/g" charts/podinfo/values-prod.yaml && \
	sed -i '' "s/appVersion: $$current/appVersion: $$next/g" charts/podinfo/Chart.yaml && \
	sed -i '' "s/version: $$current/version: $$next/g" charts/podinfo/Chart.yaml && \
	sed -i '' "s/podinfo:$$current/podinfo:$$next/g" kustomize/deployment.yaml && \
	sed -i '' "s/podinfo:$$current/podinfo:$$next/g" deploy/webapp/frontend/deployment.yaml && \
	sed -i '' "s/podinfo:$$current/podinfo:$$next/g" deploy/webapp/backend/deployment.yaml && \
	sed -i '' "s/podinfo:$$current/podinfo:$$next/g" deploy/bases/frontend/deployment.yaml && \
	sed -i '' "s/podinfo:$$current/podinfo:$$next/g" deploy/bases/backend/deployment.yaml && \
	echo "Version $$next set in code, deployment, chart and kustomize"

release:
	git tag $(VERSION)
	git push origin $(VERSION)

swagger:
	go get github.com/swaggo/swag/cmd/swag
	cd pkg/api && $$(go env GOPATH)/bin/swag init -g server.go

get-deps:
	@echo "downloading protoc-gen tool"
	go get google.golang.org/grpc/cmd/protoc-gen-go-grpc

gen:
	buf generate
	cp proto/fis_service.swagger.json proto/swagger/swagger.json

.PHONY: start-skaffold
dev:
	minikube start
	skaffold config set --global local-cluster true
	eval $(minikube docker-env)
	skaffold dev

.PHONY: start-minikube-deployment
start-minikube-deployment:
	./scripts/local-build.sh
	./scripts/local-deploy.sh

.PHONY: stop-minikube-deployment
stop-minikube-deployment:
	./scripts/local-nuke.sh

stop: ## Stop all Docker Containers run in Compose
	$(DC) stop

clean: stop ## Clean all Docker Containers and Volumes
	$(DC) down --rmi local --remove-orphans -v
	$(DC) rm -f -v

build: clean ## Rebuild the Docker Image for use by Compose
	$(DC) build

start: stop ## Run the Application as a docker compose workflow
	$(DC) up

run-background: stop ## Run the Application
	$(DC) up --detach

.PHONY: test
test: run-background
	echo "waiting for services to be ready to accept connections"
	sleep 60
	go test ./... -coverprofile cover.out
	docker ps -a 
	make stop

generate:
	docker run -v $$(pwd):/src -w /src --rm bufbuild/buf:$(BUF_VERSION) generate

lint:
	docker run -v $$(pwd):/src -w /src --rm bufbuild/buf:$(BUF_VERSION) lint
	docker run -v $$(pwd):/src -w /src --rm bufbuild/buf:$(BUF_VERSION) breaking --against 'https://github.com/johanbrandhorst/grpc-gateway-boilerplate.git#branch=master'

deploy:
	./deploy/deploy.sh

stop.cluster: ## Delete kind cluster
	kind delete cluster 

start.cluster: ## Starts a local kind cluster
	kind create cluster || true

start.linkerd:  ## Setup LinkerD in local kubernetes cluster
	linkerd check --pre
	linkerd install | kubectl apply -f -
	linkerd viz install | kubectl apply -f -
	linkerd check
	linkerd viz dashboard &

.PHONY: tilt
start.k8s: stop.cluster start.cluster start.linkerd ## Start local tilt dev workflow
	tilt up

test.unit: run-background ## Run the tests of the project
ifeq ($(EXPORT_RESULT), true)
	GO111MODULE=off go get -u github.com/jstemmer/go-junit-report
	$(eval OUTPUT_OPTIONS = | tee /dev/tty | go-junit-report -set-exit-code > junit-report.xml)
endif
	$(GOTEST) -v -race ./... $(OUTPUT_OPTIONS)

test.integration:
	$(ENV_INTEGRATION_TEST) $(GOTEST) -tags=integration ./internal/integration -v -count=1

coverage: ## Generate coverage report
	$(GOTEST) -cover -covermode=count -coverprofile=profile.cov ./...
	$(GOCMD) tool cover -func profile.cov
ifeq ($(EXPORT_RESULT), true)
	GO111MODULE=off go get -u github.com/AlekSi/gocov-xml
	GO111MODULE=off go get -u github.com/axw/gocov/gocov
	gocov convert profile.cov | gocov-xml > coverage.xml
endif

clean.files: ## Remove
	rm -f ./junit-report.xml checkstyle-report.xml ./coverage.xml ./profile.cov

kill.docker.desktop:
	pkill -SIGHUP -f /Applications/Docker.app 'docker serve' 

start.docker.desktop:
	./integration-test/docker-desktop.sh


gen:
	cd api && make && cd ..