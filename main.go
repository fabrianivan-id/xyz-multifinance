package main

import (
	"log"
	"net/http"

	"myapp/config"
	"myapp/controllers"

	"github.com/gorilla/mux"
)

func main() {
	config.LoadConfig()

	router := mux.NewRouter()

	router.HandleFunc("/consumers", controllers.AddConsumer).Methods("POST")
	router.HandleFunc("/transactions", controllers.AddTransaction).Methods("POST")
	router.HandleFunc("/consumers/credit_limit", controllers.GetCreditLimit).Methods("GET")
	router.HandleFunc("/consumers/credit_score", controllers.GetCreditScore).Methods("GET")
	router.HandleFunc("/consumers", controllers.GetConsumer).Methods("GET")
	router.HandleFunc("/consumers", controllers.UpdateConsumer).Methods("PUT")
	router.HandleFunc("/transactions", controllers.GetTransactions).Methods("GET")
	router.HandleFunc("/transactions/contract_number", controllers.GetTransactionByContractNumber).Methods("GET")
	router.HandleFunc("/transactions", controllers.UpdateTransaction).Methods("PUT")
	router.HandleFunc("/transactions", controllers.DeleteTransaction).Methods("DELETE")
	router.HandleFunc("/transactions/nik", controllers.GetTransactionsByNIK).Methods("GET")
	router.HandleFunc("/transactions/status", controllers.GetTransactionsByStatus).Methods("GET")

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
