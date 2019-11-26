#!/bin/bash

RUN_NAME="go_monitor"

mkdir output output/log
export GO111MODULE=on
go build -a -o output/${RUN_NAME}