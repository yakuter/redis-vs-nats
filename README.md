# NATS vs REDIS
This repo is prepared to test the performance of nats and redis in pubsub messaging.

## PREQUISITES
Redis server - https://redis.io/  
Nats server - https://nats.io/download/nats-io/nats-server/  
NATS streaming server - https://nats.io/download/nats-io/nats-streaming-server/


## FOLDERS
### faker
With faker, the user can generate json files with fake messages. Usage:
```go
    // In faker folder
    go run faker.go
```
### json
JSON files generated with faker are located in this folder. Default files and sizes are listed below:
```shell
➜  json git:(master) ✗ ls -lah
total 244896
drwxr-xr-x   5 yakuter  staff   160B Apr 10 10:47 .
drwxr-xr-x  11 yakuter  staff   352B Apr 10 10:53 ..
-rw-r--r--   1 yakuter  staff   108M Apr 10 10:47 100k.json
-rw-r--r--   1 yakuter  staff    11M Apr 10 10:47 10k.json
-rw-r--r--   1 yakuter  staff   1.1M Apr 10 10:47 1k.json
```
### redis, nats and nats-streaming
These folders have same structure. They all have a sender and a receiver. The test steps are below. The order is important!

1. Start the relevant server according to the test (Redis, Nats or Nats Streaming Server).
2. Go into receiver folder of the test and run receiver with `go run receiver.go`
3. Go into sender folder of the test and run sender with `go run sender.go`

### monitoring
#### Nats
NATS has a great tool for monioring called `nats-top` which is a `top`-like tool for monitoring NATS servers.
https://github.com/nats-io/nats-top

```sh
$ nats-top

NATS server version 0.7.3 (uptime: 3m34s)
Server:
  Load: CPU:  58.3%  Memory: 8.6M  Slow Consumers: 0
  In:   Msgs: 568.7K  Bytes: 1.7M  Msgs/Sec: 13129.0  Bytes/Sec: 38.5K
  Out:  Msgs: 1.6M  Bytes: 4.7M  Msgs/Sec: 131290.9  Bytes/Sec: 384.6K    

Connections: 10
  HOST                 CID    NAME        SUBS    PENDING     MSGS_TO   MSGS_FROM   BYTES_TO    BYTES_FROM  LANG     VERSION  UPTIME   LAST ACTIVITY
  127.0.0.1:57487      13     example     1       12.0K       161.6K    0           484.7K      0           go       1.1.7    17s      2016-02-09 00:13:24.753062715 -0800 PST
  127.0.0.1:57488      14     example     1       11.9K       161.6K    0           484.7K      0           go       1.1.7    17s      2016-02-09 00:13:24.753040168 -0800 PST
  127.0.0.1:57489      15     example     1       12.1K       161.6K    0           484.7K      0           go       1.1.7    17s      2016-02-09 00:13:24.753069442 -0800 PST
  127.0.0.1:57490      16     example     1       12.0K       161.6K    0           484.7K      0           go       1.1.7    17s      2016-02-09 00:13:24.753057413 -0800 PST
```

#### Redis
Unfortunately, couldn't find a useful tool like nats-top. If you know any, I would like to hear.

### My Test Results

**File**: 100k.json | 100.000 JSON messages | 108mb

**NATS**:   6.207s  
**REDIS**:  8.212s
