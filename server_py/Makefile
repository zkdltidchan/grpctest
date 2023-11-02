PIP=pip3
PYTHON=python3


SOURCE_PROTO_PATH=../proto
PYTHON_PROTO_PATH=./
# PROTO_NAME = goServer.proto
PROTO_NAME = pyServer.proto

protoinstall:
	$(PIP) install grpcio
	$(PIP) install grpcio-tools
	

protos:
	$(PYTHON) -m grpc_tools.protoc \
	-I$(SOURCE_PROTO_PATH) \
	--python_out=$(PYTHON_PROTO_PATH) \
	--grpc_python_out=$(PYTHON_PROTO_PATH) \
	$(SOURCE_PROTO_PATH)/$(PROTO_NAME)
