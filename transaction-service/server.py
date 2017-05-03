from concurrent import futures
import time
import grpc
import redis
from protobuf_to_dict import protobuf_to_dict

import transactionapi_pb2
import transactionapi_pb2_grpc

_ONE_DAY_IN_SECONDS = 60 * 60 * 24


class TransactionApi(transactionapi_pb2_grpc.TransactionApiServicer):

    def __init__(self):
        self.redis = redis.StrictRedis(host='redis-master', port=6379, db=0)

    def CreateTransaction(self, request, context):
        request_dict = protobuf_to_dict(request)
        # TODO: validate request
        success = self.redis.hmset(request.id, request_dict)
        return transactionapi_pb2.TransactionReply(success=success)

    def GetTransaction(self, request, context):
        result = self.redis.hgetall(request.id)
        # TODO: validate result
        return transactionapi_pb2.Transaction(name=result['name'], id=int(result['id']))


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    transactionapi_pb2_grpc.add_TransactionApiServicer_to_server(TransactionApi(), server)
    server.add_insecure_port('[::]:50051')
    server.start()

    # gRPC starts a new thread to service requests. Just make the main thread
    # sleep.
    try:
        while True:
            time.sleep(_ONE_DAY_IN_SECONDS)
    except KeyboardInterrupt:
        server.stop(grace=0)


if __name__ == '__main__':
    serve()
