load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

proto_library(
    name = "authpb_proto",
    srcs = ["auth.proto"],
    visibility = ["//visibility:public"],
    deps = ["//gogoproto:gogoproto_proto"],
)

go_proto_library(
    name = "authpb_go_proto",
    importpath = "github.com/coreos/etcd/auth/authpb",
    proto = ":authpb_proto",
    visibility = ["//visibility:public"],
    deps = ["//gogoproto:go_default_library"],
)

go_library(
    name = "go_default_library",
    srcs = ["auth.pb.go"],
    importpath = "github.com/coreos/etcd/auth/authpb",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/golang/protobuf/proto:go_default_library"],
)

filegroup(
    name = "go_default_library_protos",
    srcs = ["auth.proto"],
    visibility = ["//visibility:public"],
)
