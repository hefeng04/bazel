load("@bazel_skylib//rules:write_file.bzl", "write_file")
load("@build_bazel_rules_nodejs//:index.bzl", "copy_to_bin", "nodejs_test")
load("@npm//react-scripts:index.bzl", "react_scripts", "react_scripts_test")

package(default_visibility = ["//go:go_visibility"])

# Filename conventions described at
# https://create-react-app.dev/docs/running-tests#filename-conventions
_TESTS = [
    "src/**/*.test.js*",
    "src/**/*.test.ts*",
    "src/**/*.spec.js*",
    "src/**/*.spec.ts*",
    "src/**/__tests__/**/*.js*",
    "src/**/__tests__/**/*.ts*",
]

# We don't want to teach react-scripts to include from multiple directories
# So we copy everything it wants to read to the output "bin" directory
copy_to_bin(
    name = "copy_static_files",
    srcs = glob(
        [
            "public/*",
            "src/**/*",
        ],
        exclude = _TESTS,
    ) + [
        "package.json",
    ],
)

# react-scripts can only work if the working directory is the root of the application.
# So we'll need to chdir before it runs.
write_file(
    name = "write_chdir_script",
    out = "chdir.js",
    content = ["process.chdir(__dirname)"],
)

_RUNTIME_DEPS = [
    "chdir.js",
    "copy_static_files",
    "@npm//axios",
    "@npm//antd",
    "@npm//react",
    "@npm//react-dom",
    "@npm//react-router-dom",
    "@npm//react-scripts",
    "@npm//web-vitals",
]

react_scripts(
    # Note: this must be named "build" since react-scripts hard-codes that as the output dir
    name = "build",
    args = [
        "--node_options=--require=./$(execpath chdir.js)",
        "build",
    ],
    data = _RUNTIME_DEPS + [
        "@npm//@types",
    ],
    output_dir = True,
)
