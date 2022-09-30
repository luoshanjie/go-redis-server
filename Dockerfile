# build
FROM golang:1.9 as builder
WORKDIR /go-redis-server
ENV GO111MODULE "on"
ENV GOPROXY=https://goproxy.cn,direct
RUN go env
COPY . .
RUN go mod download
RUN make build

# runtime
FROM alpine
WORKDIR /app
COPY ./conf/ ./conf/
COPY --from=builder /go-redis-server ./bin/
RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo Asia/Shanghai > /etc/timezone
ENTRYPOINT ["/app/bin/rediserver"]
