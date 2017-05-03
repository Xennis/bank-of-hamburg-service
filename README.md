# Bank of Hamburg Service

This is a simple service of a fake bank named Bank of Hamburg. You can create a bank account and pay money in our out. The service consists of
* [Account service written in Go](account-service/README.md)
* [Transaction service written in Python](transaction-service/README.md)
* Redis master for updates
* Redis slave for search

## Run service on Kubernetes cluster

Create a cluster, for instance on Google Cloud
```bash
gcloud container clusters create <name> --num-nodes=3 --disk-size=10
```

```bash
# Config map
kubectl create configmap backend-configmap --from-literal=redis.port=6379
# Service and deployment of the redis master and slaves
kubectl create -f redis-master-slave.yaml
# Service (load balancer) and deployment of the account service
kubectl create -f account-service.yaml
# Service (and deployment of the transaction service
kubectl create -f transaction-service.yaml
```

```bash
# Generate protocol buffer file
cd transaction-service
python -m grpc_tools.protoc --include_imports --include_source_info -I protos protos/transactionapi.proto --descriptor_set_out out.pb
# Deploy rpc api
gcloud service-management deploy out.pb ../transaction-api-config.yaml
```

## Credits

Redis setup: [Create a Guestbook with Redis and PHP](https://cloud.google.com/container-engine/docs/tutorials/guestbook) by Google Cloud Platform
