package main

import (
	"friends_management/cmd/handlers"
	"friends_management/utils"
	"log"
	"net/http"
)

func main() {
	db := utils.CreateConnection()
	defer db.Close()

	log.Println("Successfully connected!")

	r := handlers.API(db)
	log.Println("Server started on: http://localhost:3000")
	http.ListenAndServe(":3000", r)
}
