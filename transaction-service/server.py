from concurrent import futures
import time
import grpc

import transactionapi_pb2
import transactionapi_pb2_grpc

_ONE_DAY_IN_SECONDS = 60 * 60 * 24


class TransactionApi(transactionapi_pb2_grpc.TransactionServicer):

    def CreateTransaction(self, request, context):
        return transactionapi_pb2.TransactionReply(message='Hello, %s!' % request.name)


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    transactionapi_pb2_grpc.add_TransactionServicer_to_server(TransactionApi(), server)
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
