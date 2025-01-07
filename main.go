package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"myapp/config"
	"myapp/controllers"
)

func main() {
	config.LoadConfig()

	router := mux.NewRouter()

	router.HandleFunc("/consumers", controllers.AddConsumer).Methods("POST")
	router.HandleFunc("/transactions", controllers.AddTransaction).Methods("POST")

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}