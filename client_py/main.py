import grpc
import pyServer_pb2

import pyServer_pb2_grpc

PYTHON_SERVER_PORT="50052"

def run():
    # connect to grpc server
    channel = grpc.insecure_channel('localhost:'+PYTHON_SERVER_PORT)
    stub = pyServer_pb2_grpc.PyTestServiceStub(channel)

    # call rpc
    response = stub.GetTestPY(
        pyServer_pb2.TestRequest(testName='test python')
        )
    print("Client received: " + response.message)

if __name__ == '__main__':
    run()
