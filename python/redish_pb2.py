# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: redish.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import wrappers_pb2 as google_dot_protobuf_dot_wrappers__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='redish.proto',
  package='redish',
  syntax='proto3',
  serialized_options=None,
  serialized_pb=_b('\n\x0credish.proto\x12\x06redish\x1a\x1egoogle/protobuf/wrappers.proto\"\x12\n\x03Key\x12\x0b\n\x03key\x18\x01 \x01(\t\"&\n\x08KeyValue\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t\"$\n\x07KeyList\x12\x19\n\x04keys\x18\x01 \x03(\x0b\x32\x0b.redish.Key\"/\n\x0cKeyValueList\x12\x1f\n\x05pairs\x18\x01 \x03(\x0b\x32\x10.redish.KeyValue\":\n\x0bSingleValue\x12+\n\x05value\x18\x01 \x01(\x0b\x32\x1c.google.protobuf.StringValue\"\x19\n\x08IntValue\x12\r\n\x05value\x18\x01 \x01(\x12\")\n\x0bKeyIntValue\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\x12\"0\n\tValueList\x12#\n\x06values\x18\x01 \x03(\x0b\x32\x13.redish.SingleValue\"(\n\nSetRequest\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t\"\x04\n\x02OK2\xa3\x07\n\x06Redish\x12)\n\x03get\x12\x0b.redish.Key\x1a\x13.redish.SingleValue\"\x00\x12\'\n\x03set\x12\x12.redish.SetRequest\x1a\n.redish.OK\"\x00\x12+\n\x04\x64\x65le\x12\x0f.redish.KeyList\x1a\x10.redish.IntValue\"\x00\x12-\n\x06\x65xists\x12\x0f.redish.KeyList\x1a\x10.redish.IntValue\"\x00\x12\'\n\x04incr\x12\x0b.redish.Key\x1a\x10.redish.IntValue\"\x00\x12\'\n\x04\x64\x65\x63r\x12\x0b.redish.Key\x1a\x10.redish.IntValue\"\x00\x12\x31\n\x06incrby\x12\x13.redish.KeyIntValue\x1a\x10.redish.IntValue\"\x00\x12\x31\n\x06\x64\x65\x63rby\x12\x13.redish.KeyIntValue\x1a\x10.redish.IntValue\"\x00\x12)\n\x06strlen\x12\x0b.redish.Key\x1a\x10.redish.IntValue\"\x00\x12\x31\n\x06getset\x12\x10.redish.KeyValue\x1a\x13.redish.SingleValue\"\x00\x12,\n\x04mget\x12\x0f.redish.KeyList\x1a\x11.redish.ValueList\"\x00\x12*\n\x04mset\x12\x14.redish.KeyValueList\x1a\n.redish.OK\"\x00\x12*\n\x04type\x12\x0b.redish.Key\x1a\x13.redish.SingleValue\"\x00\x12\x31\n\x06\x65xpire\x12\x13.redish.KeyIntValue\x1a\x10.redish.IntValue\"\x00\x12\x32\n\x07pexpire\x12\x13.redish.KeyIntValue\x1a\x10.redish.IntValue\"\x00\x12\x33\n\x08\x65xpireat\x12\x13.redish.KeyIntValue\x1a\x10.redish.IntValue\"\x00\x12\x34\n\tpexpireat\x12\x13.redish.KeyIntValue\x1a\x10.redish.IntValue\"\x00\x12*\n\x07persist\x12\x0b.redish.Key\x1a\x10.redish.IntValue\"\x00\x12&\n\x03ttl\x12\x0b.redish.Key\x1a\x10.redish.IntValue\"\x00\x12\'\n\x04pttl\x12\x0b.redish.Key\x1a\x10.redish.IntValue\"\x00\x62\x06proto3')
  ,
  dependencies=[google_dot_protobuf_dot_wrappers__pb2.DESCRIPTOR,])




_KEY = _descriptor.Descriptor(
  name='Key',
  full_name='redish.Key',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='redish.Key.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
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
  serialized_start=56,
  serialized_end=74,
)


_KEYVALUE = _descriptor.Descriptor(
  name='KeyValue',
  full_name='redish.KeyValue',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='redish.KeyValue.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='redish.KeyValue.value', index=1,
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
  serialized_start=76,
  serialized_end=114,
)


_KEYLIST = _descriptor.Descriptor(
  name='KeyList',
  full_name='redish.KeyList',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='keys', full_name='redish.KeyList.keys', index=0,
      number=1, type=11, cpp_type=10, label=3,
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
  serialized_start=116,
  serialized_end=152,
)


_KEYVALUELIST = _descriptor.Descriptor(
  name='KeyValueList',
  full_name='redish.KeyValueList',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='pairs', full_name='redish.KeyValueList.pairs', index=0,
      number=1, type=11, cpp_type=10, label=3,
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
  serialized_start=154,
  serialized_end=201,
)


