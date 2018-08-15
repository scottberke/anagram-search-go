## Ruby With Sinatra Benchmark Test

{"anagrams":["ared","daer","dare","dear"]}
LOG: Response code = 200
LOG: header received:
HTTP/1.1 200 OK
Content-Type: application/json
Content-Length: 42
X-Content-Type-Options: nosniff
Connection: keep-alive
Server: thin

{"anagrams":["ared","daer","dare","dear"]}
LOG: Response code = 200
Completed 10000 requests
Finished 10000 requests


Server Software:        thin
Server Hostname:        localhost
Server Port:            3000

Document Path:          /anagrams/read.json
Document Length:        42 bytes

Concurrency Level:      10
Time taken for tests:   6.834 seconds
Complete requests:      10000
Failed requests:        0
Keep-Alive requests:    10000
Total transferred:      1840000 bytes
HTML transferred:       420000 bytes
Requests per second:    1463.37 [#/sec] (mean)
Time per request:       6.834 [ms] (mean)
Time per request:       0.683 [ms] (mean, across all concurrent requests)
Transfer rate:          262.95 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.0      0       1
Processing:     2    7   3.7      6      82
Waiting:        1    7   3.7      6      81
Total:          2    7   3.7      6      82

Percentage of the requests served within a certain time (ms)
  50%      6
  66%      7
  75%      8
  80%      8
  90%     10
  95%     12
  98%     15
  99%     17
 100%     82 (longest request)
