load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "fds_other.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "fds_linux.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "github.com/coreos/etcd/pkg/runtime",
    visibility = ["//visibility:public"],
)
