#!/usr/bin/env sh

podman run --rm -it -v "$(pwd)":/project -v go:/root/go/pkg/mod -w=/project --network=host golang:alpine
