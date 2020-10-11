#!/bin/bash

ps aux | grep -E "go run|micro" | grep "\-\-registry\=etcd" | awk '{print $2}' | xargs kill -9
