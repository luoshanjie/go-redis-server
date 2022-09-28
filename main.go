package main

import (
	"fmt"
	"os"

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
	s := configure.NewServerProperties(filename)
	fmt.Println(s.Bind)
}
