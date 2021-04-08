package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_jsonnet//jsonnet:jsonnet.bzl",
    "jsonnet_library",
    "jsonnet_to_json",
    "jsonnet_to_json_test",
)

jsonnet_library(
    name = "workflow",
    srcs = ["workflow.libsonnet"],
)

jsonnet_to_json(
    name = "wordcount",
    src = "wordcount.jsonnet",
    outs = ["wordcount.json"],
    deps = [":workflow"],
)

jsonnet_to_json_test(
    name = "wordcount_test",
    src = "wordcount.jsonnet",
    golden = "wordcount.json",
    deps = [":workflow"],
)
