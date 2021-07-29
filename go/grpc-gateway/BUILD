package(default_visibility = ["//visibility:public"])
load("@build_stack_rules_proto//github.com/grpc-ecosystem/grpc-gateway:gateway_grpc_compile.bzl", "gateway_grpc_compile")

proto_library(
    name = "echo_service_proto",
    srcs = glob(
        ["*.proto"],
    ),
    deps = [
      "@com_google_googleapis//google/api:annotations_proto",
    ],
)


gateway_grpc_compile(
    name = "echo_service_grpc",
    deps = [":echo_service_proto"],
)

