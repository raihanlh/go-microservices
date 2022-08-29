# kubectl --kubeconfig=go-microservices-kubeconfig.yaml get nodes
# kubectl --kubeconfig=go-microservices-kubeconfig.yaml apply -f ./k8
kubectl apply -f ./k8
# kubectl rollout restart -f ./k8