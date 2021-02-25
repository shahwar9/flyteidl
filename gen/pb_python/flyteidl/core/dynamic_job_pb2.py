# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: flyteidl/core/dynamic_job.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from flyteidl.core import tasks_pb2 as flyteidl_dot_core_dot_tasks__pb2
from flyteidl.core import workflow_pb2 as flyteidl_dot_core_dot_workflow__pb2
from flyteidl.core import literals_pb2 as flyteidl_dot_core_dot_literals__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='flyteidl/core/dynamic_job.proto',
  package='flyteidl.core',
  syntax='proto3',
  serialized_options=_b('Z0github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/core'),
  serialized_pb=_b('\n\x1f\x66lyteidl/core/dynamic_job.proto\x12\rflyteidl.core\x1a\x19\x66lyteidl/core/tasks.proto\x1a\x1c\x66lyteidl/core/workflow.proto\x1a\x1c\x66lyteidl/core/literals.proto\"\xd7\x01\n\x0e\x44ynamicJobSpec\x12\"\n\x05nodes\x18\x01 \x03(\x0b\x32\x13.flyteidl.core.Node\x12\x15\n\rmin_successes\x18\x02 \x01(\x03\x12\'\n\x07outputs\x18\x03 \x03(\x0b\x32\x16.flyteidl.core.Binding\x12*\n\x05tasks\x18\x04 \x03(\x0b\x32\x1b.flyteidl.core.TaskTemplate\x12\x35\n\x0csubworkflows\x18\x05 \x03(\x0b\x32\x1f.flyteidl.core.WorkflowTemplateB2Z0github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/coreb\x06proto3')
  ,
  dependencies=[flyteidl_dot_core_dot_tasks__pb2.DESCRIPTOR,flyteidl_dot_core_dot_workflow__pb2.DESCRIPTOR,flyteidl_dot_core_dot_literals__pb2.DESCRIPTOR,])




_DYNAMICJOBSPEC = _descriptor.Descriptor(
  name='DynamicJobSpec',
  full_name='flyteidl.core.DynamicJobSpec',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='nodes', full_name='flyteidl.core.DynamicJobSpec.nodes', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='min_successes', full_name='flyteidl.core.DynamicJobSpec.min_successes', index=1,
      number=2, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='outputs', full_name='flyteidl.core.DynamicJobSpec.outputs', index=2,
      number=3, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='tasks', full_name='flyteidl.core.DynamicJobSpec.tasks', index=3,
      number=4, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='subworkflows', full_name='flyteidl.core.DynamicJobSpec.subworkflows', index=4,
      number=5, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=138,
  serialized_end=353,
)

_DYNAMICJOBSPEC.fields_by_name['nodes'].message_type = flyteidl_dot_core_dot_workflow__pb2._NODE
_DYNAMICJOBSPEC.fields_by_name['outputs'].message_type = flyteidl_dot_core_dot_literals__pb2._BINDING
_DYNAMICJOBSPEC.fields_by_name['tasks'].message_type = flyteidl_dot_core_dot_tasks__pb2._TASKTEMPLATE
_DYNAMICJOBSPEC.fields_by_name['subworkflows'].message_type = flyteidl_dot_core_dot_workflow__pb2._WORKFLOWTEMPLATE
DESCRIPTOR.message_types_by_name['DynamicJobSpec'] = _DYNAMICJOBSPEC
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

DynamicJobSpec = _reflection.GeneratedProtocolMessageType('DynamicJobSpec', (_message.Message,), dict(
  DESCRIPTOR = _DYNAMICJOBSPEC,
  __module__ = 'flyteidl.core.dynamic_job_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.core.DynamicJobSpec)
  ))
_sym_db.RegisterMessage(DynamicJobSpec)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
