## 注册 redis.conf

具体参考本目录的`redis.conf`

## 配置 redis 及映射目录

```bash
$ sudo apt install redis-tools
$ sudo mkdir -p /usr/local/redis/conf /usr/local/redis/data /usr/local/redis/logs
$ sudo cp redis.conf /usr/local/redis/conf/
$ sudo docker pull redis:6.0
$ sudo docker run --privileged=true -itd \
-v /usr/local/redis/conf/redis.conf:/usr/local/etc/redis/redis.conf \
-v /usr/local/redis/data:/data -v /usr/local/redis/logs:/logs \
-p 16379:6379 \
--name session-redis \
redis:6.0 \
redis-server /usr/local/etc/redis/redis.conf
```

## 检查 redis 容器状态

```bash
$ sudo docker ps | grep session-redis
6afa6a06e5b4        redis:6.0                     "docker-entrypoint.s…"   4 minutes ago       Up 4 minutes        0.0.0.0:16379->6379/tcp                                                     session-redis
```

## 检查 redis 连接 (宿主机上访问 redis)

```shell
$ redis-cli -h 127.0.0.1 -p 16379
127.0.0.1:16379> auth test123456
OK
127.0.0.1:16379> keys *
(empty list or set)
127.0.0.1:16379>
```
