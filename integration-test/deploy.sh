#!/bin/bash 

namespace=default
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

echo "installing user service in default namespace"
helm upgrade --install service ./charts/financial-integration-service \
             --values ./charts/financial-integration-service/values.production.yaml \
             --set replicaCount=1 \
             --set linkerd.profile.enabled=false \
             -n ${namespace}