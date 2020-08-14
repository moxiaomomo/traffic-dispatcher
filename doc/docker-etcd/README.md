- [etcd] docker 安装 https://www.jianshu.com/p/140a16408e98

```
docker run -d -v /usr/share/ca-certificates/:/etc/ssl/certs -p 4001:4001 -p 2380:2380 -p 2379:2379 \
 --name etcd etcd /usr/local/bin/etcd \
 -name etcd0 \
 -advertise-client-urls http://0.0.0.0:2379,http://0.0.0.0:4001 \
 -listen-client-urls http://0.0.0.0:2379,http://0.0.0.0:4001 \
 -initial-advertise-peer-urls http://0.0.0.0:2380 \
 -listen-peer-urls http://0.0.0.0:2380 \
 -initial-cluster-token etcd-cluster-1 \
 -initial-cluster etcd0=http://0.0.0.0:2380 \
 -initial-cluster-state new
```

```
root@xiaomo:/home/xiaomo# curl -L http://192.168.2.244:2379/v2/members
{"members":[{"id":"8d82831b0940236c","name":"etcd0","peerURLs":["http://192.168.2.244:2380"],"clientURLs":["http://192.168.2.244:2379","http://192.168.2.244:4001"]}]}
```

- docker 安装 etcdkeeper (webUI)

```
docker run -it -d --name etcdkeeper -p 8899:8080 evildecay/etcdkeeper
```


- docker-compose启动(etcd及etcdkeeper)

```
# sudo apt install docker-compose
docker-compose up -d
```
