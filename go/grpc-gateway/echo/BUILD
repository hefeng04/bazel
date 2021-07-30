load("@rules_proto//proto:defs.bzl", "proto_library")

package(default_visibility = ["//visibility:public"])

load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")
load("@grpc_ecosystem_grpc_gateway//protoc-gen-openapiv2:defs.bzl", "protoc_gen_openapiv2")


go_proto_library(
    name = "echoservice_go_proto",
    compilers = [
		    "@io_bazel_rules_go//proto:go_grpc",
	"@grpc_ecosystem_grpc_gateway//protoc-gen-grpc-gateway:go_gen_grpc_gateway",
		    ],
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

go_library(
    name = "grpc-gateway",
    srcs = ["echo.go"],
    importpath = "github.com/example/project/go/grpc-gateway/echo",
    deps = [
        ":echoservice_go_proto",
        "@com_github_golang_glog//:glog",
        "@grpc_ecosystem_grpc_gateway//runtime",
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_google_grpc//metadata",
    ],
)
