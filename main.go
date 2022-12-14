package main

import (
	"go-redis-server/internal/service"
)

var banner = `
 ██████   ██████      ██████  ███████ ██████  ██ ███████     ███████ ███████ ██████  ██    ██ ███████ ██████  
██       ██    ██     ██   ██ ██      ██   ██ ██ ██          ██      ██      ██   ██ ██    ██ ██      ██   ██ 
██   ███ ██    ██     ██████  █████   ██   ██ ██ ███████     ███████ █████   ██████  ██    ██ █████   ██████  
██    ██ ██    ██     ██   ██ ██      ██   ██ ██      ██          ██ ██      ██   ██  ██  ██  ██      ██   ██ 
 ██████   ██████      ██   ██ ███████ ██████  ██ ███████     ███████ ███████ ██   ██   ████   ███████ ██   ██
`

func main() {
	print(banner)
	print(service.NewService().Run())
}
