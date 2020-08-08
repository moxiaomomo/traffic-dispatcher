#!/bin/bash
# v0.1 in development

rootDir=`pwd`
mkdir -p ${rootDir}/build
rm -rf ${rootDir}/build/*

# service/driver
cd ${rootDir}/service/driver/ && go build -o ${rootDir}/build/driver-api *.go 
# service/gateway
cd ${rootDir}/service/gateway/ && go build -o ${rootDir}/build/apigw *.go 