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
	
compose-up:
	docker-compose \
		-f compose/docker-compose.yaml \
		-f compose/docker-compose-ngrok.yaml up 

compose-up-d:
	docker-compose \
		-f compose/docker-compose.yaml \
		-f compose/docker-compose-newrelic.yaml \
		-f compose/docker-compose-ngrok.yaml up --detach

compose-down:
	docker-compose \
		-f compose/docker-compose.yaml \
		-f compose/docker-compose-newrelic.yaml \
		-f compose/docker-compose-ngrok.yaml down

.PHONY: test
test: compose-up-d
	echo "waiting for services to be ready to accept connections"
	sleep 60
	go test ./... -coverprofile cover.out
	docker ps -a 
	make compose-down

generate:
	docker run -v $$(pwd):/src -w /src --rm bufbuild/buf:$(BUF_VERSION) generate

lint:
	docker run -v $$(pwd):/src -w /src --rm bufbuild/buf:$(BUF_VERSION) lint
	docker run -v $$(pwd):/src -w /src --rm bufbuild/buf:$(BUF_VERSION) breaking --against 'https://github.com/johanbrandhorst/grpc-gateway-boilerplate.git#branch=master'

deploy:
	./deploy/deploy.sh