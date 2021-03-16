load("@build_bazel_rules_nodejs//:index.bzl", "js_library", "nodejs_binary")
load("@io_bazel_rules_docker//nodejs:image.bzl", "nodejs_image")

package_group(
    name = "javascript_visibility",
    packages = [
        "//javascript/...",
    ],
)

package(default_visibility = [":javascript_visibility"])

js_library(
    name = "server_lib",
    srcs = ["server.js"],
    deps = [
        "@npm//http-status",
        "@npm//koa",
        "@npm//koa-bodyparser",
        "@npm//koa-cors",
        "@npm//koa-logger",
        "@npm//koa-mount",
        "@npm//koa-router",
        "@npm//koa-static",
    ],
)

nodejs_binary(
    name = "server",
    # Every execution of the program gets these arguments added
    args = ["--node_options=--expose-gc"],
    data = [
        ":server_lib",
        "//javascript/portal:build",
    ],
    entry_point = ":server.js",
)

nodejs_image(
    name = "docker",
    binary = ":server",
)
