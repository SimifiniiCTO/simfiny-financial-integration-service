# financial-integration-service
financial-integration-service for the simfinii platform

Specifications:

* Health checks (readiness and liveness)
* Graceful shutdown on interrupt signals
* File watcher for secrets and configmaps
* Instrumented with Prometheus
* Tracing with Istio and Jaeger
* Linkerd service profile
* Structured logging with zap 
* 12-factor app with viper
* Fault injection (random errors and latency)
* Swagger docs
* Helm and Kustomize installers
* End-to-End testing with Kubernetes Kind and Helm
* Kustomize testing with GitHub Actions and Open Policy Agent
* Multi-arch container image with Docker buildx and Github Actions
* CVE scanning with trivy

Web API:

* `GET /` prints runtime information
* `GET /version` prints financial-integration-service version and git commit hash 
* `GET /metrics` return HTTP requests duration and Go runtime metrics
* `GET /healthz` used by Kubernetes liveness probe
* `GET /readyz` used by Kubernetes readiness probe
* `POST /readyz/enable` signals the Kubernetes LB that this instance is ready to receive traffic
* `POST /readyz/disable` signals the Kubernetes LB to stop sending requests to this instance
* `GET /status/{code}` returns the status code
* `GET /panic` crashes the process with exit code 255
* `POST /echo` forwards the call to the backend service and echos the posted content 
* `GET /env` returns the environment variables as a JSON array
* `GET /headers` returns a JSON with the request HTTP headers
* `GET /delay/{seconds}` waits for the specified period
* `POST /token` issues a JWT token valid for one minute `JWT=$(curl -sd 'anon' financial-integration-service:9898/token | jq -r .token)`
* `GET /token/validate` validates the JWT token `curl -H "Authorization: Bearer $JWT" financial-integration-service:9898/token/validate`
* `GET /configs` returns a JSON with configmaps and/or secrets mounted in the `config` volume
* `POST/PUT /cache/{key}` saves the posted content to Redis
* `GET /cache/{key}` returns the content from Redis if the key exists
* `DELETE /cache/{key}` deletes the key from Redis if exists
* `POST /store` writes the posted content to disk at /data/hash and returns the SHA1 hash of the content
* `GET /store/{hash}` returns the content of the file /data/hash if exists
* `GET /ws/echo` echos content via websockets `podcli ws ws://localhost:9898/ws/echo`
* `GET /chunked/{seconds}` uses `transfer-encoding` type `chunked` to give a partial response and then waits for the specified period
* `GET /swagger.json` returns the API Swagger docs, used for Linkerd service profiling and Gloo routes discovery

gRPC API:

* `/grpc.health.v1.Health/Check` health checking

Web UI:

