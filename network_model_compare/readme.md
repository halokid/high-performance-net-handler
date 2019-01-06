几种golang后端网络模型的性能对比， 针对HTTP
============================

### 硬件配置
CPU i7 2.6G/ 16G内存 / MACOS

### model1 测试数据

```
ab -c 2000 -n 10000 http://127.0.0.1:8089/model1/

-----

This is ApacheBench, Version 2.3 <$Revision: 1843412 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 192.168.1.101 (be patient)


Server Software:
Server Hostname:        192.168.1.101
Server Port:            8089

Document Path:          /model1/
Document Length:        22 bytes

Concurrency Level:      2000
Time taken for tests:   10.155 seconds
Complete requests:      10000
Failed requests:        0
Total transferred:      1390000 bytes
HTML transferred:       220000 bytes
Requests per second:    984.77 [#/sec] (mean)
Time per request:       2030.921 [ms] (mean)
Time per request:       1.015 [ms] (mean, across all concurrent requests)
Transfer rate:          133.68 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    1  12.5      0    1249
Processing:   459 1855 536.4   1711    2934
Waiting:        1 1093 620.5   1106    2474
Total:        459 1856 536.4   1712    2934

Percentage of the requests served within a certain time (ms)
  50%   1712
  66%   2001
  75%   2057
  80%   2081
  90%   2710
  95%   2820
  98%   2893
  99%   2916
 100%   2934 (longest request)


```