# go-redis-server
this project is use native golang implementation redis server func

## support func
- Support string, list, hash, set, sorted set, bitmap
- TTL
- Publish/Subscribe

## get started
you will need install docker, and clone project
```shell
$ docker version
$ git clone https://github.com/luoshanjie/go-redis-server.git
$ cd go-redis-server
$ docker build -t go-redis-server:1.0 .
$ docker run -itd container-redis -p 6339:6339 go-redis-server:1.0
```


# License
This project is licensed under the [Apache license](https://github.com/luoshanjie/go-redis-server/blob/master/LICENSE)
