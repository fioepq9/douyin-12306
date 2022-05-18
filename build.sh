#!/bin/env/bash
kitex -type protobuf -module douyin-12306 ./idl/user.proto
docker image rm douyin-12306-golang
docker build -f golang.Dockerfile -t douyin-12306-golang .