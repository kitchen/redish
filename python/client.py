import grpc
import redish_pb2
import redish_pb2_grpc

channel = grpc.insecure_channel('localhost:4242')
stub = redish_pb2_grpc.RedishStub(channel)

import code; code.interact(local=dict(globals(), **locals()))
