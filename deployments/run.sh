#!/bin/bash
set -eux

cd "$(dirname "$(realpath "${0}")")"

bazel run //cmd/go:go_container

docker-compose up
