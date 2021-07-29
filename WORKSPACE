load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

# Go environment setup
http_archive(
    name = "io_bazel_rules_go",
    sha256 = "7c10271940c6bce577d51a075ae77728964db285dac0a46614a7934dc34303e6",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.26.0/rules_go-v0.26.0.tar.gz",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.26.0/rules_go-v0.26.0.tar.gz",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "62ca106be173579c0a167deb23358fdfe71ffa1e4cfdddf5582af26520f1c66f",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.23.0/bazel-gazelle-v0.23.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.23.0/bazel-gazelle-v0.23.0.tar.gz",
    ],
)

http_archive(
    name = "build_stack_rules_proto",
    urls = ["https://github.com/stackb/rules_proto/archive/b2913e6340bcbffb46793045ecac928dcf1b34a5.tar.gz"],
    strip_prefix = "rules_proto-b2913e6340bcbffb46793045ecac928dcf1b34a5",
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

go_rules_dependencies()

go_register_toolchains(version = "1.16")

gazelle_dependencies()

load("@build_stack_rules_proto//github.com/grpc-ecosystem/grpc-gateway:deps.bzl", "gateway_grpc_compile")

gateway_grpc_compile()

# Nodejs related
http_archive(
    name = "build_bazel_rules_nodejs",
    sha256 = "4952ef879704ab4ad6729a29007e7094aef213ea79e9f2e94cbe1c9a753e63ef",
    urls = ["https://github.com/bazelbuild/rules_nodejs/releases/download/2.2.0/rules_nodejs-2.2.0.tar.gz"],
)

load("@build_bazel_rules_nodejs//:index.bzl", "yarn_install")

yarn_install(
    name = "npm",
    package_json = "//javascript:package.json",
    yarn_lock = "//javascript:yarn.lock",
)

# Docker related
http_archive(
    name = "io_bazel_rules_docker",
    sha256 = "1698624e878b0607052ae6131aa216d45ebb63871ec497f26c67455b34119c80",
    strip_prefix = "rules_docker-0.15.0",
    urls = ["https://github.com/bazelbuild/rules_docker/releases/download/v0.15.0/rules_docker-v0.15.0.tar.gz"],
)

load(
    "@io_bazel_rules_docker//repositories:repositories.bzl",
    container_repositories = "repositories",
)

container_repositories()

# googleapis
git_repository(
    name = "com_google_googleapis",
    remote = "https://github.com/googleapis/googleapis.git",
    commit = "8de886fe5447d097e7f51c20aed40738f7eb68e2",
)

load("@com_google_googleapis//:repository_rules.bzl", "switched_rules_by_language")

switched_rules_by_language(
    name = "com_google_googleapis_imports",
    grpc = True,
    go = True,
)


load("@io_bazel_rules_docker//repositories:deps.bzl", container_deps = "deps")

container_deps()

load(
    "@io_bazel_rules_docker//nodejs:image.bzl",
    nodejs_image_repositories = "repositories",
)

nodejs_image_repositories()

# Foreign cc
# Rule repository, note that it's recommended to use a pinned commit to a released version of the rules
# version 0.2.0
git_repository(
    name = "rules_foreign_cc",
    commit = "d02390f1363cdd2ba5a7f7907a481503d483d569",
    remote = "https://github.com/bazelbuild/rules_foreign_cc",
)

load("@rules_foreign_cc//foreign_cc:repositories.bzl", "rules_foreign_cc_dependencies")

# glog
go_repository(
    name = "com_github_golang_glog",
    importpath = "github.com/golang/glog",
    commit = "23def4e6c14b4da8ac2ed8007337bc5eb5007998",
)

# This sets up some common toolchains for building targets. For more details, please see
# https://github.com/bazelbuild/rules_foreign_cc/tree/main/docs#rules_foreign_cc_dependencies
rules_foreign_cc_dependencies()

_ALL_CONTENT = """\
filegroup(
    name = "all_srcs",
    srcs = glob(["**"]),
    visibility = ["//visibility:public"],
)
"""

# jsconnet
http_archive(
    name = "io_bazel_rules_jsonnet",
    sha256 = "7f51f859035cd98bcf4f70dedaeaca47fe9fbae6b199882c516d67df416505da",
    strip_prefix = "rules_jsonnet-0.3.0",
    urls = ["https://github.com/bazelbuild/rules_jsonnet/archive/0.3.0.tar.gz"],
)

load("@io_bazel_rules_jsonnet//jsonnet:jsonnet.bzl", "jsonnet_repositories")

jsonnet_repositories()

load("@jsonnet_go//bazel:repositories.bzl", "jsonnet_go_repositories")

jsonnet_go_repositories()

load("@jsonnet_go//bazel:deps.bzl", "jsonnet_go_dependencies")

jsonnet_go_dependencies()

# OpenBLAS
# OpenBLAS source code repository
http_archive(
    name = "openblas",
    build_file_content = _ALL_CONTENT,
    strip_prefix = "OpenBLAS-0.3.2",
    urls = ["https://github.com/xianyi/OpenBLAS/archive/v0.3.2.tar.gz"],
)

# Eigen source code repository
http_archive(
    name = "eigen",
    build_file_content = _ALL_CONTENT,
    strip_prefix = "eigen-git-mirror-3.3.5",
    urls = ["https://github.com/eigenteam/eigen-git-mirror/archive/3.3.5.tar.gz"],
)

# buildifier
http_archive(
    name = "com_google_protobuf",
    strip_prefix = "protobuf-master",
    urls = ["https://github.com/protocolbuffers/protobuf/archive/master.zip"],
)

load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

protobuf_deps()

http_archive(
    name = "com_github_bazelbuild_buildtools",
    strip_prefix = "buildtools-master",
    url = "https://github.com/bazelbuild/buildtools/archive/master.zip",
)