_SINGLEVALUE = _descriptor.Descriptor(
  name='SingleValue',
  full_name='redish.SingleValue',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='value', full_name='redish.SingleValue.value', index=0,
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
  serialized_start=203,
  serialized_end=261,
)


_INTVALUE = _descriptor.Descriptor(
  name='IntValue',
  full_name='redish.IntValue',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='value', full_name='redish.IntValue.value', index=0,
      number=1, type=18, cpp_type=2, label=1,
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
  serialized_start=263,
  serialized_end=288,
)


_KEYINTVALUE = _descriptor.Descriptor(
  name='KeyIntValue',
  full_name='redish.KeyIntValue',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='redish.KeyIntValue.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='redish.KeyIntValue.value', index=1,
      number=2, type=18, cpp_type=2, label=1,
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
  serialized_start=290,
  serialized_end=331,
)


_VALUELIST = _descriptor.Descriptor(
  name='ValueList',
  full_name='redish.ValueList',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='values', full_name='redish.ValueList.values', index=0,
      number=1, type=11, cpp_type=10, label=3,
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
  serialized_start=333,
  serialized_end=381,
)


_SETREQUEST = _descriptor.Descriptor(
  name='SetRequest',
  full_name='redish.SetRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='redish.SetRequest.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='redish.SetRequest.value', index=1,
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
  serialized_start=383,
  serialized_end=423,
)


_OK = _descriptor.Descriptor(
  name='OK',
  full_name='redish.OK',
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
  serialized_start=425,
  serialized_end=429,
)

_KEYLIST.fields_by_name['keys'].message_type = _KEY
_KEYVALUELIST.fields_by_name['pairs'].message_type = _KEYVALUE
_SINGLEVALUE.fields_by_name['value'].message_type = google_dot_protobuf_dot_wrappers__pb2._STRINGVALUE
_VALUELIST.fields_by_name['values'].message_type = _SINGLEVALUE
DESCRIPTOR.message_types_by_name['Key'] = _KEY
DESCRIPTOR.message_types_by_name['KeyValue'] = _KEYVALUE
DESCRIPTOR.message_types_by_name['KeyList'] = _KEYLIST
DESCRIPTOR.message_types_by_name['KeyValueList'] = _KEYVALUELIST
DESCRIPTOR.message_types_by_name['SingleValue'] = _SINGLEVALUE
DESCRIPTOR.message_types_by_name['IntValue'] = _INTVALUE
DESCRIPTOR.message_types_by_name['KeyIntValue'] = _KEYINTVALUE
DESCRIPTOR.message_types_by_name['ValueList'] = _VALUELIST
DESCRIPTOR.message_types_by_name['SetRequest'] = _SETREQUEST
DESCRIPTOR.message_types_by_name['OK'] = _OK
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Key = _reflection.GeneratedProtocolMessageType('Key', (_message.Message,), {
  'DESCRIPTOR' : _KEY,
  '__module__' : 'redish_pb2'
  # @@protoc_insertion_point(class_scope:redish.Key)
  })
_sym_db.RegisterMessage(Key)

KeyValue = _reflection.GeneratedProtocolMessageType('KeyValue', (_message.Message,), {
  'DESCRIPTOR' : _KEYVALUE,
  '__module__' : 'redish_pb2'
  # @@protoc_insertion_point(class_scope:redish.KeyValue)
  })
_sym_db.RegisterMessage(KeyValue)

KeyList = _reflection.GeneratedProtocolMessageType('KeyList', (_message.Message,), {
  'DESCRIPTOR' : _KEYLIST,
  '__module__' : 'redish_pb2'
  # @@protoc_insertion_point(class_scope:redish.KeyList)
  })
_sym_db.RegisterMessage(KeyList)

KeyValueList = _reflection.GeneratedProtocolMessageType('KeyValueList', (_message.Message,), {
  'DESCRIPTOR' : _KEYVALUELIST,
  '__module__' : 'redish_pb2'
  # @@protoc_insertion_point(class_scope:redish.KeyValueList)
  })
_sym_db.RegisterMessage(KeyValueList)

SingleValue = _reflection.GeneratedProtocolMessageType('SingleValue', (_message.Message,), {
  'DESCRIPTOR' : _SINGLEVALUE,
  '__module__' : 'redish_pb2'
  # @@protoc_insertion_point(class_scope:redish.SingleValue)
  })
_sym_db.RegisterMessage(SingleValue)

IntValue = _reflection.GeneratedProtocolMessageType('IntValue', (_message.Message,), {
  'DESCRIPTOR' : _INTVALUE,
  '__module__' : 'redish_pb2'
  # @@protoc_insertion_point(class_scope:redish.IntValue)
  })
_sym_db.RegisterMessage(IntValue)

