- 下载 protobuf 相关 go 库

```
go get -u github.com/golang/protobuf/proto
go get -u github.com/golang/protobuf/protoc-gen-go
go get github.com/micro/micro/v2/cmd/protoc-gen-micro/v2
```

- Linux(Ubuntu)安装 protoc

```shell
# use version 3.12
wget https://github.com/protocolbuffers/protobuf/releases/download/v3.12.3/protoc-3.12.3-linux-x86_64.zip
unzip protoc-3.12.3-linux-x86_64.zip -d protoc
sudo mv protoc /usr/local/
sudo ln -s /usr/local/protoc/bin/protoc /usr/sbin/
protoc --version
# libprotoc 3.12.3
```

- grpc 版本冲突问题

```
# github.com/coreos/etcd/clientv3/balancer/resolver/endpoint
/data/go/pkg/mod/github.com/coreos/etcd@v3.3.18+incompatible/clientv3/balancer/resolver/endpoint/endpoint.go:114:78: undefined: resolver.BuildOption
/data/go/pkg/mod/github.com/coreos/etcd@v3.3.18+incompatible/clientv3/balancer/resolver/endpoint/endpoint.go:182:31: undefined: resolver.ResolveNowOption
# github.com/coreos/etcd/clientv3/balancer/picker
/data/go/pkg/mod/github.com/coreos/etcd@v3.3.18+incompatible/clientv3/balancer/picker/err.go:37:44: undefined: balancer.PickOptions
/data/go/pkg/mod/github.com/coreos/etcd@v3.3.18+incompatible/clientv3/balancer/picker/roundrobin_balanced.go:55:54: undefined: balancer.PickOptions
```

解决方法 (指定以下库版本):

```
google.golang.org/grpc v1.26.0
google.golang.org/protobuf v1.23.0
```

- 编译 proto 示例

```shell
$ cd proto/driver
$ protoc --proto_path=. --proto_path=. --micro_out=. --go_out=. driver.proto
```
