#!/bin/bash

RUN_NAME="go_monitor"

mkdir output output/${RUN_NAME}_log
cp -r config output/
cp -r data output/
export GO111MODULE=on

# dev
# go build -a -o output/${RUN_NAME} -ldflags "-w"

# prod
go build -a -o output/${RUN_NAME}

# 取消交叉编译
# CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o output/${RUN_NAME}