#!/bin/bash

set -e
set -u

go test -v ./...
echo "EXIT CODE: $?"
