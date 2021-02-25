# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: flyteidl/admin/project_domain_attributes.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from flyteidl.admin import matchable_resource_pb2 as flyteidl_dot_admin_dot_matchable__resource__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='flyteidl/admin/project_domain_attributes.proto',
  package='flyteidl.admin',
  syntax='proto3',
  serialized_options=_b('Z1github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/admin'),
  serialized_pb=_b('\n.flyteidl/admin/project_domain_attributes.proto\x12\x0e\x66lyteidl.admin\x1a\'flyteidl/admin/matchable_resource.proto\"{\n\x17ProjectDomainAttributes\x12\x0f\n\x07project\x18\x01 \x01(\t\x12\x0e\n\x06\x64omain\x18\x02 \x01(\t\x12?\n\x13matching_attributes\x18\x03 \x01(\x0b\x32\".flyteidl.admin.MatchingAttributes\"c\n$ProjectDomainAttributesUpdateRequest\x12;\n\nattributes\x18\x01 \x01(\x0b\x32\'.flyteidl.admin.ProjectDomainAttributes\"\'\n%ProjectDomainAttributesUpdateResponse\"~\n!ProjectDomainAttributesGetRequest\x12\x0f\n\x07project\x18\x01 \x01(\t\x12\x0e\n\x06\x64omain\x18\x02 \x01(\t\x12\x38\n\rresource_type\x18\x03 \x01(\x0e\x32!.flyteidl.admin.MatchableResource\"a\n\"ProjectDomainAttributesGetResponse\x12;\n\nattributes\x18\x01 \x01(\x0b\x32\'.flyteidl.admin.ProjectDomainAttributes\"\x81\x01\n$ProjectDomainAttributesDeleteRequest\x12\x0f\n\x07project\x18\x01 \x01(\t\x12\x0e\n\x06\x64omain\x18\x02 \x01(\t\x12\x38\n\rresource_type\x18\x03 \x01(\x0e\x32!.flyteidl.admin.MatchableResource\"\'\n%ProjectDomainAttributesDeleteResponseB3Z1github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/adminb\x06proto3')
  ,
  dependencies=[flyteidl_dot_admin_dot_matchable__resource__pb2.DESCRIPTOR,])




_PROJECTDOMAINATTRIBUTES = _descriptor.Descriptor(
  name='ProjectDomainAttributes',
  full_name='flyteidl.admin.ProjectDomainAttributes',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='project', full_name='flyteidl.admin.ProjectDomainAttributes.project', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='domain', full_name='flyteidl.admin.ProjectDomainAttributes.domain', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='matching_attributes', full_name='flyteidl.admin.ProjectDomainAttributes.matching_attributes', index=2,
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
  serialized_start=107,
  serialized_end=230,
)


_PROJECTDOMAINATTRIBUTESUPDATEREQUEST = _descriptor.Descriptor(
  name='ProjectDomainAttributesUpdateRequest',
  full_name='flyteidl.admin.ProjectDomainAttributesUpdateRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='attributes', full_name='flyteidl.admin.ProjectDomainAttributesUpdateRequest.attributes', index=0,
      number=1, type=11, cpp_type=10, label=1,
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
  serialized_start=232,
  serialized_end=331,
)


_PROJECTDOMAINATTRIBUTESUPDATERESPONSE = _descriptor.Descriptor(
  name='ProjectDomainAttributesUpdateResponse',
  full_name='flyteidl.admin.ProjectDomainAttributesUpdateResponse',
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
  serialized_start=333,
  serialized_end=372,
)


_PROJECTDOMAINATTRIBUTESGETREQUEST = _descriptor.Descriptor(
  name='ProjectDomainAttributesGetRequest',
  full_name='flyteidl.admin.ProjectDomainAttributesGetRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='project', full_name='flyteidl.admin.ProjectDomainAttributesGetRequest.project', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='domain', full_name='flyteidl.admin.ProjectDomainAttributesGetRequest.domain', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='resource_type', full_name='flyteidl.admin.ProjectDomainAttributesGetRequest.resource_type', index=2,
      number=3, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
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
  serialized_start=374,
  serialized_end=500,
)


_PROJECTDOMAINATTRIBUTESGETRESPONSE = _descriptor.Descriptor(
  name='ProjectDomainAttributesGetResponse',
  full_name='flyteidl.admin.ProjectDomainAttributesGetResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='attributes', full_name='flyteidl.admin.ProjectDomainAttributesGetResponse.attributes', index=0,
      number=1, type=11, cpp_type=10, label=1,
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
  serialized_start=502,
  serialized_end=599,
)


_PROJECTDOMAINATTRIBUTESDELETEREQUEST = _descriptor.Descriptor(
  name='ProjectDomainAttributesDeleteRequest',
  full_name='flyteidl.admin.ProjectDomainAttributesDeleteRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='project', full_name='flyteidl.admin.ProjectDomainAttributesDeleteRequest.project', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='domain', full_name='flyteidl.admin.ProjectDomainAttributesDeleteRequest.domain', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='resource_type', full_name='flyteidl.admin.ProjectDomainAttributesDeleteRequest.resource_type', index=2,
      number=3, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
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
  serialized_start=602,
  serialized_end=731,
)


_PROJECTDOMAINATTRIBUTESDELETERESPONSE = _descriptor.Descriptor(
  name='ProjectDomainAttributesDeleteResponse',
  full_name='flyteidl.admin.ProjectDomainAttributesDeleteResponse',
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
  serialized_start=733,
  serialized_end=772,
)

