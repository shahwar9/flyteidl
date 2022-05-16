# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: flyteidl/admin/workflow.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from flyteidl.core import compiler_pb2 as flyteidl_dot_core_dot_compiler__pb2
from flyteidl.core import identifier_pb2 as flyteidl_dot_core_dot_identifier__pb2
from flyteidl.core import workflow_pb2 as flyteidl_dot_core_dot_workflow__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
from flyteidl.admin import entity_description_pb2 as flyteidl_dot_admin_dot_entity__description__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='flyteidl/admin/workflow.proto',
  package='flyteidl.admin',
  syntax='proto3',
  serialized_options=_b('Z5github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/admin'),
  serialized_pb=_b('\n\x1d\x66lyteidl/admin/workflow.proto\x12\x0e\x66lyteidl.admin\x1a\x1c\x66lyteidl/core/compiler.proto\x1a\x1e\x66lyteidl/core/identifier.proto\x1a\x1c\x66lyteidl/core/workflow.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\'flyteidl/admin/entity_description.proto\"j\n\x15WorkflowCreateRequest\x12%\n\x02id\x18\x01 \x01(\x0b\x32\x19.flyteidl.core.Identifier\x12*\n\x04spec\x18\x02 \x01(\x0b\x32\x1c.flyteidl.admin.WorkflowSpec\"\x18\n\x16WorkflowCreateResponse\"c\n\x08Workflow\x12%\n\x02id\x18\x01 \x01(\x0b\x32\x19.flyteidl.core.Identifier\x12\x30\n\x07\x63losure\x18\x02 \x01(\x0b\x32\x1f.flyteidl.admin.WorkflowClosure\"J\n\x0cWorkflowList\x12+\n\tworkflows\x18\x01 \x03(\x0b\x32\x18.flyteidl.admin.Workflow\x12\r\n\x05token\x18\x02 \x01(\t\"\xb8\x01\n\x0cWorkflowSpec\x12\x31\n\x08template\x18\x01 \x01(\x0b\x32\x1f.flyteidl.core.WorkflowTemplate\x12\x36\n\rsub_workflows\x18\x02 \x03(\x0b\x32\x1f.flyteidl.core.WorkflowTemplate\x12=\n\x12\x65ntity_description\x18\x03 \x01(\x0b\x32!.flyteidl.admin.EntityDescription\"\x84\x01\n\x0fWorkflowClosure\x12\x41\n\x11\x63ompiled_workflow\x18\x01 \x01(\x0b\x32&.flyteidl.core.CompiledWorkflowClosure\x12.\n\ncreated_at\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.TimestampB7Z5github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/adminb\x06proto3')
  ,
  dependencies=[flyteidl_dot_core_dot_compiler__pb2.DESCRIPTOR,flyteidl_dot_core_dot_identifier__pb2.DESCRIPTOR,flyteidl_dot_core_dot_workflow__pb2.DESCRIPTOR,google_dot_protobuf_dot_timestamp__pb2.DESCRIPTOR,flyteidl_dot_admin_dot_entity__description__pb2.DESCRIPTOR,])




_WORKFLOWCREATEREQUEST = _descriptor.Descriptor(
  name='WorkflowCreateRequest',
  full_name='flyteidl.admin.WorkflowCreateRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='flyteidl.admin.WorkflowCreateRequest.id', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='spec', full_name='flyteidl.admin.WorkflowCreateRequest.spec', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
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
  serialized_start=215,
  serialized_end=321,
)


_WORKFLOWCREATERESPONSE = _descriptor.Descriptor(
  name='WorkflowCreateResponse',
  full_name='flyteidl.admin.WorkflowCreateResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
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
  serialized_start=323,
  serialized_end=347,
)


_WORKFLOW = _descriptor.Descriptor(
  name='Workflow',
  full_name='flyteidl.admin.Workflow',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='flyteidl.admin.Workflow.id', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='closure', full_name='flyteidl.admin.Workflow.closure', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
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
  serialized_start=349,
  serialized_end=448,
)


_WORKFLOWLIST = _descriptor.Descriptor(
  name='WorkflowList',
  full_name='flyteidl.admin.WorkflowList',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='workflows', full_name='flyteidl.admin.WorkflowList.workflows', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='token', full_name='flyteidl.admin.WorkflowList.token', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
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
  serialized_start=450,
  serialized_end=524,
)


_WORKFLOWSPEC = _descriptor.Descriptor(
  name='WorkflowSpec',
  full_name='flyteidl.admin.WorkflowSpec',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='template', full_name='flyteidl.admin.WorkflowSpec.template', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='sub_workflows', full_name='flyteidl.admin.WorkflowSpec.sub_workflows', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='entity_description', full_name='flyteidl.admin.WorkflowSpec.entity_description', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
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
  serialized_start=527,
  serialized_end=711,
)


