package(default_visibility = ["//kythe:default_visibility"])

# TODO(fromberger): The baseline proto_library rule does not have a hook to
# request metadata. That probably means these examples do not work at present,
# though there are no tests to verify that.  Once we have unblocked Bazel
# upgrades, revisit how this should work.

proto_library(
    name = "example_proto",
    srcs = ["example.proto"],
)

cc_proto_library(
    name = "example_cc_proto",
    deps = [":example_proto"],
)

cc_library(
    name = "proto_user",
    srcs = ["proto_user.cc"],
    copts = [
        "-Wno-non-virtual-dtor",
        "-Wno-unused-variable",
        "-Wno-implicit-fallthrough",
    ],
    deps = [":example_cc_proto"],
)

cc_binary(
    name = "proto_indexer",
    srcs = ["proto_indexer.cc"],
    copts = [
        "-Wno-non-virtual-dtor",
        "-Wno-unused-variable",
        "-Wno-implicit-fallthrough",
    ],
    deps = [
        "//kythe/cxx/common:lib",
        "//kythe/cxx/common/indexing:lib",
        "@com_github_google_glog//:glog",
        "@com_google_protobuf//:cc_wkt_protos",
    ],
)
