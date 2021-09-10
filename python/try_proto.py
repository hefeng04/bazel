import proto.api.api_pb2 as api_pb
import proto.log.event_pb2 as event_pb
import python.py_deps as py_deps


print("cool")
py_deps.my_print()
api_instance = api_pb.Hello()
api_instance.message = "message"
print(api_instance)
