#!/usr/bin/env bash

GOOS=linux GOARCH=amd64 go build -o ./bin/linux64/binlog-parser
GOOS=darwin GOARCH=amd64 go build -o ./bin/darwin64/binglog-parser