![financial-integration-service-ui](https://raw.githubusercontent.com/simfiniicto/podinfo/gh-pages/screens/podinfo-ui-v3.png)

To access the Swagger UI open `<podinfo-host>/swagger/index.html` in a browser.

### Guides

* [GitOps Progressive Deliver with Flagger, Helm v3 and Linkerd](https://helm.workshop.flagger.dev/intro/)
* [GitOps Progressive Deliver on EKS with Flagger and AppMesh](https://eks.handson.flagger.dev/prerequisites/)
* [Automated canary deployments with Flagger and Istio](https://medium.com/google-cloud/automated-canary-deployments-with-flagger-and-istio-ac747827f9d1)
* [Kubernetes autoscaling with Istio metrics](https://medium.com/google-cloud/kubernetes-autoscaling-with-istio-metrics-76442253a45a)
* [Autoscaling EKS on Fargate with custom metrics](https://aws.amazon.com/blogs/containers/autoscaling-eks-on-fargate-with-custom-metrics/)
* [Managing Helm releases the GitOps way](https://medium.com/google-cloud/managing-helm-releases-the-gitops-way-207a6ac6ff0e)
* [Securing EKS Ingress With Contour And Let’s Encrypt The GitOps Way](https://aws.amazon.com/blogs/containers/securing-eks-ingress-contour-lets-encrypt-gitops/)

### Install

Helm:

```bash
helm repo add financial-integration-service https://simfiniicto.github.io/financial-integration-service

helm upgrade --install --wait frontend \
--namespace test \
--set replicaCount=2 \
--set backend=http://backend-financial-integration-service:9898/echo \
financial-integration-service/financial-integration-service

helm test frontend

helm upgrade --install --wait backend \
--namespace test \
--set redis.enabled=true \
financial-integration-service/financial-integration-service
```

Kustomize:

```bash
kubectl apply -k github.com/SimifiniiCTO/simfinii/src/backend/services/financial-integration-service//kustomize
```

Docker:

```bash
docker run -dp 9898:9898 simfiniicto/financial-integration-service
```

### Continuous Delivery

In order to install financial-integration-service on a Kubernetes cluster and keep it up to date with the latest
release in an automated manner, you can use [Flux](https://fluxcd.io).

Install the Flux CLI on MacOS and Linux using Homebrew:

```sh
brew install fluxcd/tap/flux
```

Install the Flux controllers needed for Helm operations:

```sh
flux install \
--namespace=flux-system \
--network-policy=false \
--components=source-controller,helm-controller
```

Add financial-integration-service's Helm repository to your cluster and
configure Flux to check for new chart releases every ten minutes:

```sh
flux create source helm financial-integration-service \
--namespace=default \
--url=https://simfiniicto.github.io/financial-integration-service \
--interval=10m
```

Create a `financial-integration-service-values.yaml` file locally:

```sh
cat > financial-integration-service-values.yaml <<EOL
replicaCount: 2
resources:
  limits:
    memory: 256Mi
  requests:
    cpu: 100m
    memory: 64Mi
EOL
```

Create a Helm release for deploying financial-integration-service in the default namespace:

```sh
flux create helmrelease financial-integration-service \
--namespace=default \
--source=HelmRepository/financial-integration-service \
--release-name=financial-integration-service \
--chart=financial-integration-service \
--chart-version=">5.0.0" \
--values=financial-integration-service-values.yaml
```

Based on the above definition, Flux will upgrade the release automatically
when a new version of financial-integration-service is released. If the upgrade fails, Flux
can [rollback](https://toolkit.fluxcd.io/components/helm/helmreleases/#configuring-failure-remediation)
to the previous working version.

You can check what version is currently deployed with:

```sh
flux get helmreleases -n default
```

To delete financial-integration-service's Helm repository and release from your cluster run:

```sh
flux -n default delete source helm financial-integration-service
flux -n default delete helmrelease financial-integration-service
```

If you wish to manage the lifecycle of your applications in a **GitOps** manner, check out
this [workflow example](https://github.com/fluxcd/flux2-kustomize-helm-example)
for multi-env deployments with Flux, Kustomize and Helm.

## Running Locally
---
To start service via docker-compose, run the below command and reference the below three urls.

```bash
# to start services and dependencies locally
make compose-up-d

# to bring down services and dependencies
make compose-down
```

To run service on local kubernetes cluster, run the below command

```bash
# to start services and dependencies in minikube
make kube-deploy

# to bring down cluster
minikube delete
```

__For more detailed setup instructions, please reference the docs folder__

---
To start service, run the below command and reference the below three urls.

```bash
# to start services and dependencies locally
make compose-up-d

# to bring down services and dependencies
make compose-down
```

RK TV: http://localhost:9196/rk/v1/tv

Swagger UI: http://localhost:9196/sw

Prometheus client: http://localhost:9196/metrics

To run service on local kubernetes cluster, run the below command

```bash
# to start services and dependencies in minikube
make kube-deploy

# to bring down cluster
minikube delete
```

ref: https://github.com/duyquang6/wager-management-be