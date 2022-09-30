# Go Redis Server
This project is use native golang implementation redis server func. use model include goroutine-per-connection, clean-arch.

## Support Func
- Support string, list, hash, set, sorted set, bitmap
- TTL
- Publish/Subscribe

## Get Started
You will need install docker, and clone project (docker version > 18.0)
```shell
$ git clone https://github.com/luoshanjie/go-redis-server.git
$ cd go-redis-server
$ docker build -t go-redis-server:1.0 .
$ docker run -itd container-redis -p 6339:6339 go-redis-server:1.0
```
Use redis client connect go-redis-server
```shell
$ rediscli -h 127.0.0.1 -p 6339
```


# License
This project is licensed under the [Apache License 2.0](https://github.com/luoshanjie/go-redis-server/blob/master/LICENSE)
