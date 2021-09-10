package(default_visibility = ["//visibility:public"])

load(
    "@com_google_protobuf//:protobuf.bzl",
    "py_proto_library",
)


proto_library(
    name = "api_protos",
    srcs = glob(
        ["*.proto"],
    ),
    deps = [
        "//python:python_protos",
    ],
)

py_proto_library(
    name = "api_py_protos",
    srcs = glob(
        ["*.proto"],
    ),
    deps = [
        "//python:python_py_protos",
    ],
)
