import grpc
from concurrent import futures
import pyServer_pb2
import pyServer_pb2_grpc

PYTHON_SERVER_PORT="50052"

class PyTestService(pyServer_pb2_grpc.PyTestServiceServicer):
    def GetTestPY(self, request, context):
        print("Server received: " + request.testName)
        res = pyServer_pb2.TestResponse(
            message=f"Hello, {request.testName}!"
            )
        return res

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    pyServer_pb2_grpc.add_PyTestServiceServicer_to_server(PyTestService(), server)
    server.add_insecure_port('[::]:'+PYTHON_SERVER_PORT)
    server.start()
    server.wait_for_termination()

if __name__ == '__main__':
    serve()
