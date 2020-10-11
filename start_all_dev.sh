#!/bin/bash

regHost=$1
if [[ $regHost == "" ]];then
  echo "需要指定registry_address"
  exit 1
fi
# in development
# --registry_address 按实际情况修改
# 启动 user backend service
nohup go run service/user/main.go --registry=etcd --registry_address=${regHost} >> /tmp/srv_user.log 2>&1 &

# 启动 order backend service
nohup go run service/order/main.go --registry=etcd --registry_address=${regHost} >> /tmp/srv_order.log 2>&1 &

# 启动 lbs backend service
nohup go run service/lbs/main.go --registry=etcd --registry_address=${regHost} >> /tmp/srv_lbs.log 2>&1 &

# 启动 driver api service
nohup go run api/driver/main.go --registry=etcd --registry_address=${regHost} >> /tmp/api_driver.log 2>&1 &

# 启动 passenger api service
nohup go run api/passenger/main.go --registry=etcd --registry_address=${regHost} >> /tmp/api_passenger.log 2>&1 &

# 启动 geo web
nohup go run web/geo/main.go --registry=etcd --registry_address=${regHost} >> /tmp/web_geo.log 2>&1 &

# 启动micro api gateway
nohup micro --registry=etcd --registry_address=${regHost} api --handler=api >> /tmp/micro_api.log 2>&1 &

# 启动micro web gateway
nohup micro --registry=etcd --registry_address=${regHost} web >> /tmp/micro_web.log 2>&1 &
