load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "gateway",
    srcs = [
        "gateway.go",
        "handlers.go",
        "main.go",
    ],
    importpath = "go/grpc-gateway/gateway",
    visibility = ["//visibility:public"],
    deps = [
        "//go/grpc-gateway/echo:echoservice_go_proto",
        "@com_github_golang_glog//:glog",
        "@go_googleapis//google/rpc:errdetails_go_proto",
        "@grpc_ecosystem_grpc_gateway//runtime",
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_google_grpc//connectivity",
    ],
)

alias(
    name = "go_default_library",
    actual = ":gateway",
    visibility = ["//examples:__subpackages__"],
)
