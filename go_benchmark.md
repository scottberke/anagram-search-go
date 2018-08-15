## GO Benchmark Test
{"anagrams":["ared","daer","dare","dear"]}
LOG: Response code = 200
LOG: header received:
HTTP/1.0 200 OK
Content-Type: application/json
Date: Wed, 15 Aug 2018 20:55:12 GMT
Content-Length: 42
Connection: keep-alive

{"anagrams":["ared","daer","dare","dear"]}
LOG: Response code = 200
Completed 10000 requests
Finished 10000 requests


Server Software:
Server Hostname:        localhost
Server Port:            3000

Document Path:          /anagrams/read.json
Document Length:        42 bytes

Concurrency Level:      10
Time taken for tests:   2.770 seconds
Complete requests:      10000
Failed requests:        0
Keep-Alive requests:    10000
Total transferred:      1740000 bytes
HTML transferred:       420000 bytes
Requests per second:    3609.81 [#/sec] (mean)
Time per request:       2.770 [ms] (mean)
Time per request:       0.277 [ms] (mean, across all concurrent requests)
Transfer rate:          613.39 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.0      0       1
Processing:     0    3  16.0      1     496
Waiting:        0    3  15.2      1     496
Total:          0    3  16.0      1     496

Percentage of the requests served within a certain time (ms)
  50%      1
  66%      2
  75%      2
  80%      3
  90%      4
  95%      7
  98%     10
  99%     16
 100%    496 (longest request)
