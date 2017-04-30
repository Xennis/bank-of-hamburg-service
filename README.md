## Bank of Hamburg Service

This is a simple service of a fake bank named Bank of Hamburg. You can create a bank account and pay money in our out.

## Go backend service

#### Run locally

Run the service
```bash
cd go-backend-service
go run *.go
```

#### Build and run Docker container

Build the docker container and run the service
```bash
cd go-backend-service
docker build -t bank-of-hamburg-backend
docker run --publish 8080:8080 --name bank-of-hamburg-backend-test --rm bank-of-hamburg-backend
```

### API

#### Accounts

Request (1) all bank accounts, (2) a single account or (3) create a account.
```bash
curl http://localhost:8080/api/accounts
curl http://localhost:8080/api/accounts/2
curl -H "Content-Type: application/json" -d '{"name":"Kat Müller"}' http://localhost:8080/api/accounts
```

#### Transactions

Transfer the amount of 1000 units from account 1 to account 2
```bash
curl -H "Content-Type: application/json" -d '{"from":1, "to": 2, "amount": 1000}' http://localhost:8080/api/transactions
```

(1) Pay in 50 units to account 1 or (2) pay it out 50 units from account 1
```bash
curl -H "Content-Type: application/json" -d '{"to": 1, "amount": 50}' http://localhost:8080/api/transactions
curl -H "Content-Type: application/json" -d '{"from": 1, "amount": 50}' http://localhost:8080/api/transactions
```

#### Service

Liveness and readiness probes
```bash
curl -v http://localhost:8080/healthz
curl -v http://localhost:8080/readiness
```

### Credits

* Go backend service: Initial service based on the idea of the tutorial [Making a RESTful JSON API in Go](https://thenewstack.io/make-a-restful-json-api-go/) by Cory Lanou.
* Redis setup: [Create a Guestbook with Redis and PHP](https://cloud.google.com/container-engine/docs/tutorials/guestbook) by Google Cloud Platform
