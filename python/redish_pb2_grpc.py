# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

import redish_pb2 as redish__pb2


class RedishStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.get = channel.unary_unary(
        '/redish.Redish/get',
        request_serializer=redish__pb2.Key.SerializeToString,
        response_deserializer=redish__pb2.SingleValue.FromString,
        )
    self.set = channel.unary_unary(
        '/redish.Redish/set',
        request_serializer=redish__pb2.SetRequest.SerializeToString,
        response_deserializer=redish__pb2.OK.FromString,
        )
    self.dele = channel.unary_unary(
        '/redish.Redish/dele',
        request_serializer=redish__pb2.KeyList.SerializeToString,
        response_deserializer=redish__pb2.IntValue.FromString,
        )
    self.exists = channel.unary_unary(
        '/redish.Redish/exists',
        request_serializer=redish__pb2.KeyList.SerializeToString,
        response_deserializer=redish__pb2.IntValue.FromString,
        )
    self.incr = channel.unary_unary(
        '/redish.Redish/incr',
        request_serializer=redish__pb2.Key.SerializeToString,
        response_deserializer=redish__pb2.IntValue.FromString,
        )
    self.decr = channel.unary_unary(
        '/redish.Redish/decr',
        request_serializer=redish__pb2.Key.SerializeToString,
        response_deserializer=redish__pb2.IntValue.FromString,
        )
    self.incrby = channel.unary_unary(
        '/redish.Redish/incrby',
        request_serializer=redish__pb2.KeyValue.SerializeToString,
        response_deserializer=redish__pb2.IntValue.FromString,
        )
    self.decrby = channel.unary_unary(
        '/redish.Redish/decrby',
        request_serializer=redish__pb2.KeyValue.SerializeToString,
        response_deserializer=redish__pb2.IntValue.FromString,
        )
    self.strlen = channel.unary_unary(
        '/redish.Redish/strlen',
        request_serializer=redish__pb2.Key.SerializeToString,
        response_deserializer=redish__pb2.IntValue.FromString,
        )
    self.getset = channel.unary_unary(
        '/redish.Redish/getset',
        request_serializer=redish__pb2.KeyValue.SerializeToString,
        response_deserializer=redish__pb2.SingleValue.FromString,
        )
    self.mget = channel.unary_unary(
        '/redish.Redish/mget',
        request_serializer=redish__pb2.KeyList.SerializeToString,
        response_deserializer=redish__pb2.ValueList.FromString,
        )
    self.mset = channel.unary_unary(
        '/redish.Redish/mset',
        request_serializer=redish__pb2.KeyValueList.SerializeToString,
        response_deserializer=redish__pb2.OK.FromString,
        )
    self.type = channel.unary_unary(
        '/redish.Redish/type',
        request_serializer=redish__pb2.Key.SerializeToString,
        response_deserializer=redish__pb2.SingleValue.FromString,
        )


class RedishServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def get(self, request, context):
    """https://redis.io/commands
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def set(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def dele(self, request, context):
    """has to be dele not del because python del is a keyword
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def exists(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def incr(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def decr(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def incrby(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def decrby(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def strlen(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def getset(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def mget(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def mset(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def type(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_RedishServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'get': grpc.unary_unary_rpc_method_handler(
          servicer.get,
          request_deserializer=redish__pb2.Key.FromString,
          response_serializer=redish__pb2.SingleValue.SerializeToString,
      ),
      'set': grpc.unary_unary_rpc_method_handler(
          servicer.set,
          request_deserializer=redish__pb2.SetRequest.FromString,
          response_serializer=redish__pb2.OK.SerializeToString,
      ),
      'dele': grpc.unary_unary_rpc_method_handler(
          servicer.dele,
          request_deserializer=redish__pb2.KeyList.FromString,
          response_serializer=redish__pb2.IntValue.SerializeToString,
      ),
      'exists': grpc.unary_unary_rpc_method_handler(
          servicer.exists,
          request_deserializer=redish__pb2.KeyList.FromString,
          response_serializer=redish__pb2.IntValue.SerializeToString,
      ),
      'incr': grpc.unary_unary_rpc_method_handler(
          servicer.incr,
          request_deserializer=redish__pb2.Key.FromString,
          response_serializer=redish__pb2.IntValue.SerializeToString,
      ),
      'decr': grpc.unary_unary_rpc_method_handler(
          servicer.decr,
          request_deserializer=redish__pb2.Key.FromString,
          response_serializer=redish__pb2.IntValue.SerializeToString,
      ),
      'incrby': grpc.unary_unary_rpc_method_handler(
          servicer.incrby,
          request_deserializer=redish__pb2.KeyValue.FromString,
          response_serializer=redish__pb2.IntValue.SerializeToString,
      ),
      'decrby': grpc.unary_unary_rpc_method_handler(
          servicer.decrby,
          request_deserializer=redish__pb2.KeyValue.FromString,
          response_serializer=redish__pb2.IntValue.SerializeToString,
      ),
      'strlen': grpc.unary_unary_rpc_method_handler(
          servicer.strlen,
          request_deserializer=redish__pb2.Key.FromString,
          response_serializer=redish__pb2.IntValue.SerializeToString,
      ),
      'getset': grpc.unary_unary_rpc_method_handler(
          servicer.getset,
          request_deserializer=redish__pb2.KeyValue.FromString,
          response_serializer=redish__pb2.SingleValue.SerializeToString,
      ),
      'mget': grpc.unary_unary_rpc_method_handler(
          servicer.mget,
          request_deserializer=redish__pb2.KeyList.FromString,
          response_serializer=redish__pb2.ValueList.SerializeToString,
      ),
      'mset': grpc.unary_unary_rpc_method_handler(
          servicer.mset,
          request_deserializer=redish__pb2.KeyValueList.FromString,
          response_serializer=redish__pb2.OK.SerializeToString,
      ),
      'type': grpc.unary_unary_rpc_method_handler(
          servicer.type,
          request_deserializer=redish__pb2.Key.FromString,
          response_serializer=redish__pb2.SingleValue.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'redish.Redish', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
