package main

import "github.com/wycliff-ochieng/cart-service/cmd/api"

func main() {
	server := api.NewAPIServer(":9000")
	server.Run()
}
