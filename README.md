# traffic-dispatcher

共享出行调度服务

### 架构设计(V0.2)

![archi_0.2.png](http://47.107.169.20/archi_0.2.png)

### 微服务划分

- admin 后台管理调度 (默认端口: 18080)
- driver 司机 api 服务 (默认端口: 18000)
- passanger 乘客 api 服务 (默认端口: 18001)
- order 订单管理服务 (默认端口: 18002)
- lbs 地理位置服务 (默认端口： 18003)
- dispatcher 派遣调度服务 (默认端口：18004)
- notification 全局消息服务 (默认端口：18005)
