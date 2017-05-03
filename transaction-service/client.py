import argparse
import grpc

from transactionapi_pb2 import Transaction, TransactionRequest
from transactionapi_pb2_grpc import TransactionApiStub


def run(host, api_key=None):
    channel = grpc.insecure_channel(host)
    stub = TransactionApiStub(channel)
    metadata = []
    if api_key:
        metadata.append(('x-api-key', api_key))

    response = stub.CreateTransaction(Transaction(id=1, name='Hello World'), metadata=metadata)
    print('Client received: {}'.format(response.success))

    response = stub.GetTransaction(TransactionRequest(id=1), metadata=metadata)
    print('Client received: {}'.format(response))

if __name__ == '__main__':
    parser = argparse.ArgumentParser(description=__doc__, formatter_class=argparse.RawDescriptionHelpFormatter)
    parser.add_argument('--host', default='localhost:50051', help='The server host.')
    parser.add_argument('--api_key', default=None, help='The API key to use for the call.')
    args = parser.parse_args()
    run(args.host, args.api_key)
