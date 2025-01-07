package controllers

import (
	"encoding/json"
	"net/http"
	"myapp/models"
	"myapp/services"
)

func AddConsumer(w http.ResponseWriter, r *http.Request) {
	var consumer models.Consumer
	if err := json.NewDecoder(r.Body).Decode(&consumer); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	services.CreateConsumer(consumer)
	w.WriteHeader(http.StatusCreated)
}