package main

import (
	"os"
	"waiting/cmd/client"
	"waiting/cmd/server"
)

func main() {
	args := os.Args
	if len(args) == 2 {
		arg := args[1]
		if arg == "server" {
			server.Run()
		} else {
			client.Run()
		}
	}
}
