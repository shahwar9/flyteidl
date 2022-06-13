# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: flyteidl/service/signal.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from flyteidl.admin import signal_pb2 as flyteidl_dot_admin_dot_signal__pb2
from protoc_gen_swagger.options import annotations_pb2 as protoc__gen__swagger_dot_options_dot_annotations__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='flyteidl/service/signal.proto',
  package='flyteidl.service',
  syntax='proto3',
  serialized_options=_b('Z7github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/service'),
  serialized_pb=_b('\n\x1d\x66lyteidl/service/signal.proto\x12\x10\x66lyteidl.service\x1a\x1cgoogle/api/annotations.proto\x1a\x1b\x66lyteidl/admin/signal.proto\x1a,protoc-gen-swagger/options/annotations.proto2\xee\x05\n\rSignalService\x12\x90\x01\n\x11GetOrCreateSignal\x12(.flyteidl.admin.SignalGetOrCreateRequest\x1a\x16.flyteidl.admin.Signal\"9\x92\x41\x36\x1a\x34Retrieve a signal, creating it if it does not exist.\x12\x95\x02\n\x0bListSignals\x12!.flyteidl.admin.SignalListRequest\x1a!.flyteidl.admin.SignalListRequest\"\xbf\x01\x82\xd3\xe4\x93\x02m\x12k/api/v1/signals/{workflow_execution_id.project}/{workflow_execution_id.domain}/{workflow_execution_id.name}\x92\x41I\x1aGFetch existing signal definitions matching the input signal id filters.\x12\xb1\x02\n\tSetSignal\x12 .flyteidl.admin.SignalSetRequest\x1a!.flyteidl.admin.SignalSetResponse\"\xde\x01\x82\xd3\xe4\x93\x02\x14\"\x0f/api/v1/signals:\x01*\x92\x41\xc0\x01\x1a\x13Set a signal value.JB\n\x03\x34\x30\x30\x12;\n9Returned for bad request that may have failed validation.Je\n\x03\x34\x30\x39\x12^\n\\Returned for a request that references an identical entity that has already been registered.B9Z7github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/serviceb\x06proto3')
  ,
  dependencies=[google_dot_api_dot_annotations__pb2.DESCRIPTOR,flyteidl_dot_admin_dot_signal__pb2.DESCRIPTOR,protoc__gen__swagger_dot_options_dot_annotations__pb2.DESCRIPTOR,])



_sym_db.RegisterFileDescriptor(DESCRIPTOR)


DESCRIPTOR._options = None

_SIGNALSERVICE = _descriptor.ServiceDescriptor(
  name='SignalService',
  full_name='flyteidl.service.SignalService',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=157,
  serialized_end=907,
  methods=[
  _descriptor.MethodDescriptor(
    name='GetOrCreateSignal',
    full_name='flyteidl.service.SignalService.GetOrCreateSignal',
    index=0,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_signal__pb2._SIGNALGETORCREATEREQUEST,
    output_type=flyteidl_dot_admin_dot_signal__pb2._SIGNAL,
    serialized_options=_b('\222A6\0324Retrieve a signal, creating it if it does not exist.'),
  ),
  _descriptor.MethodDescriptor(
    name='ListSignals',
    full_name='flyteidl.service.SignalService.ListSignals',
    index=1,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_signal__pb2._SIGNALLISTREQUEST,
    output_type=flyteidl_dot_admin_dot_signal__pb2._SIGNALLISTREQUEST,
    serialized_options=_b('\202\323\344\223\002m\022k/api/v1/signals/{workflow_execution_id.project}/{workflow_execution_id.domain}/{workflow_execution_id.name}\222AI\032GFetch existing signal definitions matching the input signal id filters.'),
  ),
  _descriptor.MethodDescriptor(
    name='SetSignal',
    full_name='flyteidl.service.SignalService.SetSignal',
    index=2,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_signal__pb2._SIGNALSETREQUEST,
    output_type=flyteidl_dot_admin_dot_signal__pb2._SIGNALSETRESPONSE,
    serialized_options=_b('\202\323\344\223\002\024\"\017/api/v1/signals:\001*\222A\300\001\032\023Set a signal value.JB\n\003400\022;\n9Returned for bad request that may have failed validation.Je\n\003409\022^\n\\Returned for a request that references an identical entity that has already been registered.'),
  ),
])
_sym_db.RegisterServiceDescriptor(_SIGNALSERVICE)

DESCRIPTOR.services_by_name['SignalService'] = _SIGNALSERVICE

# @@protoc_insertion_point(module_scope)
