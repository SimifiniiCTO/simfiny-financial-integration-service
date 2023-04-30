# financial-integration-service

Financial Integration Service Helm chart for Kubernetes

## TL;DR

```console
$ helm repo add simfinii https://github.com/SimifiniiCTO/simfinii
$ helm install my-release simfinii/financial-integration-service
```

**Homepage:** <https://github.com/SimifiniiCTO/simfiny-financial-integration-service>

## Maintainers

| Name       | Email             | Url |
|------------|-------------------|-----|
| yoan yomba | yoan@simfinii.com |     |

![Version: 6.0.3](https://img.shields.io/badge/Version-6.0.3-informational?style=flat-square) ![AppVersion: 6.0.3](https://img.shields.io/badge/AppVersion-6.0.3-informational?style=flat-square) ![Version: 6.0.3](https://img.shields.io/badge/Version-6.0.3-informational?style=flat-square) ![AppVersion: 6.0.3](https://img.shields.io/badge/AppVersion-6.0.3-informational?style=flat-square)

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

This chart bootstraps the financial integration service deployment on a Kubernetes cluster using the Helm package
manager.

This chart can be used with Kubeapps for deployment and management of Helm Charts in clusters. This Helm chart has been
tested on top of a Kubernetes Runtime.

## Requirements

Kubernetes: `>=1.19.0-0`

* Kubernetes 1.19+
* Helm 3.2.0+
* PV provisioner support in the underlying infrastructure

## Source Code

* <https://github.com/SimifiniiCTO/simfiny-financial-integration-service>

## Installing the Chart

To install the chart with the release name `my-release`:

```console
$ helm repo add simfinii https://github.com/SimifiniiCTO/simfinii
$ helm install my-release simfinii/financial-integration-service
```

These commands deploy the financial-integration-service on the Kubernetes cluster in the default configuration. The
Parameters section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall/delete the `my-release` deployment:

```console
$ helm delete my-release
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Values

| Key                                          | Type   | Default                                                       | Description |
|----------------------------------------------|--------|---------------------------------------------------------------|-------------|
| affinity                                     | object | `{}`                                                          |             |
| backend                                      | string | `nil`                                                         |             |
| backends                                     | list   | `[]`                                                          |             |
| cache                                        | string | `""`                                                          |             |
| certificate.create                           | bool   | `false`                                                       |             |
| certificate.dnsNames[0]                      | string | `"financial-integration-service"`                             |             |
| certificate.issuerRef.kind                   | string | `"ClusterIssuer"`                                             |             |
| certificate.issuerRef.name                   | string | `"self-signed"`                                               |             |
| database.host                                | string | `"financial-integration-service-db"`                          |             |
| database.max_connection_attempts             | int    | `2`                                                           |             |
| database.max_connection_retries              | int    | `2`                                                           |             |
| database.max_connection_retry_sleep_interval | string | `"100ms"`                                                     |             |
| database.max_connection_retry_timeout        | string | `"500ms"`                                                     |             |
| database.name                                | string | `"financial_integration_service"`                             |             |
| database.password                            | string | `"financial_integration_service"`                             |             |
| database.port                                | int    | `5432`                                                        |             |
| database.sslMode                             | string | `"disable"`                                                   |             |
| database.user                                | string | `"financial_integration_service"`                             |             |
| faults.delay                                 | bool   | `false`                                                       |             |
| faults.error                                 | bool   | `false`                                                       |             |
| faults.test_fail                             | bool   | `false`                                                       |             |
| faults.test_timeout                          | bool   | `false`                                                       |             |
| faults.unhealthy                             | bool   | `false`                                                       |             |
| faults.unready                               | bool   | `false`                                                       |             |
| h2c.enabled                                  | bool   | `false`                                                       |             |
| host                                         | string | `nil`                                                         |             |
| hpa.cpu                                      | string | `nil`                                                         |             |
| hpa.enabled                                  | bool   | `false`                                                       |             |
| hpa.maxReplicas                              | int    | `10`                                                          |             |
| hpa.memory                                   | string | `nil`                                                         |             |
| hpa.requests                                 | string | `nil`                                                         |             |
| image.pullPolicy                             | string | `"IfNotPresent"`                                              |             |
| image.repository                             | string | `"feelguuds/financial-integration-service"`                   |             |
| image.tag                                    | string | `"latest"`                                                    |             |
| ingress.annotations                          | object | `{}`                                                          |             |
| ingress.className                            | string | `""`                                                          |             |
| ingress.enabled                              | bool   | `false`                                                       |             |
| ingress.hosts[0].host                        | string | `"financial-integration-service.local"`                       |             |
| ingress.hosts[0].paths[0].path               | string | `"/"`                                                         |             |
| ingress.hosts[0].paths[0].pathType           | string | `"ImplementationSpecific"`                                    |             |
| ingress.tls                                  | list   | `[]`                                                          |             |
| integrationTests.enabled                     | bool   | `false`                                                       |             |
| linkerd.profile.enabled                      | bool   | `false`                                                       |             |
| logLevel                                     | string | `"info"`                                                      |             |
| newRelic.enabled                             | bool   | `true`                                                        |             |
| newRelic.key                                 | string | `"62fd721c712d5863a4e75b8f547b7c1ea884NRAL"`                  |             |
| nodeSelector                                 | object | `{}`                                                          |             |
| plaid.client_id                              | string | `"61eb5d49ea3b4700127560d1"`                                  |             |
| plaid.enabled                                | bool   | `true`                                                        |             |
| plaid.env                                    | string | `"sandbox"`                                                   |             |
| plaid.products                               | string | `"transactions,auth,identity,assets,investments,liabilities"` |             |
| plaid.redirect_url                           | string | `""`                                                          |             |
| plaid.secret_key                             | string | `"465686056e8fd1b87db3d993d096d8"`                            |             |
| plaid.webhook_url                            | string | `""`                                                          |             |
| podAnnotations                               | object | `{}`                                                          |             |
| postgresLocalDB.enabled                      | bool   | `true`                                                        |             |
| redis.enabled                                | bool   | `false`                                                       |             |
| redis.repository                             | string | `"redis"`                                                     |             |
| redis.tag                                    | string | `"6.0.8"`                                                     |             |
| replicaCount                                 | int    | `1`                                                           |             |
| resources.limits                             | string | `nil`                                                         |             |
| resources.requests.cpu                       | string | `"1m"`                                                        |             |
| resources.requests.memory                    | string | `"16Mi"`                                                      |             |
| securityContext                              | object | `{}`                                                          |             |
| service.annotations                          | object | `{}`                                                          |             |
| service.enabled                              | bool   | `true`                                                        |             |
| service.env                                  | string | `"dev"`                                                       |             |
| service.environment                          | string | `"local"`                                                     |             |
| service.externalPort                         | int    | `9898`                                                        |             |
| service.grpcPort                             | int    | `9896`                                                        |             |
| service.grpcService                          | string | `"financial-integration-service"`                             |             |
| service.hostPort                             | string | `nil`                                                         |             |
| service.httpPort                             | int    | `9898`                                                        |             |
| service.metricsPort                          | int    | `9797`                                                        |             |
| service.nodePort                             | int    | `31198`                                                       |             |
| service.type                                 | string | `"ClusterIP"`                                                 |             |
| serviceAccount.enabled                       | bool   | `false`                                                       |             |
| serviceAccount.name                          | string | `nil`                                                         |             |
| serviceMonitor.additionalLabels              | object | `{}`                                                          |             |
| serviceMonitor.enabled                       | bool   | `false`                                                       |             |
| serviceMonitor.interval                      | string | `"15s"`                                                       |             |
| tls.certPath                                 | string | `"/data/cert"`                                                |             |
| tls.enabled                                  | bool   | `false`                                                       |             |
| tls.hostPort                                 | string | `nil`                                                         |             |
| tls.port                                     | int    | `9899`                                                        |             |
| tls.secretName                               | string | `nil`                                                         |             |
| tolerations                                  | list   | `[]`                                                          |             |
| ui.color                                     | string | `"#34577c"`                                                   |             |
| ui.logo                                      | string | `""`                                                          |             |
| ui.message                                   | string | `""`                                                          |             |

Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`. For example,

```console
$ helm install my-release \
  --set image.repository=keratin/authn-server \
    simfinii/authentication-service
```

The above command sets the repository from which to get the container user of interest.

> NOTE: Once this chart is deployed, it is not possible to change the application's access credentials, such as
> usernames or passwords, using Helm. To change these application credentials after deployment, delete any persistent
> volumes (PVs) used by the chart and re-deploy it, or use the application's built-in administrative tools if available.

Alternatively, a YAML file that specifies the values for the parameters can be provided while installing the chart. For
example,

```console
$ helm install my-release -f values.yaml simfinii/financial-integration-service
```

> **Tip**: You can use the default [values.yaml](values.yaml)

## Troubleshooting

Find more information about how to deal with common errors related to Bitnami's Helm charts
in [this troubleshooting guide](https://docs.bitnami.com/general/how-to/troubleshoot-helm-chart-issues).

## Upgrading

### To 0.0.1

Base chart version

---

## License

Copyright &copy; 2022 Simfinii

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
Autogenerated from chart metadata using [helm-docs v1.6.0](https://github.com/norwoodj/helm-docs/releases/v1.6.0)