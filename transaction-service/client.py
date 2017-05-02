import argparse
import grpc

import transactionapi_pb2
import transactionapi_pb2_grpc


def run(host, api_key=None):
    channel = grpc.insecure_channel(host)
    stub = transactionapi_pb2_grpc.TransactionStub(channel)
    metadata = []
    if api_key:
        metadata.append(('x-api-key', api_key))
    response = stub.CreateTransaction(transactionapi_pb2.TransactionRequest(name='you'), metadata=metadata)
    print('Client received: ' + response.message)


if __name__ == '__main__':
    parser = argparse.ArgumentParser(description=__doc__, formatter_class=argparse.RawDescriptionHelpFormatter)
    parser.add_argument('--host', default='localhost:50051', help='The server host.')
    parser.add_argument('--api_key', default=None, help='The API key to use for the call.')
    args = parser.parse_args()
    run(args.host, args.api_key)
