package controllers

import (
	"encoding/json"
	"net/http"
	"myapp/models"
	"myapp/services"
)

func AddTransaction(w http.ResponseWriter, r *http.Request) {
	var transaction models.Transaction
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	services.CreateTransaction(transaction)
	w.WriteHeader(http.StatusCreated)
}