KeyIntValue = _reflection.GeneratedProtocolMessageType('KeyIntValue', (_message.Message,), {
  'DESCRIPTOR' : _KEYINTVALUE,
  '__module__' : 'redish_pb2'
  # @@protoc_insertion_point(class_scope:redish.KeyIntValue)
  })
_sym_db.RegisterMessage(KeyIntValue)

ValueList = _reflection.GeneratedProtocolMessageType('ValueList', (_message.Message,), {
  'DESCRIPTOR' : _VALUELIST,
  '__module__' : 'redish_pb2'
  # @@protoc_insertion_point(class_scope:redish.ValueList)
  })
_sym_db.RegisterMessage(ValueList)

SetRequest = _reflection.GeneratedProtocolMessageType('SetRequest', (_message.Message,), {
  'DESCRIPTOR' : _SETREQUEST,
  '__module__' : 'redish_pb2'
  # @@protoc_insertion_point(class_scope:redish.SetRequest)
  })
_sym_db.RegisterMessage(SetRequest)

OK = _reflection.GeneratedProtocolMessageType('OK', (_message.Message,), {
  'DESCRIPTOR' : _OK,
  '__module__' : 'redish_pb2'
  # @@protoc_insertion_point(class_scope:redish.OK)
  })
_sym_db.RegisterMessage(OK)



_REDISH = _descriptor.ServiceDescriptor(
  name='Redish',
  full_name='redish.Redish',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=432,
  serialized_end=1363,
  methods=[
  _descriptor.MethodDescriptor(
    name='get',
    full_name='redish.Redish.get',
    index=0,
    containing_service=None,
    input_type=_KEY,
    output_type=_SINGLEVALUE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='set',
    full_name='redish.Redish.set',
    index=1,
    containing_service=None,
    input_type=_SETREQUEST,
    output_type=_OK,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='dele',
    full_name='redish.Redish.dele',
    index=2,
    containing_service=None,
    input_type=_KEYLIST,
    output_type=_INTVALUE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='exists',
    full_name='redish.Redish.exists',
    index=3,
    containing_service=None,
    input_type=_KEYLIST,
    output_type=_INTVALUE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='incr',
    full_name='redish.Redish.incr',
    index=4,
    containing_service=None,
    input_type=_KEY,
    output_type=_INTVALUE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='decr',
    full_name='redish.Redish.decr',
    index=5,
    containing_service=None,
    input_type=_KEY,
    output_type=_INTVALUE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='incrby',
    full_name='redish.Redish.incrby',
    index=6,
    containing_service=None,
    input_type=_KEYINTVALUE,
    output_type=_INTVALUE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='decrby',
    full_name='redish.Redish.decrby',
    index=7,
    containing_service=None,
    input_type=_KEYINTVALUE,
    output_type=_INTVALUE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='strlen',
    full_name='redish.Redish.strlen',
    index=8,
    containing_service=None,
    input_type=_KEY,
    output_type=_INTVALUE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='getset',
    full_name='redish.Redish.getset',
    index=9,
    containing_service=None,
    input_type=_KEYVALUE,
    output_type=_SINGLEVALUE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='mget',
    full_name='redish.Redish.mget',
    index=10,
    containing_service=None,
    input_type=_KEYLIST,
    output_type=_VALUELIST,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='mset',
    full_name='redish.Redish.mset',
    index=11,
    containing_service=None,
    input_type=_KEYVALUELIST,
    output_type=_OK,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='type',
    full_name='redish.Redish.type',
    index=12,
    containing_service=None,
    input_type=_KEY,
    output_type=_SINGLEVALUE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='expire',
    full_name='redish.Redish.expire',
    index=13,
    containing_service=None,
    input_type=_KEYINTVALUE,
    output_type=_INTVALUE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='pexpire',
    full_name='redish.Redish.pexpire',
    index=14,
    containing_service=None,
    input_type=_KEYINTVALUE,
    output_type=_INTVALUE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='expireat',
    full_name='redish.Redish.expireat',
    index=15,
    containing_service=None,
    input_type=_KEYINTVALUE,
    output_type=_INTVALUE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='pexpireat',
    full_name='redish.Redish.pexpireat',
    index=16,
    containing_service=None,
    input_type=_KEYINTVALUE,
    output_type=_INTVALUE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='persist',
    full_name='redish.Redish.persist',
    index=17,
    containing_service=None,
    input_type=_KEY,
    output_type=_INTVALUE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='ttl',
    full_name='redish.Redish.ttl',
    index=18,
    containing_service=None,
    input_type=_KEY,
    output_type=_INTVALUE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='pttl',
    full_name='redish.Redish.pttl',
    index=19,
    containing_service=None,
    input_type=_KEY,
    output_type=_INTVALUE,
    serialized_options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_REDISH)

DESCRIPTOR.services_by_name['Redish'] = _REDISH

# @@protoc_insertion_point(module_scope)
