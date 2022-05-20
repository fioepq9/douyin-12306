#!/bin/env/bash

# 启动 etcd 和 jaeger
docker compose up &

# 启动 api 服务
cd cmd/api && ./run.sh &

# 启动 user 服务
cd cmd/user && ./build.sh && output/bootstrap.sh &
