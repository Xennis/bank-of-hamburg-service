# Bank of Hamburg Service

This is a simple service of a fake bank named Bank of Hamburg. You can create a bank account and pay money in our out. The service consists of
* [Account service written in Go](account-service/README.md)
* Redis master for updates
* Redis slave for search

## Run service on Kubernetes cluster

```bash
# Config map
kubectl create configmap backend-configmap --from-literal=redis.port=6379
# Service and deployment of the redis master
kubectl create -f redis-master.yaml
# Service and deployment of the redis slave
kubectl create -f redis-slave.yaml
# Service (load balancer) and deployment of the account service
kubectl create -f account-service.yaml
```

## Credits

Redis setup: [Create a Guestbook with Redis and PHP](https://cloud.google.com/container-engine/docs/tutorials/guestbook) by Google Cloud Platform
