## Bank of Hamburg Service

This is a simple service of a fake bank named Bank of Hamburg. You can create
a bank account and pay money in our out.


### Run

#### Run locally

Run the service
```bash
go run *.go
```

#### Run as Docker container

Build the docker container and run the service
```bash
docker build -t bank-of-hamburg .
docker run --publish 6060:8080 --name bank-of-hamburg-test --rm bank-of-hamburg
```

### API

* Locally: http://localhost:8080
* Docker: http://localhost:6060

#### Accounts

Model: [account.go](account.go)

Request (1) all bank accounts, (2) a single account or (3) create a account.
```bash
curl http://localhost:8080/accounts
curl http://localhost:8080/accounts/2
curl -H "Content-Type: application/json" -d '{"name":"Kat Müller"}' http://localhost:8080/accounts
```

#### Transactions

Model: [transaction.go](transaction.go)

Transfer the amount of 1000 units from account 1 to account 2
```bash
curl -H "Content-Type: application/json" -d '{"from":1, "to": 2, "amount": 1000}' http://localhost:8080/transactions
```

(1) Pay in 50 units to account 1 or (2) pay it out 50 units from account 1
```bash
curl -H "Content-Type: application/json" -d '{"to": 1, "amount": 50}' http://localhost:8080/transactions
curl -H "Content-Type: application/json" -d '{"from": 1, "amount": 50}' http://localhost:8080/transactions
```

### Credits

Parts of the code based on the tutorial [Making a RESTful JSON API in Go](https://thenewstack.io/make-a-restful-json-api-go/) by Cory Lanou.
