import argparse
from concurrent import futures
import time
import grpc
import redis
from protobuf_to_dict import protobuf_to_dict

import transactionapi_pb2
import transactionapi_pb2_grpc

_ONE_DAY_IN_SECONDS = 60 * 60 * 24


class TransactionApi(transactionapi_pb2_grpc.TransactionApiServicer):

    def __init__(self, host, port):
        self.redis = redis.StrictRedis(host=host, port=port, db=0)

    def CreateTransaction(self, request, context):
        request_dict = protobuf_to_dict(request)
        # TODO: validate request
        success = self.redis.hmset(request.id, request_dict)
        return transactionapi_pb2.TransactionReply(success=success)

    def GetTransaction(self, request, context):
        result = self.redis.hgetall(request.id)
        # TODO: validate result
        return transactionapi_pb2.Transaction(name=result['name'], id=int(result['id']))


def serve(options):
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    transactionapi_pb2_grpc.add_TransactionApiServicer_to_server(TransactionApi(host=options.redis_host, port=options.redis_port), server)
    server.add_insecure_port(options.api_port)
    server.start()

    # gRPC starts a new thread to service requests. Just make the main thread
    # sleep.
    try:
        while True:
            time.sleep(_ONE_DAY_IN_SECONDS)
    except KeyboardInterrupt:
        server.stop(grace=0)


def parse_args():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument('--redis-host', default='redis-master', help='Redis host')
    parser.add_argument('--redis-port', default=6379, help='Redis port', type=int)
    parser.add_argument('--api-port', default='[::]:50051', help='Port of the gRPC api this service provides.')
    return parser.parse_args()


if __name__ == '__main__':
    options = parse_args()
    serve(options)
