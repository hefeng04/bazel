package(default_visibility = ["//visibility:public"])

load(
    "@com_google_protobuf//:protobuf.bzl",
    "py_proto_library",
)


proto_library(
    name = "event_protos",
    srcs = glob(
        ["*.proto"],
    ),
)

py_proto_library(
    name = "event_py_protos",
    srcs = glob(
        ["*.proto"],
    ),
    tags = [
        "manual",
    ],
)
