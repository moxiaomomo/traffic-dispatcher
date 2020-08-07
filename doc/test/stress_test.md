### http 接口压力测试示例

`go-wrk`是一个用 Go 语言实现的轻量级的 http 基准测试工具，类似于`wrk`，本文将简单介绍一下如何使用`go-wrk`实现接口的性能(压力)测试。

- 安装 go-wrk

```
https://github.com/wg/wrk
# 本次采用go-wrk测试
https://github.com/adjust/go-wrk
```

```shell
git clone git://github.com/adeven/go-wrk.git
cd go-wrk
# 开启go modules的情况， 需要初始化配置(go1.11或以上)
go mod init go-wrk
go build
```

- go-wrk 选项说明

```shell
# #go-wrk [flags] url
#./go-wrk --help
Usage of ./go-wrk:
  -CA string
    	A PEM eoncoded CA's certificate file. (default "someCertCAFile")
  -H string
    	the http headers sent separated by '\n' (default "User-Agent: go-wrk 0.1 benchmark\nContent-Type: text/html;")
  -b string
    	the http request body
  -c int
    	the max numbers of connections used (default 100)
  -cert string
    	A PEM eoncoded certificate file. (default "someCertFile")
  -d string
    	dist mode
  -f string
    	json config file
  -i	TLS checks are disabled
  -k	if keep-alives are disabled (default true)
  -key string
    	A PEM encoded private key file. (default "someKeyFile")
  -m string
    	the http request method (default "GET")
  -n int
    	the total number of calls processed (default 1000)
  -p string
    	the http request body data file
  -r	in the case of having stream or file in the response,
    	 it reads all response body to calculate the response size
  -s string
    	if specified, it counts how often the searched string s is contained in the responses
  -t int
    	the numbers of threads used (default 1)
```

- 测试`/test/qurey`接口 (假设提前开启了一个 web 服务，监听了 8082 端口)

```shell
# 8个线程，400个连接， 模拟10w次请求
./go-wrk -c=400 -t=8 -n=100000 "http://localhost:8082/test/query?lat=39.915&lng=116.404"
```

- 测试结果说明

```
==========================BENCHMARK==========================
URL:				http://localhost:8082/test/query?lat=39.915&lng=116.404

Used Connections:		400
Used Threads:			8
Total number of calls:		100000

===========================TIMINGS===========================
Total time passed:		99.46s
Avg time per request:		395.89ms
Requests per second:		1005.39
Median time per request:	368.22ms
99th percentile time:		775.90ms
Slowest time for request:	1479.00ms

=============================DATA=============================
Total response body sizes:		501200000
Avg response body per request:		5012.00 Byte
Transfer rate per second:		5038989.99 Byte/s (5.04 MByte/s)
==========================RESPONSES==========================
20X Responses:		100000	(100.00%)
30X Responses:		0	(0.00%)
40X Responses:		0	(0.00%)
50X Responses:		0	(0.00%)
Errors:			0	(0.00%)
```

可以看到，关于接口`/test/query`，(1)每秒可以处理`1005次`请求(即 QPS)；(2)每秒传输`5.04MB`数据(吞吐量)；(3)响应码为`20x`开头的请求为 100%, 即没有发生业务之外的错误(比如 502);(4)%99 的请求的平均处理时间为`775.90ms`。
