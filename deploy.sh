#!/bin/bash

# kubectl --kubeconfig=go-microservices-kubeconfig.yaml get nodes
# kubectl --kubeconfig=go-microservices-kubeconfig.yaml apply -f ./k8
kubectl apply -f ./k8
./configure-prometheus.sh
sleep 1m # Wait for prometheus config to finish
./configure-ingress.sh
# kubectl rollout restart -f ./k8