load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

proto_library(
    name = "moredefaults_proto",
    srcs = ["md.proto"],
    visibility = ["//visibility:public"],
    deps = [
        "//github.com/gogo/protobuf/gogoproto:gogoproto_proto",
        "//github.com/gogo/protobuf/test/example:example_proto",
    ],
)

go_proto_library(
    name = "moredefaults_go_proto",
    importpath = "github.com/gogo/protobuf/test/moredefaults",
    proto = ":moredefaults_proto",
    visibility = ["//visibility:public"],
    deps = [
        "//github.com/gogo/protobuf/gogoproto:go_default_library",
        "//github.com/gogo/protobuf/test/example:go_default_library",
    ],
)

go_library(
    name = "go_default_library",
    srcs = ["md.pb.go"],
    importpath = "github.com/gogo/protobuf/test/moredefaults",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/gogo/protobuf/gogoproto:go_default_library",
        "//vendor/github.com/gogo/protobuf/proto:go_default_library",
        "//vendor/github.com/gogo/protobuf/test/example:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "md_test.go",
        "mdpb_test.go",
    ],
    importpath = "github.com/gogo/protobuf/test/moredefaults",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/gogo/protobuf/gogoproto:go_default_library",
        "//vendor/github.com/gogo/protobuf/jsonpb:go_default_library",
        "//vendor/github.com/gogo/protobuf/proto:go_default_library",
        "//vendor/github.com/gogo/protobuf/test/example:go_default_library",
    ],
)

filegroup(
    name = "go_default_library_protos",
    srcs = ["md.proto"],
    visibility = ["//visibility:public"],
)
