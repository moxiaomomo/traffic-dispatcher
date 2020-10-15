## go (1.14)

官网下载包即可安装。

## docker, docker-compose, etcd

详细参考当前目录下的`docker-etcd`文档。

## protobuf

详细参考当前目录的`proto_cfg.md`文档。

## mongodb

- 安装 mongodb

```bash
sudo apt install mongodb
# MongoDB server version: 3.6.8
# sudo yum install mongodb
```

- 创建 index

```
> db.geoInfo.createIndex({geoinfo: "2dsphere"})
> db.geoInfo.createIndex({h3index:1})
> db.geoInfo.createIndex({uid:1}, {unique:1})
```

- 示例

```
> db.geoInfo.find()
{ "_id" : ObjectId("5f38d6193f24b503fa68a35e"), "uid" : "testuid_0", "geoinfo" : { "type" : "Point", "coordinates" : [ 116.404, 39.915 ] }, "h3index" : "0x8731aa428ffffff", "name" : "testuser_0" }
{ "_id" : ObjectId("5f38d6193f24b503fa68a35f"), "uid" : "testuid_2", "geoinfo" : { "type" : "Point", "coordinates" : [ 118.404, 41.915 ] }, "h3index" : "0x8731adaa4ffffff", "name" : "testuser_2" }
{ "_id" : ObjectId("5f38d6193f24b503fa68a360"), "uid" : "testuid_1", "geoinfo" : { "coordinates" : [ 117.404, 40.915 ], "type" : "Point" }, "h3index" : "0x8731a8002ffffff", "name" : "testuser_1" }
```

## redis

```bash
sudo apt install redis-server
# 5.0.7
# yum install redis
```
