## Go backend service

#### Run locally

Run the service
```bash
cd go-backend-service
go run *.go
```

#### Run Docker container

Build the docker container and run the service
```bash
cd go-backend-service
docker build -t bank-of-hamburg-backend .
docker run --publish 8080:8080 --name bank-of-hamburg-backend-test --rm bank-of-hamburg-backend
```

## gRPC

```bash
# Generate Go file
protoc -I protos --go_out=plugins=grpc:../account-service protos/transactionapi.proto
```

## API (REST)

#### Accounts

```bash
# Request all bank accounts
curl http://localhost:8080/api/accounts
# Request account with ID 2
curl http://localhost:8080/api/accounts/2
# Create an account
curl -H "Content-Type: application/json" -d '{"name":"Kat Müller"}' http://localhost:8080/api/accounts
```

#### Transactions

```bash
# Transfer the amount of 1000 units from account 1 to account 2
curl -H "Content-Type: application/json" -d '{"from":1, "to": 2, "amount": 1000}' http://localhost:8080/api/transactions
# Pay in 50 units to account 1
curl -H "Content-Type: application/json" -d '{"to": 1, "amount": 50}' http://localhost:8080/api/transactions
# Pay out 50 units from account 1
curl -H "Content-Type: application/json" -d '{"from": 1, "amount": 50}' http://localhost:8080/api/transactions
```

#### Service

```bash
# Liveness probe
curl -v http://localhost:8080/healthz
# Readiness probe
curl -v http://localhost:8080/readiness
```

## Credits

Initial service based on the idea of the tutorial [Making a RESTful JSON API in Go](https://thenewstack.io/make-a-restful-json-api-go/) by Cory Lanou.
