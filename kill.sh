#!/bin/env/bash
lsof_9090=$(lsof -i | grep 9090)
pid_of_9090=$(echo $lsof_9090 | cut -d ' ' -f 2) 
kill -9 $pid_of_9090


lsof_9091=$(lsof -i | grep 9091)
pid_of_9091=$(echo $lsof_9091 | cut -d ' ' -f 2) 
kill -9 $pid_of_9091

docker-compose down
