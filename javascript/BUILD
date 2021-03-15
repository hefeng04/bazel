
load("@build_bazel_rules_nodejs//:index.bzl", "js_library", "nodejs_binary")
load("@io_bazel_rules_docker//nodejs:image.bzl", "nodejs_image")

js_library(
    name = "server_lib",
    srcs = [ "server.js" ],
    deps = [
        "@npm//koa",
        "@npm//koa-bodyparser",
        "@npm//koa-router",
        "@npm//koa-logger",
        "@npm//koa-cors",
        "@npm//koa-static",
        "@npm//koa-mount",
        "@npm//http-status",
    ],
)

nodejs_binary(
    name = "server",
    # Every execution of the program gets these arguments added
    args = ["--node_options=--expose-gc"],
    data = [
        ":server_lib",
    ],
    entry_point = ":server.js",
)

nodejs_image(
    name = "docker",
    binary = ":server",
)
