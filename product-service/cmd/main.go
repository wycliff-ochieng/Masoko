package main

import (
	"fmt"
	"log"

	"github.com/wycliff-ochieng/cmd/api"
	"github.com/wycliff-ochieng/db"
)

func main() {
	fmt.Print("Server is starting right now ....\n")

	//l := log.New(os.Stdout,"ECOMMERCE API ",log.LstdFlags)

	store, err := db.NewPostgrestore()
	if err != nil {
		log.Fatal("unable to connect to db", err)
	}
	if err := store.Init(); err != nil {
		fmt.Printf("Couldnt initialize %v", err)
	}

	server := api.NewAPIserver(":9000", store)
	server.Run()
}
