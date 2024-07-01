package main

import (
	"fmt"
	"github.com/fadegor05/linkshortee/backend/api"
	"github.com/fadegor05/linkshortee/backend/database"
	"log"
	"os"
)

func main() {
	db := database.Connect()
	port := fmt.Sprintf(":%s", os.Getenv("BACKEND_PORT"))
	server := api.NewAPIServer(port, db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
