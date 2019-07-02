#!/bin/bash
set -eux

cd "$(dirname "$(realpath "${0}")")"

bazel run --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 //cmd/go:go_container

docker-compose up
