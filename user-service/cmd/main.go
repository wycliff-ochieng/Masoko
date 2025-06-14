package main

import "github.com/wycliff-ochieng/user-service/cmd/api"

func main() {
	server := api.NewAPIServer(":8000")
	server.Run()
}
