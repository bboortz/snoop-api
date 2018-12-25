#!/bin/bash

set -e
set -u


if [ -n "${DEV:-}" ]; then
    go build -o snoop cmd/main.go
else
    docker build -t snoop -f build/Dockerfile .
fi
