package(default_visibility = ["//visibility:public"])

load(
    "@com_google_protobuf//:protobuf.bzl",
    "py_proto_library",
)

py_binary(
  name = "try_proto",
  srcs = ["try_proto.py"],
  deps = [
    ":py_deps",
    "//proto/api:api_py_protos",
    "//proto/log:event_py_protos",
  ],
)

py_binary(
  name = "py_deps",
  srcs = ["py_deps.py"],
  deps = [
    "//proto/api:api_py_protos",
    "//proto/log:event_py_protos",
  ],
)

proto_library(
    name = "python_protos",
    srcs = glob(
        ["*.proto"],
    ),
)

py_proto_library(
    name = "python_py_protos",
    srcs = glob(
        ["*.proto"],
    ),
)
