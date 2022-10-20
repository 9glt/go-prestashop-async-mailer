#!/bin/bash

docker run --rm -it -w /app -v $(pwd):/app  -e CGO_ENABLED=0 golang:1.17-alpine go mod download && g
o build
