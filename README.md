# Anagram Search

## Description
This application provides fast searches for anagrams. Anagrams are ingested when the application loads and provides the following endpoints.
Three endpoints currently exist:
 1. POST [`/words.json`](#words)
 2. GET [`/anagrams/:word.json`](#anagrams)
 3. DELETE [`/words/:word.json`](#words)
 4. DELETE [`/words.json`](#words)


These endpoints are documented below with example usage.



## To Run Locally
To install and build, execute:
```bash
  $ go get github.com/scottberke/anagram-search-go
  $ cd $GOPATH/src/github.com/scottberke/anagram-search-go
  $ go build
```

To see available flags:
```bash
$ ./anagram-search-go -h
Usage of ./anagram-search-go:
  -port int
    	a port to start the server on (default 8080)
```

To run:
```bash
$ cd $GOPATH/src/github.com/scottberke/anagram-search-go
$ ./anagram-search-go -port=3000
```

To run tests:
```bash
$ cd $GOPATH/src/github.com/scottberke/anagram-search-go
$ go test ./...
```

## Endpoints

### Anagrams
#### GET /anagrams/:words.json
Used to get anagrams for a word. Consumes a word and returns JSON of matching anagrams.

##### Request
```bash
curl -X GET \
  http://localhost:3000/anagrams/read.json \
```
##### Response 200 OK
```json
{
    "anagrams": [
        "ared",
        "daer",
        "dare",
        "dear"
    ]
}
```

### Words
#### POST /words.json
Use to add words to the anagrams dictionary. Takes a JSON array of English-language words.

##### Request
```bash
curl -X POST \
  http://localhost:3000/words.json \
  -H 'Content-Type: application/json' \
  -d '{ "words": ["read", "dear", "dare"] }'
```
##### Response 201 Created
```json

```

#### DELETE /words.json
Use to delete all contents in the dictionary.

##### Request
```bash
curl -X DELETE \
  http://localhost:3000/words.json \
```
##### Response 204 No Content
```json

```

#### DELETE /words/:word.json
Use to delete a single word from the dictionary.

##### Request
```bash
curl -X DELETE \
  http://localhost:3000/words/read.json \
```
##### Response 204 No Content
```json

```






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

***
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

***
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
