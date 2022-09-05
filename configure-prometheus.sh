helm repo add prometheus-community https://prometheus-community.github.io/helm-charts 
helm repo update
helm install prometheus prometheus-community/prometheus -f prometheus/prometheus-values.yaml
kubectl apply -f ingress/ingress-prometheus.yaml
# helm upgrade prometheus prometheus-community/prometheus -f prometheus/prometheus-values.yaml