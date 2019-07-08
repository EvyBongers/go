#!/bin/bash
set -eux

SCRIPT="$(realpath "${0}")"
SCRIPT_DIR="$(dirname "${SCRIPT}")"
cd "${SCRIPT_DIR}"

bazel run --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 //cmd/go:go_container

docker-compose up
