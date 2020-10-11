#!/bin/bash

ps aux | grep "\-\-registry\=etcd" | awk '{print $2}' | xargs kill -9
