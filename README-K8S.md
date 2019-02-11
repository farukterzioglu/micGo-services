# Install dashboard

kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v1.10.1/src/deploy/recommended/kubernetes-dashboard.yaml

# Additional features

kubectl create -f https://raw.githubusercontent.com/kubernetes/heapster/master/deploy/kube-config/influxdb/influxdb.yaml  
kubectl create -f https://raw.githubusercontent.com/kubernetes/heapster/master/deploy/kube-config/influxdb/heapster.yaml  
kubectl create -f https://raw.githubusercontent.com/kubernetes/heapster/master/deploy/kube-config/influxdb/grafana.yaml

# Starting dashboard

kubectl proxy

# Navigate to dashboard

http://localhost:8001/api/v1/namespaces/kube-system/services/https:kubernetes-dashboard:/proxy/#!/namespace?namespace=default

# Getting token

kubectl -n kube-system get secret  
kubectl -n kube-system describe secrets replicaset-controller-token-btkvl

# Start zookeper

kubectl run zookeeper --image=zookeeper:3.4 --port=2181  
kubectl expose deployment zookeeper --type=NodePort  
kubectl get services (get zookeper ip)

# Start kafka

kubectl run kafka --image=ches/kafka --port=7203 --port=9092 --env="KAFKA_MESSAGE_MAX_BYTES=3000000" --env="KAFKA_REPLICA_FETCH_MAX_BYTES=3100000" --env="ZOOKEEPER_IP=10.100.165.232" --env="KAFKA_ADVERTISED_HOST_NAME=172.24.96.1"  
kubectl get pods  
kubectl logs kafka-698855978-2wnhl  
kubectl expose deployment kafka --type=NodePort

# Create topic

// TODO : Run once commands ???  
kubectl run topiccreation --image=ches/kafka --env="ZOOKEEPER_IP=10.100.165.232" --command -- sh -c "sleep 5 && kafka-topics.sh --create --topic review-commands --replication-factor 1 --partitions 1 --zookeeper 10.100.165.232:2181"

# Deploy Review api

docker build -f .\build\Review.API\Dockerfile -t review-api:latest .  
kubectl get services

# Util.

kubectl delete deployments [deploymebnt_name]

NOTES  
https://www.hanselman.com/blog/HowToSetUpKubernetesOnWindows10WithDockerForWindowsAndRunASPNETCore.aspx  
https://github.com/kubernetes/dashboard/wiki/Access-control  
https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/

TOKEN  
eyJhbGciOiJSUzI1NiIsImtpZCI6IiJ9.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJrdWJlLXN5c3RlbSIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VjcmV0Lm5hbWUiOiJyZXBsaWNhc2V0LWNvbnRyb2xsZXItdG9rZW4tYnRrdmwiLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlcnZpY2UtYWNjb3VudC5uYW1lIjoicmVwbGljYXNldC1jb250cm9sbGVyIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQudWlkIjoiYjQ3MjYwNmQtMmUwNC0xMWU5LWEyNDktMDAxNTVkMGM1MzFiIiwic3ViIjoic3lzdGVtOnNlcnZpY2VhY2NvdW50Omt1YmUtc3lzdGVtOnJlcGxpY2FzZXQtY29udHJvbGxlciJ9.eEhFUPa1A581citnCdq8pzvMi34JNCnlwgtBD3nUpFvW6qBP_hl6y74jBoy9FWcRMMwsdZx9V6ZivQ5vPCgTLlIqBU-\_1rEHvboWuzobXMA9n6z_3ay7vuAdOqO0mVeHJ_5qOtN1XFXoF5EalHUXrIH5jiI37k4ugJt4hCHAcL6recGZpn0-CCODb96ESohcJbhbhQrivU5BAQZOcgTWc280E9yyO4ABS23cIultzedWGlE-iJYnwB1SUQFjXyaxw3PeztdiueHkfP-lO42-FZ9pOgDuscU29SYPp_ldWct4yxP3c60GthfZftUBjJfn_ELX6dCo9chNsYeuqqWBBA
