### Install dashboard

`kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v1.10.1/src/deploy/recommended/kubernetes-dashboard.yaml`

### Additional features

```
kubectl create -f https://raw.githubusercontent.com/kubernetes/heapster/master/deploy/kube-config/influxdb/influxdb.yaml
kubectl create -f https://raw.githubusercontent.com/kubernetes/heapster/master/deploy/kube-config/influxdb/heapster.yaml
kubectl create -f https://raw.githubusercontent.com/kubernetes/heapster/master/deploy/kube-config/influxdb/grafana.yaml
```

### Starting dashboard

`kubectl proxy`

### Navigate to dashboard

http://localhost:8001/api/v1/namespaces/kube-system/services/https:kubernetes-dashboard:/proxy/#!/namespace?namespace=default

### Getting token

```
kubectl -n kube-system get secret
kubectl -n kube-system describe secrets replicaset-controller-token-btkvl
```

### Start zookeper

```
kubectl apply -f .\build\Kafka\deployment-zookeeper.yaml
kubectl apply -f .\build\Kafka\service-zookeeper.yaml
```

Or

```
kubectl run zookeeper --image=zookeeper:3.4 --port=2181
kubectl expose deployment zookeeper --type=NodePort
kubectl get services (for info)
```

### Start kafka

```
kubectl apply -f .\build\Kafka\deployment-kafka.yaml
kubectl apply -f .\build\Kafka\service-kafka.yaml
```

Or

```
kubectl run kafka --image=ches/kafka --port=7203 --port=9092 --env="KAFKA_MESSAGE_MAX_BYTES=3000000" --env="KAFKA_REPLICA_FETCH_MAX_BYTES=3100000" --env="ZOOKEEPER_IP=zookeeper" --env="KAFKA_PORT=9092"
kubectl get pods
kubectl logs [kafka-***] (for info)
kubectl expose deployment kafka --type=NodePort
kubectl get services (for info)
```

### Deploy RPC Server

```
docker build -f .\build\Review.CommandRpcServer\Dockerfile -t command-rpcserver:latest .
kubectl apply -f .\build\Review.CommandRpcServer\deployment.yaml
kubectl apply -f .\build\Review.CommandRpcServer\service.yaml
```

### Deploy Review api

```
docker build -f .\build\Review.API\Dockerfile -t review-api:latest .
kubectl apply -f .\build\Review.API\deployment.yaml
kubectl apply -f .\build\Review.API\service.yaml

kubectl get services // for checking details
kubectl get pods
kubectl attach review-api-[****]
```

Navigate to http://localhost:31115/swaggerui/

### Some helper codes

`kubectl delete deployments [deployment_name]`
`kubectl run topiccreation --image=ches/kafka --env="ZOOKEEPER_IP=10.100.165.232" --command -- sh -c "sleep 5 && kafka-topics.sh --create --topic review-commands --replication-factor 1 --partitions 1 --zookeeper 10.100.165.232:2181"`
`kubectl attach POD_NAME`

### Notes

https://www.hanselman.com/blog/HowToSetUpKubernetesOnWindows10WithDockerForWindowsAndRunASPNETCore.aspx  
https://github.com/kubernetes/dashboard/wiki/Access-control  
https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/

TOKEN  
eyJhbGciOiJSUzI1NiIsImtpZCI6IiJ9.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJrdWJlLXN5c3RlbSIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VjcmV0Lm5hbWUiOiJyZXBsaWNhc2V0LWNvbnRyb2xsZXItdG9rZW4tYnRrdmwiLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlcnZpY2UtYWNjb3VudC5uYW1lIjoicmVwbGljYXNldC1jb250cm9sbGVyIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQudWlkIjoiYjQ3MjYwNmQtMmUwNC0xMWU5LWEyNDktMDAxNTVkMGM1MzFiIiwic3ViIjoic3lzdGVtOnNlcnZpY2VhY2NvdW50Omt1YmUtc3lzdGVtOnJlcGxpY2FzZXQtY29udHJvbGxlciJ9.eEhFUPa1A581citnCdq8pzvMi34JNCnlwgtBD3nUpFvW6qBP_hl6y74jBoy9FWcRMMwsdZx9V6ZivQ5vPCgTLlIqBU-\_1rEHvboWuzobXMA9n6z_3ay7vuAdOqO0mVeHJ_5qOtN1XFXoF5EalHUXrIH5jiI37k4ugJt4hCHAcL6recGZpn0-CCODb96ESohcJbhbhQrivU5BAQZOcgTWc280E9yyO4ABS23cIultzedWGlE-iJYnwB1SUQFjXyaxw3PeztdiueHkfP-lO42-FZ9pOgDuscU29SYPp_ldWct4yxP3c60GthfZftUBjJfn_ELX6dCo9chNsYeuqqWBBA
