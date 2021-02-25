# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: flyteidl/plugins/spark.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='flyteidl/plugins/spark.proto',
  package='flyteidl.plugins',
  syntax='proto3',
  serialized_options=_b('Z3github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/plugins'),
  serialized_pb=_b('\n\x1c\x66lyteidl/plugins/spark.proto\x12\x10\x66lyteidl.plugins\"B\n\x10SparkApplication\".\n\x04Type\x12\n\n\x06PYTHON\x10\x00\x12\x08\n\x04JAVA\x10\x01\x12\t\n\x05SCALA\x10\x02\x12\x05\n\x01R\x10\x03\"\xf5\x02\n\x08SparkJob\x12@\n\x0f\x61pplicationType\x18\x01 \x01(\x0e\x32\'.flyteidl.plugins.SparkApplication.Type\x12\x1b\n\x13mainApplicationFile\x18\x02 \x01(\t\x12\x11\n\tmainClass\x18\x03 \x01(\t\x12<\n\tsparkConf\x18\x04 \x03(\x0b\x32).flyteidl.plugins.SparkJob.SparkConfEntry\x12>\n\nhadoopConf\x18\x05 \x03(\x0b\x32*.flyteidl.plugins.SparkJob.HadoopConfEntry\x12\x14\n\x0c\x65xecutorPath\x18\x06 \x01(\t\x1a\x30\n\x0eSparkConfEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\x1a\x31\n\x0fHadoopConfEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\x42\x35Z3github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/pluginsb\x06proto3')
)



_SPARKAPPLICATION_TYPE = _descriptor.EnumDescriptor(
  name='Type',
  full_name='flyteidl.plugins.SparkApplication.Type',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='PYTHON', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='JAVA', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SCALA', index=2, number=2,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='R', index=3, number=3,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=70,
  serialized_end=116,
)
_sym_db.RegisterEnumDescriptor(_SPARKAPPLICATION_TYPE)


_SPARKAPPLICATION = _descriptor.Descriptor(
  name='SparkApplication',
  full_name='flyteidl.plugins.SparkApplication',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _SPARKAPPLICATION_TYPE,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=50,
  serialized_end=116,
)


_SPARKJOB_SPARKCONFENTRY = _descriptor.Descriptor(
  name='SparkConfEntry',
  full_name='flyteidl.plugins.SparkJob.SparkConfEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='flyteidl.plugins.SparkJob.SparkConfEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='flyteidl.plugins.SparkJob.SparkConfEntry.value', index=1,
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
  serialized_options=_b('8\001'),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=393,
  serialized_end=441,
)

_SPARKJOB_HADOOPCONFENTRY = _descriptor.Descriptor(
  name='HadoopConfEntry',
  full_name='flyteidl.plugins.SparkJob.HadoopConfEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='flyteidl.plugins.SparkJob.HadoopConfEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='flyteidl.plugins.SparkJob.HadoopConfEntry.value', index=1,
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
  serialized_options=_b('8\001'),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=443,
  serialized_end=492,
)

_SPARKJOB = _descriptor.Descriptor(
  name='SparkJob',
  full_name='flyteidl.plugins.SparkJob',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='applicationType', full_name='flyteidl.plugins.SparkJob.applicationType', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='mainApplicationFile', full_name='flyteidl.plugins.SparkJob.mainApplicationFile', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='mainClass', full_name='flyteidl.plugins.SparkJob.mainClass', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='sparkConf', full_name='flyteidl.plugins.SparkJob.sparkConf', index=3,
      number=4, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='hadoopConf', full_name='flyteidl.plugins.SparkJob.hadoopConf', index=4,
      number=5, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='executorPath', full_name='flyteidl.plugins.SparkJob.executorPath', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_SPARKJOB_SPARKCONFENTRY, _SPARKJOB_HADOOPCONFENTRY, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=119,
  serialized_end=492,
)

_SPARKAPPLICATION_TYPE.containing_type = _SPARKAPPLICATION
_SPARKJOB_SPARKCONFENTRY.containing_type = _SPARKJOB
_SPARKJOB_HADOOPCONFENTRY.containing_type = _SPARKJOB
_SPARKJOB.fields_by_name['applicationType'].enum_type = _SPARKAPPLICATION_TYPE
_SPARKJOB.fields_by_name['sparkConf'].message_type = _SPARKJOB_SPARKCONFENTRY
_SPARKJOB.fields_by_name['hadoopConf'].message_type = _SPARKJOB_HADOOPCONFENTRY
DESCRIPTOR.message_types_by_name['SparkApplication'] = _SPARKAPPLICATION
DESCRIPTOR.message_types_by_name['SparkJob'] = _SPARKJOB
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

SparkApplication = _reflection.GeneratedProtocolMessageType('SparkApplication', (_message.Message,), dict(
  DESCRIPTOR = _SPARKAPPLICATION,
  __module__ = 'flyteidl.plugins.spark_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.plugins.SparkApplication)
  ))
_sym_db.RegisterMessage(SparkApplication)

SparkJob = _reflection.GeneratedProtocolMessageType('SparkJob', (_message.Message,), dict(

  SparkConfEntry = _reflection.GeneratedProtocolMessageType('SparkConfEntry', (_message.Message,), dict(
    DESCRIPTOR = _SPARKJOB_SPARKCONFENTRY,
    __module__ = 'flyteidl.plugins.spark_pb2'
    # @@protoc_insertion_point(class_scope:flyteidl.plugins.SparkJob.SparkConfEntry)
    ))
  ,

  HadoopConfEntry = _reflection.GeneratedProtocolMessageType('HadoopConfEntry', (_message.Message,), dict(
    DESCRIPTOR = _SPARKJOB_HADOOPCONFENTRY,
    __module__ = 'flyteidl.plugins.spark_pb2'
    # @@protoc_insertion_point(class_scope:flyteidl.plugins.SparkJob.HadoopConfEntry)
    ))
  ,
  DESCRIPTOR = _SPARKJOB,
  __module__ = 'flyteidl.plugins.spark_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.plugins.SparkJob)
  ))
_sym_db.RegisterMessage(SparkJob)
_sym_db.RegisterMessage(SparkJob.SparkConfEntry)
_sym_db.RegisterMessage(SparkJob.HadoopConfEntry)


DESCRIPTOR._options = None
_SPARKJOB_SPARKCONFENTRY._options = None
_SPARKJOB_HADOOPCONFENTRY._options = None
# @@protoc_insertion_point(module_scope)
