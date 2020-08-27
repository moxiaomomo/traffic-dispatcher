# traffic-dispatcher _`(In development)`_

共享出行调度服务

### 架构设计(V0.3)

![archi_0.3.png](https://raw.githubusercontent.com/moxiaomomo/traffic-dispatcher/master/doc/archi/archi_0.3.png)

### API 列表简介(V0.2)

![api-design_0.2.png](https://raw.githubusercontent.com/moxiaomomo/traffic-dispatcher/master/doc/archi/api-design_0.2.png)

### 环境要求及相关配置

(后端服务)

- Go: 支持 1.14 或以上
- Go 包管理: Go Modules
- Redis
- MongoDB
- MySQL
- Etcd
- RabbitMQ (Optional)
- Protobuf (V3)
- Docker (部署微服务等)

(前端测试)

[web_admin](https://github.com/moxiaomomo/traffic-dispatcher-admin)

- Vue 2.x
- Baidu map API
- Typescript 3.x

[app_test](https://github.com/moxiaomomo/traffic-dispatcher-cli)

- uni-app
- Vue 2.x
- Baidu map API
- Typescript 3.x

### 微服务划分

- admin 后台管理调度 (默认端口: 18080)
- driver 司机 api 服务 (默认端口: 18000)
- passenger 乘客 api 服务 (默认端口: 18001)
- order 订单管理服务 (默认端口: 18002)
- lbs 地理位置服务 (默认端口： 18003)
- dispatcher 派遣调度服务 (默认端口：18004)
- notification 全局消息服务 (默认端口：18005)

### 编译

- 编译 proto

```shell
# geo.proto
protoc --proto_path=. --micro_out=./proto/geo/ --go_out=./proto/geo/ proto/geo/geo.proto
```

### 测试

- 测试 web 接口

```bash
# in development
# --registry_address 按实际情况修改
# 启动 user backend service
go run service/user/main.go --registry=etcd --registry_address=172.30.0.10:2379
# 启动 order backend service
go run service/order/main.go --registry=etcd --registry_address=172.30.0.10:2379
# 启动 driver api service
go run api/driver/main.go --registry=etcd --registry_address=172.30.0.10:2379
# 启动 passenger api service
go run api/passenger/main.go --registry=etcd --registry_address=172.30.0.10:2379
# 启动micro api gateway
micro --registry=etcd --registry_address=172.30.0.10:2379 api --handler=api
# 测试

## signup
curl -X POST "http://localhost:8080/passenger/user/signup" -H "content-type:application/json" -d '{"role":0,"userName":"xiaomo","userPwd":"123456"}'
# {"code":1,"msg":"Signup succeeded."}
curl -X POST "http://localhost:8080/driver/user/signup" -H "content-type:application/json" -d '{"role":1,"userName":"xiaohua","userPwd":"123456"}'
# {"code":1,"msg":"Signup succeeded."}

## signin
curl -X POST "http://localhost:8080/passenger/user/signin" -H "content-type:application/json" -d '{"role":0,"userName":"xiaomo","userPwd":"123456"}'
# {"code":1,"msg":"Signin succeeded.","user":{"id":6,"userID":"97d09d9efec8df12cfd093a79599efff","userName":"xiaomo","userPwd":"123456","lastActive":18446744011573954816,"token":"a3170402fa50535a38fe1a63aff749335f47d8fa"}}
curl -X POST "http://localhost:8080/driver/user/signin" -H "content-type:application/json" -d '{"role":0,"userName":"xiaohua","userPwd":"123456"}'
# {"code":1,"msg":"Signin succeeded.","user":{"id":7,"role":1,"userID":"8c920cfdf46bdcc7744335e44684e594","userName":"xiaohua","userPwd":"123456","token":"c009573d73644d81e126668527ef05f65f47d8d8"}}

## create and accept order
curl -X POST "http://localhost:8080/passenger/order/createOrder" -H "content-type:application/json" -d '{"srcGeo":"[110,26]","destGeo":"[112,30]","passengerId":"97d09d9efec8df12cfd093a79599efff"}'
# {{"code":10000,"msg":"","order":{"id":2,"orderId":"53463c941b2b6e7b7b6ab28bd13b31ae","srcGeo":"[110,26]","destGeo":"[112,30]","createAt":1598544547,"passengerId":"97d09d9efec8df12cfd093a79599efff"}}
curl -X POST "http://localhost:8080/driver/order/acceptOrder" -H "content-type:application/json" -d '{"orderId":"53463c941b2b6e7b7b6ab28bd13b31ae","driverId":"8c920cfdf46bdcc7744335e44684e594"}'
# {"code":10000,"msg":"","order":{"orderId":"53463c941b2b6e7b7b6ab28bd13b31ae","acceptAt":1598544858,"driverId":"8c920cfdf46bdcc7744335e44684e594","status":1}}
```

- 测试 websocket 传输

```bash
# service/lbs
go run service/lbs/main.go --registry=etcd --registry_address=172.30.0.10:2379
# web/geo
go run web/geo/main.go --registry=etcd --registry_address=172.30.0.10:2379
# micro web
micro --registry=etcd --registry_address=172.30.0.10:2379 web
```
