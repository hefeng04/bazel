load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "gateway-lib",
    srcs = ["main.go"],
    importpath = "go/grpc-gateway/",
    visibility = ["//visibility:private"],
    deps = [
	"//go/grpc-gateway/echo:echo-server",
	"//go/grpc-gateway/gateway",
        "@com_github_golang_glog//:glog",
    ],
)

go_binary(
    name = "grpc-gateway",
    embed = [":gateway-lib"],
    visibility = ["//visibility:public"],
)
