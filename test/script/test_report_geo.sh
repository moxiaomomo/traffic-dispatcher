#!/bin/bash
# in development

registry="--registry=etcd --registry_address=172.30.0.10:2379"
rootDir=`pwd`
mkdir -p ${rootDir}/build
mkdir -p ${rootDir}/logs

# service/lbs
echo "build and run service/lbs..."
rm -f ${rootDir}/build/lbs-srv
go build -o ${rootDir}/build/lbs-srv ${rootDir}/service/lbs/*.go
nohup ${rootDir}/build/lbs-srv $registry >> ${rootDir}/logs/lbs-srv.log 2>&1
# web/geo
echo "build and run web/geo..."
rm -f ${rootDir}/build/geo-web
go build -o ${rootDir}/build/geo-web ${rootDir}/web/geo/*.go
nohup ${rootDir}/build/geo-web $registry >> ${rootDir}/logs/geo-web.log 2>&1
# micro api
echo "run api web"
nohup micro $registry web >> ${rootDir}/logs/micro-api-web.log 2>&1