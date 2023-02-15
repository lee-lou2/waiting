package main

import (
	"os"
	"waiting/cmd/client"
	"waiting/cmd/server"
	"waiting/config/env"
)

func main() {
	// 환경 변수 실행
	env.Load()

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

//func main() {
//	// 환경 변수 실행
//	env.Load()
//
//	// 타입별 프로그램 실행
//	ProjectType := os.Getenv("PROJECT_TYPE")
//	if ProjectType == "SERVER" {
//		server.Run()
//	} else if ProjectType == "CLIENT" {
//		client.Run()
//	}
//}
