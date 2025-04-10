package main

import (
	"Auth-microservice/db"
	"Auth-microservice/router"
	"fmt"
	"log"
)

func main() {
	// init database with grom
	if err := db.Init(); err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	// init router
	router.InitRouter()
}