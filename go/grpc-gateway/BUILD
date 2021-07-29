load("@rules_proto//proto:defs.bzl", "proto_library")

package(default_visibility = ["//visibility:public"])

load("@build_stack_rules_proto//github.com/grpc-ecosystem/grpc-gateway:gateway_grpc_compile.bzl", "gateway_grpc_compile")
load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

go_library(
    name = "echo_lib",
    srcs = ["echo.go"],
    importpath = "go/grpcgateway/echo",
    deps = [
        ":echo_service_go_proto",
        "@com_github_golang_glog//:glog",
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_google_grpc//metadata:go_default_library",
    ],
)

go_proto_library(
    name = "echo_service_go_proto",
    compilers = ["@io_bazel_rules_go//proto:go_grpc"],
    importpath = "go/grpcgateway/proto/echo",
    proto = ":echo_service_proto",
    visibility = ["//visibility:public"],
    deps = [
        "@com_google_googleapis//google/api:annotations_go_proto",
    ],
)

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

proto_library(
    name = "echoservice_proto",
    srcs = ["echo_service.proto"],
    deps = ["@go_googleapis//google/api:annotations_proto"],
)

go_proto_library(
    name = "echoservice_go_proto",
    compilers = ["@io_bazel_rules_go//proto:go_grpc"],
    importpath = "github.com/example/project/go/grpc-gateway",
    proto = ":echoservice_proto",
    deps = ["@go_googleapis//google/api:annotations_go_proto"],
)

go_library(
    name = "grpc-gateway",
    srcs = ["echo.go"],
    embed = [":echoservice_go_proto"],
    importpath = "github.com/example/project/go/grpc-gateway",
    deps = [
        "@com_github_golang_glog//:glog",
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_google_grpc//metadata",
    ],
)
