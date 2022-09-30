package main

import (
	"fmt"
	"os"

	"go-redis-server/internal/delivery"
	"go-redis-server/internal/service"
	"go-redis-server/pkg/configure"
)

var filename string
var banner = `
 ██████   ██████      ██████  ███████ ██████  ██ ███████     ███████ ███████ ██████  ██    ██ ███████ ██████  
██       ██    ██     ██   ██ ██      ██   ██ ██ ██          ██      ██      ██   ██ ██    ██ ██      ██   ██ 
██   ███ ██    ██     ██████  █████   ██   ██ ██ ███████     ███████ █████   ██████  ██    ██ █████   ██████  
██    ██ ██    ██     ██   ██ ██      ██   ██ ██      ██          ██ ██      ██   ██  ██  ██  ██      ██   ██ 
 ██████   ██████      ██   ██ ███████ ██████  ██ ███████     ███████ ███████ ██   ██   ████   ███████ ██   ██
`

func init() {
	current, _ := os.Getwd()
	filename = fmt.Sprintf("%s/conf/redis.conf", current)
}

func main() {
	print(banner)
	properties := configure.NewServerProperties(filename)
	tcp := delivery.NewTCPDelivery(properties)
	print(service.NewBaseService(tcp).Run())
}
