#!/bin/bash

RUN_NAME="go_monitor"

mkdir output output/${RUN_NAME}_log
cp -r config output/
export GO111MODULE=on
go build -a -o output/${RUN_NAME}_dev
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o output/${RUN_NAME}