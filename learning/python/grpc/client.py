from __future__ import print_function

import grpc

from proto import chat_controll_pb2
from proto import chat_controll_pb2_grpc


def run():
    channel = grpc.insecure_channel('localhost:50051')
    stub = chat_controll_pb2_grpc.GreeterStub(channel)
    response = stub.SayHello(chat_controll_pb2.HelloRequest(name='Hello World！ This is message from client!'))
    print("Greeter client received: " + response.message)


if __name__ == '__main__':
    run()