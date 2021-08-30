load("@rules_cc//cc:defs.bzl", "cc_binary")

cc_binary(
    name = "hello-world",
    srcs = ["hello-world.cc"],
    defines = select({
	":fault_injector": ["FAULT_INJECTOR"],
        "//conditions:default": [],
    }),
)

config_setting(
    name = "fault_injector",
    values = {"define": "FAULT_INJECTOR=1"},
)
