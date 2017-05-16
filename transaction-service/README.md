# Python Transaction Service

## Run

#### Run locally

Setup service
```bash
# Install dependencies (consider to create a virtual environment first)
pip install -r requirements.txt
```

Run server
```bash
python server.py
```

When running the client it should receive a message
```bash
python client.py
```

## gRPC

```bash
# Generate Python files
python -m grpc_tools.protoc -I protos --python_out=. --grpc_python_out=. protos/transactionapi.proto
# Generate protocol buffer file
python -m grpc_tools.protoc --include_imports --include_source_info -I protos protos/transactionapi.proto --descriptor_set_out out.pb
```

## API (gRPC)

#### Transactions

see [transactionapi.proto](protos/transactionapi.proto)

## Credits

Bootstrapped with the tutorial [Endpoints Getting Started with gRPC & Python Quickstart](https://github.com/GoogleCloudPlatform/python-docs-samples/tree/master/endpoints/getting-started-grpc)
