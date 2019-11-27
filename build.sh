#!/bin/bash

RUN_NAME="go_monitor"

mkdir output output/monitor_log
export GO111MODULE=on
go build -a -o output/${RUN_NAME}