# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/chat_controll.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x19proto/chat_controll.proto\x12\x0f\x63hat_controller\"M\n\x08\x43hatData\x12*\n\x03msg\x18\x01 \x03(\x0b\x32\x1d.chat_controller.QueryMessage\x12\x15\n\rauditStrategy\x18\x02 \x01(\t\"j\n\x0b\x43hatRequest\x12\x0f\n\x07traceId\x18\x01 \x01(\t\x12\x0e\n\x06userId\x18\x02 \x01(\t\x12\x11\n\tsessionId\x18\x03 \x01(\t\x12\'\n\x04\x64\x61ta\x18\x04 \x01(\x0b\x32\x19.chat_controller.ChatData\"\x1d\n\nChatAnswer\x12\x0f\n\x07\x63ontent\x18\x01 \x01(\t\"b\n\tChatReply\x12\x0f\n\x07traceId\x18\x01 \x01(\t\x12\x0c\n\x04\x63ode\x18\x02 \x01(\x05\x12\x0b\n\x03msg\x18\x03 \x01(\t\x12)\n\x04\x64\x61ta\x18\x04 \x01(\x0b\x32\x1b.chat_controller.ChatAnswer\"s\n\x0cQueryMessage\x12\x30\n\x04role\x18\x01 \x01(\x0e\x32\".chat_controller.QueryMessage.Role\x12\x0f\n\x07\x63ontent\x18\x02 \x01(\t\" \n\x04Role\x12\t\n\x05HUMAN\x10\x00\x12\r\n\tASSISTANT\x10\x01\x32T\n\x0c\x43hatControll\x12\x44\n\x04\x43hat\x12\x1c.chat_controller.ChatRequest\x1a\x1a.chat_controller.ChatReply\"\x00\x30\x01\x42\x12Z\x10./;chat_controllb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'proto.chat_controll_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\020./;chat_controll'
  _globals['_CHATDATA']._serialized_start=46
  _globals['_CHATDATA']._serialized_end=123
  _globals['_CHATREQUEST']._serialized_start=125
  _globals['_CHATREQUEST']._serialized_end=231
  _globals['_CHATANSWER']._serialized_start=233
  _globals['_CHATANSWER']._serialized_end=262
  _globals['_CHATREPLY']._serialized_start=264
  _globals['_CHATREPLY']._serialized_end=362
  _globals['_QUERYMESSAGE']._serialized_start=364
  _globals['_QUERYMESSAGE']._serialized_end=479
  _globals['_QUERYMESSAGE_ROLE']._serialized_start=447
  _globals['_QUERYMESSAGE_ROLE']._serialized_end=479
  _globals['_CHATCONTROLL']._serialized_start=481
  _globals['_CHATCONTROLL']._serialized_end=565
# @@protoc_insertion_point(module_scope)