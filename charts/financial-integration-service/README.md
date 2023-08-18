# financial-integration-service

Financial Integration Service Helm chart for Kubernetes

## TL;DR
```console
$ helm repo add simfinii https://github.com/SimifiniiCTO/simfinii
$ helm install my-release simfinii/financial-integration-service
```

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| yoan yomba | <yoan@simfinii.com> |  |

![Version: 0.1.1](https://img.shields.io/badge/Version-0.1.1-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.16.0](https://img.shields.io/badge/AppVersion-1.16.0-informational?style=flat-square) ![Version: 0.1.1](https://img.shields.io/badge/Version-0.1.1-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.16.0](https://img.shields.io/badge/AppVersion-1.16.0-informational?style=flat-square)

```
          ,-.
 ,     ,-.   ,-.
/ \   (   )-(   )
\ |  ,.>-(   )-<
 \|,' (   )-(   )
  Y ___`-'   `-'
  |/__/   `-'
  |
  |
  |    SIMFINY PLATFORM - This platform is attempting to revolutionize the finacial gamification as well
  |                        as the social experience users consumers by both enabling them to better reach their
  |                        target goals as well as allowing such consumers to share ideas and progress in communities
__|_____________
```

## Service Background

Simfinii provides a medium by which numerous stakeholders can not only track the health of their finances, but
 leverage their peers and our various offerings to achieve all conceivable financial goals they may withhold.
 Given the nature of the platform we are building, properly architecting the flow of financial data across
 the Simfinii ecosystem is a component critical to providing the value we believe our users need. The financial
 integration service plays an instrumental role in this realm from the context of the backend.

Through the financial integration service, users are able to obtain a holistic view of their financial health
spanning investments, liabilities, and past transactions. This will serve as the core service driving all
features requiring user financial data.

## Introduction
This chart bootstraps the headless authentication service deployment on a Kubernetes cluster using the Helm package manager.

This chart can be used with Kubeapps for deployment and management of Helm Charts in clusters. This Helm chart has been tested on top of a Kubernetes Runtime.

## Requirements

Kubernetes: `>=1.19.0-0`

* Kubernetes 1.19+
* Helm 3.2.0+
* PV provisioner support in the underlying infrastructure

## Installing the Chart

To install the chart with the release name `my-release`:

```console
$ helm repo add simfinii https://github.com/SimifiniiCTO/simfinii
$ helm install my-release simfinii/financial-integration-service
```

These commands deploy the financial integration service on the Kubernetes cluster in the default configuration. The Parameters section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall/delete the `my-release` deployment:

```console
$ helm delete my-release
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| affinity | object | `{}` |  |
| autoscaling | object | `{"behavior":{"scaleDown":{"percent":100,"periodSeconds":15,"stabilizationWindowSeconds":300},"scaleUp":{"percent":100,"periodSeconds":15,"stabilizationWindowSeconds":0}},"cpu":90,"enabled":true,"maxReplicas":10,"memory":90,"minReplicas":1}` | autoscaling configuration for pod traffic scalibility |
| autoscaling.behavior | object | `{"scaleDown":{"percent":100,"periodSeconds":15,"stabilizationWindowSeconds":300},"scaleUp":{"percent":100,"periodSeconds":15,"stabilizationWindowSeconds":0}}` | Define the scaling policies and behavior for scaling up and scaling down. This is particularly useful to avoid rapid fluctuations in pod counts. |
| autoscaling.behavior.scaleDown | object | `{"percent":100,"periodSeconds":15,"stabilizationWindowSeconds":300}` | Scale down configuration |
| autoscaling.behavior.scaleDown.percent | int | `100` | The percentage by which the autoscaler can increase or decrease the current replica count. |
| autoscaling.behavior.scaleDown.periodSeconds | int | `15` | The time window to consider when applying the percent scaling policy. |
| autoscaling.behavior.scaleDown.stabilizationWindowSeconds | int | `300` | The number of seconds for which past recommendations should be considered when scaling up or down. |
| autoscaling.behavior.scaleUp | object | `{"percent":100,"periodSeconds":15,"stabilizationWindowSeconds":0}` | Scale up configuration |
| autoscaling.behavior.scaleUp.percent | int | `100` | The percentage by which the autoscaler can increase or decrease the current replica count. |
| autoscaling.behavior.scaleUp.periodSeconds | int | `15` | The time window to consider when applying the percent scaling policy. |
| autoscaling.behavior.scaleUp.stabilizationWindowSeconds | int | `0` | The number of seconds for which past recommendations should be considered when scaling up or down. |
| autoscaling.cpu | int | `90` | Target average utilizations for CPU and Memory, respectively. These are thresholds that, when surpassed, will trigger the autoscaler to increase the pod count, or decrease if the utilization is below the threshold for a certain time. |
| autoscaling.enabled | bool | `true` | Whether the autoscaler should be applied or not. |
| autoscaling.maxReplicas | int | `10` | Maximum replicas for the deployment |
| autoscaling.memory | int | `90` | Target average utilization for Memory in percentage |
| autoscaling.minReplicas | int | `1` | Minimum replicas for the deployment |
| aws.accessKeyID | string | `"AKIA5HFOAJRN7YDEYPST"` |  |
| aws.kmsID | string | `"mrk-e44f269bc0034feb95ede34154c3cfe4"` |  |
| aws.region | string | `"us-east-2"` |  |
| aws.secretAccessKey | string | `"c4XOO/7RLxjonKrmZIvdIOef8TiG4C/fnOgm3JsL"` |  |
| backend | string | `nil` |  |
| backends | list | `[]` |  |
| batchJobs.actionableInsights.enabled | bool | `true` |  |
| batchJobs.actionableInsights.interval | string | `"@weekly"` |  |
| batchJobs.recurringTransactions.enabled | bool | `true` |  |
| batchJobs.recurringTransactions.interval | string | `"@every 3h"` |  |
| batchJobs.syncAllAccounts.enabled | bool | `true` |  |
| batchJobs.syncAllAccounts.interval | string | `"@every 1h"` |  |
| cache | string | `""` |  |
| cacheTLSEnabled | bool | `true` |  |
| cacheTTLInSeconds | int | `3600` |  |
| certificate | object | `{"create":false,"dnsNames":["financial-integration-service"],"duration":"24h","ipAddresses":{},"issuerRef":{"kind":"ClusterIssuer","name":"self-signed"},"key":{"algorithm":"rsa","rotationPolicy":"Never","size":"2048"},"namespace":"","renewBefore":"12h","usages":["server auth","client auth"]}` | Create a Certificate Manager certificate (cert-manager required). |
| certificate.dnsNames | list | `["financial-integration-service"]` | The hostname and any subject alternative names for the certificate. |
| certificate.duration | string | `"24h"` | The validity duration of the certificate. |
| certificate.ipAddresses | object | `{}` | IP addresses to associate with the certificate (usually for internal services). |
| certificate.issuerRef | object | `{"kind":"ClusterIssuer","name":"self-signed"}` | Reference to the Issuer or ClusterIssuer resource that will issue the certificate. |
| certificate.issuerRef.kind | string | `"ClusterIssuer"` | Type of the issuer resource. Can be either 'Issuer' (namespace-scoped) or 'ClusterIssuer' (cluster-scoped). |
| certificate.issuerRef.name | string | `"self-signed"` | Name of the issuer resource. |
| certificate.key | object | `{"algorithm":"rsa","rotationPolicy":"Never","size":"2048"}` | Configuration related to the private key of the certificate. |
| certificate.key.algorithm | string | `"rsa"` | Cryptographic algorithm used for the key. |
| certificate.key.rotationPolicy | string | `"Never"` | Policy that determines when to rotate the private key. 'Never' means no automatic rotation. |
| certificate.key.size | string | `"2048"` | Size (in bits) of the key. Common sizes are 2048, 3072 or 4096 for RSA. |
| certificate.namespace | string | `""` | The namespace in which the certificate should be created. Leave empty for default namespace. |
| certificate.renewBefore | string | `"12h"` | Duration before the certificate expiration to renew it. |
| certificate.usages | list | `["server auth","client auth"]` | Intended usage of the certificate, determining its key usages and extended key usages. |
| clickhouse.host | string | `"l77zn1po0n.us-east-1.aws.clickhouse.cloud"` |  |
| clickhouse.maxConnectionAttempts | int | `4` |  |
| clickhouse.maxConnectionRetries | int | `4` |  |
| clickhouse.maxConnectionRetrySleepInterval | string | `"100ms"` |  |
| clickhouse.maxConnectionRetryTimeout | string | `"500ms"` |  |
| clickhouse.maxQueryTimeout | string | `"500ms"` |  |
| clickhouse.password | string | `"8OnfmVp~U6xnZ"` |  |
| clickhouse.port | int | `8443` |  |
| clickhouse.uri | string | `"https://default:8OnfmVp~U6xnZ@l77zn1po0n.us-east-1.aws.clickhouse.cloud:8443/default?secure=true&skip_verify=true"` |  |
| clickhouse.user | string | `"default"` |  |
| database.connectionMaxLifetime | string | `"10h"` |  |
| database.host | string | `"financial-integration-service-db"` |  |
| database.maxConnectionAttempts | int | `4` |  |
| database.maxConnectionRetries | int | `4` |  |
| database.maxConnectionRetrySleepInterval | string | `"100ms"` |  |
| database.maxConnectionRetryTimeout | string | `"500ms"` |  |
| database.maxIdleConnections | int | `10` |  |
| database.maxOpenConnections | int | `10` |  |
| database.maxQueryTimeout | string | `"500ms"` |  |
| database.name | string | `"financial_integration_service"` |  |
| database.password | string | `"financial_integration_service"` |  |
| database.port | int | `5432` |  |
| database.sslMode | string | `"disable"` |  |
| database.user | string | `"financial_integration_service"` |  |
| dependencies.temporal.clusterNamespace | string | `"temporal"` |  |
| dependencies.temporal.config.concurrentActivityExecutionSize | int | `5000` |  |
| dependencies.temporal.config.concurrentWorkflowTaskPollers | int | `100` |  |
| dependencies.temporal.config.namespace | string | `"simfiny"` |  |
| dependencies.temporal.config.taskQueue | string | `"financial-integration-service-task-queue"` |  |
| dependencies.temporal.config.workflowExecutionTimeout | string | `"1s"` |  |
| dependencies.temporal.config.workflowRunTimeout | string | `"1s"` |  |
| dependencies.temporal.config.workflowTaskTimeout | string | `"1s"` |  |
| dependencies.temporal.enabled | bool | `true` |  |
| dependencies.temporal.name | string | `"temporal-cluster-frontend"` |  |
| dependencies.temporal.port | int | `7233` |  |
| dependencies.temporal.retry.backoffCoefficient | float | `1.5` |  |
| dependencies.temporal.retry.maxRetryAttempts | int | `1` |  |
| dependencies.temporal.retry.maxRetryInterval | string | `"1s"` |  |
| dependencies.temporal.retry.retryInterval | string | `"100ms"` |  |
| dependencies.temporal.rpc.timeout | string | `"700ms"` |  |
| faults.delay | bool | `false` |  |
| faults.error | bool | `false` |  |
| faults.test_fail | bool | `false` |  |
| faults.test_timeout | bool | `false` |  |
| faults.unhealthy | bool | `false` |  |
| faults.unready | bool | `false` |  |
| h2c.enabled | bool | `false` |  |
| host | string | `nil` |  |
| image.pullPolicy | string | `"IfNotPresent"` |  |
| image.repository | string | `"feelguuds/financial-integration-service"` |  |
| image.tag | string | `"latest"` |  |
| ingress.annotations | object | `{}` |  |
| ingress.className | string | `""` |  |
| ingress.enabled | bool | `false` |  |
| ingress.hosts[0].host | string | `"financial-integration-service.local"` |  |
| ingress.hosts[0].paths[0].path | string | `"/"` |  |
| ingress.hosts[0].paths[0].pathType | string | `"ImplementationSpecific"` |  |
| ingress.tls | list | `[]` |  |
| integrationTests.enabled | bool | `false` |  |
| linkerd.annotations."linkerd.io/inject" | string | `"enabled"` |  |
| linkerd.annotations."prometheus.io/path" | string | `"/metrics"` |  |
| linkerd.annotations."prometheus.io/port" | string | `"4191\""` |  |
| linkerd.annotations."prometheus.io/scrape" | string | `"true"` |  |
| linkerd.profile.enabled | bool | `true` |  |
| logLevel | string | `"info"` |  |
| newRelic.apiKey | string | `"62fd721c712d5863a4e75b8f547b7c1ea884NRAL"` |  |
| newRelic.enabled | bool | `true` |  |
| nodeSelector | object | `{}` |  |
| openai.apiKey | string | `"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"` |  |
| openai.configs.frequencyPenalty | float | `0` |  |
| openai.configs.maxToken | int | `2000` |  |
| openai.configs.model | string | `"text-davinci-003"` |  |
| openai.configs.presencePenalty | int | `1` |  |
| openai.configs.temperature | float | `0.7` |  |
| openai.configs.topP | float | `1` |  |
| plaid.clientID | string | `"61eb5d49ea3b4700127560d1"` |  |
| plaid.enabled | bool | `true` |  |
| plaid.env | string | `"sandbox"` |  |
| plaid.oauthDomain | string | `"simfiny"` |  |
| plaid.products | string | `"transactions,auth,balance,investments,liabilities"` |  |
| plaid.secretKey | string | `"465686056e8fd1b87db3d993d096d8"` |  |
| plaid.webhookEnabled | bool | `true` |  |
| plaid.webhookOauthDomain | string | `"simfiny"` |  |
| podAnnotations | object | `{}` |  |
| postgresLocalDB.enabled | bool | `true` |  |
| redis.enabled | bool | `false` |  |
| redis.repository | string | `"redis"` |  |
| redis.tag | string | `"6.0.8"` |  |
| replicaCount | int | `3` |  |
| resources.limits | string | `nil` |  |
| resources.requests.cpu | string | `"1m"` |  |
| resources.requests.memory | string | `"16Mi"` |  |
| securityContext | object | `{}` |  |
| service.annotations | object | `{}` |  |
| service.enabled | bool | `true` |  |
| service.env | string | `"dev"` |  |
| service.environment | string | `"local"` |  |
| service.externalPort | int | `9898` |  |
| service.grpcPort | int | `9896` |  |
| service.grpcService | string | `"financial-integration-service"` |  |
| service.hostPort | string | `nil` |  |
| service.http.client.timeout | string | `"1s"` |  |
| service.http.server.shutdownTimeout | string | `"5s"` |  |
| service.http.server.timeout | string | `"1s"` |  |
| service.httpPort | int | `9898` |  |
| service.metricsPort | int | `9797` |  |
| service.nodePort | int | `31198` |  |
| service.rpc.timeout | string | `"1s"` |  |
| service.type | string | `"ClusterIP"` |  |
| serviceAccount.enabled | bool | `true` |  |
| serviceAccount.name | string | `nil` |  |
| serviceMonitor.additionalLabels | object | `{}` |  |
| serviceMonitor.enabled | bool | `false` |  |
| serviceMonitor.interval | string | `"15s"` |  |
| stripe.apiKey | string | `"sk_test_4eC39HqLyjWDarjtT1zdp7dc"` |  |
| stripe.enabled | bool | `true` |  |
| stripe.endpointSigningKey | string | `"whsec_21441814697a4a51dc01395a030498131d56ec4d7155bb216cc75f36548c86bf"` |  |
| taskProcessorWorkers | int | `5` | number of task processor worker threads to spin up |
| tls.certPath | string | `"/data/cert"` |  |
| tls.enabled | bool | `false` |  |
| tls.hostPort | string | `nil` |  |
| tls.port | int | `9899` |  |
| tls.secretName | string | `nil` |  |
| tolerations | list | `[]` |  |
| ui.color | string | `"#34577c"` |  |
| ui.logo | string | `""` |  |
| ui.message | string | `""` |  |

Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`. For example,

```console
$ helm install my-release \
  --set image.repository=feelguuds/financial-integration-service \
    simfinii/financial-integration-service
```

The above command sets the repository from which to get the container of interest.

> NOTE: Once this chart is deployed, it is not possible to change the application's access credentials, such as usernames or passwords, using Helm. To change these application credentials after deployment, delete any persistent volumes (PVs) used by the chart and re-deploy it, or use the application's built-in administrative tools if available.

Alternatively, a YAML file that specifies the values for the parameters can be provided while installing the chart. For example,

```console
$ helm install my-release -f values.yaml simfinii/financial-integration-service
```

> **Tip**: You can use the default [values.yaml](values.yaml)

## Troubleshooting

Find more information about how to deal with common errors related to Melodiy's Helm charts in [this troubleshooting guide](https://docs.bitnami.com/general/how-to/troubleshoot-helm-chart-issues).

## Upgrading

### To 0.0.1
Base chart version

---
## License

Copyright &copy; 2022 Melodiy

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.11.0](https://github.com/norwoodj/helm-docs/releases/v1.11.0)