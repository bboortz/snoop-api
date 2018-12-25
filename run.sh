#!/bin/bash

set -e
set -u


if [ -n "${DEV:-}" ]; then
    go run cmd/main.go
else
    docker run  -i -t --rm --name snoop  -p 8443:8443 snoop
fi