_WORKFLOWCLOSURE = _descriptor.Descriptor(
  name='WorkflowClosure',
  full_name='flyteidl.admin.WorkflowClosure',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='compiled_workflow', full_name='flyteidl.admin.WorkflowClosure.compiled_workflow', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='created_at', full_name='flyteidl.admin.WorkflowClosure.created_at', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
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
  serialized_start=714,
  serialized_end=846,
)

_WORKFLOWCREATEREQUEST.fields_by_name['id'].message_type = flyteidl_dot_core_dot_identifier__pb2._IDENTIFIER
_WORKFLOWCREATEREQUEST.fields_by_name['spec'].message_type = _WORKFLOWSPEC
_WORKFLOW.fields_by_name['id'].message_type = flyteidl_dot_core_dot_identifier__pb2._IDENTIFIER
_WORKFLOW.fields_by_name['closure'].message_type = _WORKFLOWCLOSURE
_WORKFLOWLIST.fields_by_name['workflows'].message_type = _WORKFLOW
_WORKFLOWSPEC.fields_by_name['template'].message_type = flyteidl_dot_core_dot_workflow__pb2._WORKFLOWTEMPLATE
_WORKFLOWSPEC.fields_by_name['sub_workflows'].message_type = flyteidl_dot_core_dot_workflow__pb2._WORKFLOWTEMPLATE
_WORKFLOWSPEC.fields_by_name['entity_description'].message_type = flyteidl_dot_admin_dot_entity__description__pb2._ENTITYDESCRIPTION
_WORKFLOWCLOSURE.fields_by_name['compiled_workflow'].message_type = flyteidl_dot_core_dot_compiler__pb2._COMPILEDWORKFLOWCLOSURE
_WORKFLOWCLOSURE.fields_by_name['created_at'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
DESCRIPTOR.message_types_by_name['WorkflowCreateRequest'] = _WORKFLOWCREATEREQUEST
DESCRIPTOR.message_types_by_name['WorkflowCreateResponse'] = _WORKFLOWCREATERESPONSE
DESCRIPTOR.message_types_by_name['Workflow'] = _WORKFLOW
DESCRIPTOR.message_types_by_name['WorkflowList'] = _WORKFLOWLIST
DESCRIPTOR.message_types_by_name['WorkflowSpec'] = _WORKFLOWSPEC
DESCRIPTOR.message_types_by_name['WorkflowClosure'] = _WORKFLOWCLOSURE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

WorkflowCreateRequest = _reflection.GeneratedProtocolMessageType('WorkflowCreateRequest', (_message.Message,), dict(
  DESCRIPTOR = _WORKFLOWCREATEREQUEST,
  __module__ = 'flyteidl.admin.workflow_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.admin.WorkflowCreateRequest)
  ))
_sym_db.RegisterMessage(WorkflowCreateRequest)

WorkflowCreateResponse = _reflection.GeneratedProtocolMessageType('WorkflowCreateResponse', (_message.Message,), dict(
  DESCRIPTOR = _WORKFLOWCREATERESPONSE,
  __module__ = 'flyteidl.admin.workflow_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.admin.WorkflowCreateResponse)
  ))
_sym_db.RegisterMessage(WorkflowCreateResponse)

Workflow = _reflection.GeneratedProtocolMessageType('Workflow', (_message.Message,), dict(
  DESCRIPTOR = _WORKFLOW,
  __module__ = 'flyteidl.admin.workflow_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.admin.Workflow)
  ))
_sym_db.RegisterMessage(Workflow)

WorkflowList = _reflection.GeneratedProtocolMessageType('WorkflowList', (_message.Message,), dict(
  DESCRIPTOR = _WORKFLOWLIST,
  __module__ = 'flyteidl.admin.workflow_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.admin.WorkflowList)
  ))
_sym_db.RegisterMessage(WorkflowList)

WorkflowSpec = _reflection.GeneratedProtocolMessageType('WorkflowSpec', (_message.Message,), dict(
  DESCRIPTOR = _WORKFLOWSPEC,
  __module__ = 'flyteidl.admin.workflow_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.admin.WorkflowSpec)
  ))
_sym_db.RegisterMessage(WorkflowSpec)

WorkflowClosure = _reflection.GeneratedProtocolMessageType('WorkflowClosure', (_message.Message,), dict(
  DESCRIPTOR = _WORKFLOWCLOSURE,
  __module__ = 'flyteidl.admin.workflow_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.admin.WorkflowClosure)
  ))
_sym_db.RegisterMessage(WorkflowClosure)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
