## Ruby On Rails API Benchmark Test
{"anagrams":["ared","daer","dare","dear"]}
LOG: header received:
HTTP/1.0 200 OK
Content-Type: application/json; charset=utf-8
ETag: W/"1f56a3ded2faad42b7a62d3ef23a70f4"
Cache-Control: max-age=0, private, must-revalidate
X-Request-Id: 670ff1ec-5689-46bc-8faa-c8c18d7a7b4b
X-Runtime: 0.001709
Connection: Keep-Alive


LOG: Response code = 200
Completed 10000 requests
Finished 10000 requests


Server Software:
Server Hostname:        0.0.0.0
Server Port:            3000

Document Path:          /anagrams/read.json
Document Length:        0 bytes

Concurrency Level:      10
Time taken for tests:   19.992 seconds
Complete requests:      10000
Failed requests:        1453
   (Connect: 0, Receive: 0, Length: 1453, Exceptions: 0)
Keep-Alive requests:    10000
Total transferred:      3009958 bytes
HTML transferred:       61026 bytes
Requests per second:    500.21 [#/sec] (mean)
Time per request:       19.992 [ms] (mean)
Time per request:       1.999 [ms] (mean, across all concurrent requests)
Transfer rate:          147.03 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.0      0       1
Processing:     1   20 386.4      8   17034
Waiting:        0   14 386.4      2   17034
Total:          1   20 386.4      8   17034

Percentage of the requests served within a certain time (ms)
  50%      8
  66%     10
  75%     12
  80%     14
  90%     19
  95%     25
  98%     36
  99%     45
 100%  17034 (longest request)
