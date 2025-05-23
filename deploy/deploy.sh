#!/bin/bash 

namespace=staging
while getopts n: flag
do
    case "${flag}" in
        n) namespace=${OPTARG};;
    esac
done

namespaceStatus=$(kubectl get ns ${namespace} -o json | jq .status.phase -r)
if [[ $namespaceStatus != "Active" ]]; then 
    echo "creating namespace ($namespace) in which all services will be deployed"
    kubectl create namespace ${namespace}
fi

if [[ $namespace == "prod" ]]; then 
    echo "installing service in production namespace"
    helm upgrade --install financial-integration-service ./charts/financial-integration-service --values ./charts/financial-integration-service/values.production.yaml -n ${namespace}
else 
    echo "installing service in staging namespace"
    helm upgrade --install financial-integration-service ./charts/financial-integration-service --values ./charts/financial-integration-service/values.staging.yaml -n ${namespace}
fi