_PROJECTDOMAINATTRIBUTES.fields_by_name['matching_attributes'].message_type = flyteidl_dot_admin_dot_matchable__resource__pb2._MATCHINGATTRIBUTES
_PROJECTDOMAINATTRIBUTESUPDATEREQUEST.fields_by_name['attributes'].message_type = _PROJECTDOMAINATTRIBUTES
_PROJECTDOMAINATTRIBUTESGETREQUEST.fields_by_name['resource_type'].enum_type = flyteidl_dot_admin_dot_matchable__resource__pb2._MATCHABLERESOURCE
_PROJECTDOMAINATTRIBUTESGETRESPONSE.fields_by_name['attributes'].message_type = _PROJECTDOMAINATTRIBUTES
_PROJECTDOMAINATTRIBUTESDELETEREQUEST.fields_by_name['resource_type'].enum_type = flyteidl_dot_admin_dot_matchable__resource__pb2._MATCHABLERESOURCE
DESCRIPTOR.message_types_by_name['ProjectDomainAttributes'] = _PROJECTDOMAINATTRIBUTES
DESCRIPTOR.message_types_by_name['ProjectDomainAttributesUpdateRequest'] = _PROJECTDOMAINATTRIBUTESUPDATEREQUEST
DESCRIPTOR.message_types_by_name['ProjectDomainAttributesUpdateResponse'] = _PROJECTDOMAINATTRIBUTESUPDATERESPONSE
DESCRIPTOR.message_types_by_name['ProjectDomainAttributesGetRequest'] = _PROJECTDOMAINATTRIBUTESGETREQUEST
DESCRIPTOR.message_types_by_name['ProjectDomainAttributesGetResponse'] = _PROJECTDOMAINATTRIBUTESGETRESPONSE
DESCRIPTOR.message_types_by_name['ProjectDomainAttributesDeleteRequest'] = _PROJECTDOMAINATTRIBUTESDELETEREQUEST
DESCRIPTOR.message_types_by_name['ProjectDomainAttributesDeleteResponse'] = _PROJECTDOMAINATTRIBUTESDELETERESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ProjectDomainAttributes = _reflection.GeneratedProtocolMessageType('ProjectDomainAttributes', (_message.Message,), dict(
  DESCRIPTOR = _PROJECTDOMAINATTRIBUTES,
  __module__ = 'flyteidl.admin.project_domain_attributes_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.admin.ProjectDomainAttributes)
  ))
_sym_db.RegisterMessage(ProjectDomainAttributes)

ProjectDomainAttributesUpdateRequest = _reflection.GeneratedProtocolMessageType('ProjectDomainAttributesUpdateRequest', (_message.Message,), dict(
  DESCRIPTOR = _PROJECTDOMAINATTRIBUTESUPDATEREQUEST,
  __module__ = 'flyteidl.admin.project_domain_attributes_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.admin.ProjectDomainAttributesUpdateRequest)
  ))
_sym_db.RegisterMessage(ProjectDomainAttributesUpdateRequest)

ProjectDomainAttributesUpdateResponse = _reflection.GeneratedProtocolMessageType('ProjectDomainAttributesUpdateResponse', (_message.Message,), dict(
  DESCRIPTOR = _PROJECTDOMAINATTRIBUTESUPDATERESPONSE,
  __module__ = 'flyteidl.admin.project_domain_attributes_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.admin.ProjectDomainAttributesUpdateResponse)
  ))
_sym_db.RegisterMessage(ProjectDomainAttributesUpdateResponse)

ProjectDomainAttributesGetRequest = _reflection.GeneratedProtocolMessageType('ProjectDomainAttributesGetRequest', (_message.Message,), dict(
  DESCRIPTOR = _PROJECTDOMAINATTRIBUTESGETREQUEST,
  __module__ = 'flyteidl.admin.project_domain_attributes_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.admin.ProjectDomainAttributesGetRequest)
  ))
_sym_db.RegisterMessage(ProjectDomainAttributesGetRequest)

ProjectDomainAttributesGetResponse = _reflection.GeneratedProtocolMessageType('ProjectDomainAttributesGetResponse', (_message.Message,), dict(
  DESCRIPTOR = _PROJECTDOMAINATTRIBUTESGETRESPONSE,
  __module__ = 'flyteidl.admin.project_domain_attributes_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.admin.ProjectDomainAttributesGetResponse)
  ))
_sym_db.RegisterMessage(ProjectDomainAttributesGetResponse)

ProjectDomainAttributesDeleteRequest = _reflection.GeneratedProtocolMessageType('ProjectDomainAttributesDeleteRequest', (_message.Message,), dict(
  DESCRIPTOR = _PROJECTDOMAINATTRIBUTESDELETEREQUEST,
  __module__ = 'flyteidl.admin.project_domain_attributes_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.admin.ProjectDomainAttributesDeleteRequest)
  ))
_sym_db.RegisterMessage(ProjectDomainAttributesDeleteRequest)

ProjectDomainAttributesDeleteResponse = _reflection.GeneratedProtocolMessageType('ProjectDomainAttributesDeleteResponse', (_message.Message,), dict(
  DESCRIPTOR = _PROJECTDOMAINATTRIBUTESDELETERESPONSE,
  __module__ = 'flyteidl.admin.project_domain_attributes_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.admin.ProjectDomainAttributesDeleteResponse)
  ))
_sym_db.RegisterMessage(ProjectDomainAttributesDeleteResponse)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
