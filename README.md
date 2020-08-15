# traffic-dispatcher _`(In development)`_

共享出行调度服务

### 架构设计(V0.2)

![archi_0.2.png](https://raw.githubusercontent.com/moxiaomomo/traffic-dispatcher/master/doc/archi/archi_0.2.png)

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

(前端测试) [web_admin](https://github.com/moxiaomomo/traffic-dispatcher-admin)

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

### 测试

```shell
# in development
# --registry_address 按实际情况修改
# 启动 order backend service
go run service/order/main.go --registry=etcd --registry_address=192.168.2.244:2379
# 启动 driver api service
go run api/driver/main.go --registry=etcd --registry_address=192.168.2.244:2379
# 启动micro api gateway
micro --registry=etcd --registry_address=192.168.2.244:2379 api
```
