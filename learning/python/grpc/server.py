from concurrent import futures
import time
import grpc
from proto import chat_controll_pb2
from proto import chat_controll_pb2_grpc

_ONE_DAY_IN_SECONDS = 60 * 60 * 24


class Greeter(chat_controll_pb2_grpc.GreeterServicer):
    # 工作函数
    def SayHello(self, request, context):
        print(request.name)
        message = "This message is from Server.And what i want to say is hello \" " + request.name + " \"";
        return chat_controll_pb2.HelloReply(message = message)


def serve():
    # gRPC 服务器
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    chat_controll_pb2_grpc.add_GreeterServicer_to_server(Greeter(), server)
    server.add_insecure_port('[::]:50051')
    print("sever is opening ,waiting for message...")
    server.start()  # start() 不会阻塞，如果运行时你的代码没有其它的事情可做，你可能需要循环等待。
    try:
        while True:
            time.sleep(_ONE_DAY_IN_SECONDS)
    except KeyboardInterrupt:
        server.stop(0)

if __name__ == '__main__':
    serve()