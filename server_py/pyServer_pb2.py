# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: pyServer.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0epyServer.proto\x12\x0cpyServerTest\"\x1f\n\x0bTestRequest\x12\x10\n\x08testName\x18\x01 \x01(\t\"\x1f\n\x0cTestResponse\x12\x0f\n\x07message\x18\x01 \x01(\t2S\n\rPyTestService\x12\x42\n\tGetTestPY\x12\x19.pyServerTest.TestRequest\x1a\x1a.pyServerTest.TestResponseB\x10Z\x0e.;pyServerTestb\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'pyServer_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\016.;pyServerTest'
  _TESTREQUEST._serialized_start=32
  _TESTREQUEST._serialized_end=63
  _TESTRESPONSE._serialized_start=65
  _TESTRESPONSE._serialized_end=96
  _PYTESTSERVICE._serialized_start=98
  _PYTESTSERVICE._serialized_end=181
# @@protoc_insertion_point(module_scope)