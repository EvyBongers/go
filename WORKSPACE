load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "7f1aa43d986df189f7cf30e81dd2dc9d8ed7c74e356341a17267f6d7e5748382",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.24.1/rules_go-v0.24.1.tar.gz",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.24.1/rules_go-v0.24.1.tar.gz",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "d4113967ab451dd4d2d767c3ca5f927fec4b30f3b2c6f8135a2033b9c05a5687",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.22.0/bazel-gazelle-v0.22.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.22.0/bazel-gazelle-v0.22.0.tar.gz",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

go_rules_dependencies()

go_register_toolchains()

gazelle_dependencies()

http_archive(
    name = "io_bazel_rules_docker",
    sha256 = "4521794f0fba2e20f3bf15846ab5e01d5332e587e9ce81629c7f96c793bb7036",
    strip_prefix = "rules_docker-0.14.4",
    urls = ["https://github.com/bazelbuild/rules_docker/releases/download/v0.14.4/rules_docker-v0.14.4.tar.gz"],
)

load(
    "@io_bazel_rules_docker//repositories:repositories.bzl",
    container_repositories = "repositories",
)

container_repositories()

load("@io_bazel_rules_docker//repositories:deps.bzl", container_deps = "deps")

container_deps()

load("@io_bazel_rules_docker//repositories:pip_repositories.bzl", "pip_deps")

pip_deps()

go_repository(
    name = "com_github_gorilla_mux",
    importpath = "github.com/gorilla/mux",
    sha256 = "56ea3c39b8c19360357970a4dcd717485f36ac7da9c6b5a107bd4b80848c4aee",
    strip_prefix = "mux-1.7.4",
    urls = ["https://github.com/gorilla/mux/archive/v1.7.4.tar.gz"],
)
