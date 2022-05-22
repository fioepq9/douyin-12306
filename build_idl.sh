#!/bin/env/bash

# user
kitex -type protobuf -module douyin-12306 ./idl/user.proto

# video
kitex -type protobuf -module douyin-12306 ./idl/video.proto
