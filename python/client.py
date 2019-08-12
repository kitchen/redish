import grpc
import redish_pb2 as redish
import redish_pb2_grpc as redish_grpc

channel = grpc.insecure_channel('localhost:4242')
stub = redish_grpc.RedishStub(channel)

fookey = redish.Key(key = "foo")
barkey = redish.Key(key = "bar")

setfoo = redish.SetRequest(key = "foo", value = "foo")
keys = redish.KeyList()
keys.keys.add().key = "foo"
keys.keys.add().key = "bar"
keys.keys.add().key = "baz"

import code; code.interact(local=dict(globals(), **locals()))
