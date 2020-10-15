## Ubuntu20.04安装docker-ce

```bash
$ curl -fsSL https://mirrors.ustc.edu.cn/docker-ce/linux/ubuntu/gpg | sudo apt-key add -
$ sudo add-apt-repository "deb [arch=amd64] https://mirrors.ustc.edu.cn/docker-ce/linux/ubuntu \
> $(lsb_release -cs) stable"

# sudo add-apt-repository \
#   "deb [arch=amd64] https://mirrors.tuna.tsinghua.edu.cn/docker-ce/linux/ubuntu \
#   $(lsb_release -cs) \
#   stable"

$ sudo apt update
$ sudo apt install docker-ce
$ docker info
Client: Docker Engine - Community
 Version:           19.03.12
 API version:       1.40
 Go version:        go1.13.10
 Git commit:        48a66213fe
 Built:             Mon Jun 22 15:45:44 2020
 OS/Arch:           linux/amd64
 Experimental:      false
# ...
```

## Docker配置国内源

```bash
$ sudo vim /etc/docker/daemon.json
```
填入如下内容：
```json
{
  "registry-mirrors": ["https://docker.mirrors.ustc.edu.cn"]
}
```
然后重启docker:
```bash
$ sudo service docker restart
```

## Ubuntu20.04安装docker-compose

```bash
$ sudo apt install docker-compose
```

## Docker方式部署etcd (可选)

- [etcd] docker 安装 https://www.jianshu.com/p/140a16408e98

```bash
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

```bash
root@xiaomo:/home/xiaomo# curl -L http://192.168.2.244:2379/v2/members
{"members":[{"id":"8d82831b0940236c","name":"etcd0","peerURLs":["http://192.168.2.244:2380"],"clientURLs":["http://192.168.2.244:2379","http://192.168.2.244:4001"]}]}
```

- docker 安装 etcdkeeper (webUI)

```bash
docker run -it -d --name etcdkeeper -p 8899:8080 evildecay/etcdkeeper:v0.7.6
```

## Docker-compose部署etcd (推荐)

- 批量启动容器
```bash
# sudo apt install docker-compose
# https://www.cnblogs.com/luliAngel/p/etcd.html
# https://www.jianshu.com/p/44022c67f117

# 详细参考docker-compose.yml文件
docker-compose up -d
```

- 打开etcdkeeper
```
# 打开浏览器，输入: http://localhost:8899/etcdkeeper/
# ifconfig /ipconfig 获得本机ip后
# 然后输入 172.30.0.10:2379  (或172.30.0.10:12379)， 正常情况下可以看到Nodes根目录
```
