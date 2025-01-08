package controllers

import (
	"encoding/json"
	"myapp/models"
	"myapp/services"
	"net/http"
	"strconv"
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

func GetCreditLimit(w http.ResponseWriter, r *http.Request) {
	nik := r.URL.Query().Get("nik")
	salaryStr := r.URL.Query().Get("salary")
	salary, err := strconv.ParseFloat(salaryStr, 64)
	if err != nil {
		http.Error(w, "Invalid salary value", http.StatusBadRequest)
		return
	}
	creditLimit := services.GetCreditLimit(nik, salary)
	if creditLimit == nil {
		http.Error(w, "Consumer not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(creditLimit)
}

func GetCreditScore(w http.ResponseWriter, r *http.Request) {
	nik := r.URL.Query().Get("nik")
	creditScore := services.GetCreditScore(nik)
	json.NewEncoder(w).Encode(creditScore)
}

func GetConsumer(w http.ResponseWriter, r *http.Request) {
	nik := r.URL.Query().Get("nik")
	consumer, exists := services.GetConsumer(nik)
	if !exists {
		http.Error(w, "Consumer not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(consumer)
}

func UpdateConsumer(w http.ResponseWriter, r *http.Request) {
	var consumer models.Consumer
	if err := json.NewDecoder(r.Body).Decode(&consumer); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !services.UpdateConsumer(consumer) {
		http.Error(w, "Consumer not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteConsumer(w http.ResponseWriter, r *http.Request) {
	nik := r.URL.Query().Get("nik")
	if !services.DeleteConsumer(nik) {
		http.Error(w, "Consumer not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetConsumers(w http.ResponseWriter, r *http.Request) {
	consumers := services.GetConsumers()
	json.NewEncoder(w).Encode(consumers)
}

func GetConsumersByStatus(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	consumers := services.GetConsumersByStatus(status)
	json.NewEncoder(w).Encode(consumers)
}

func GetConsumersByCreditScore(w http.ResponseWriter, r *http.Request) {
	creditScoreStr := r.URL.Query().Get("credit_score")
	creditScore, err := strconv.Atoi(creditScoreStr)
	if err != nil {
		http.Error(w, "Invalid creditScore value", http.StatusBadRequest)
		return
	}
	consumers := services.GetConsumersByCreditScore(creditScore)
	json.NewEncoder(w).Encode(consumers)
}

func GetConsumersByCreditLimit(w http.ResponseWriter, r *http.Request) {
	creditLimitStr := r.URL.Query().Get("credit_limit")
	creditLimit, err := strconv.ParseFloat(creditLimitStr, 64)
	if err != nil {
		http.Error(w, "Invalid credit limit value", http.StatusBadRequest)
		return
	}
	consumers := services.GetConsumersByCreditLimit(creditLimit)
	json.NewEncoder(w).Encode(consumers)
}

func GetConsumersBySalary(w http.ResponseWriter, r *http.Request) {
	salaryStr := r.URL.Query().Get("salary")
	salary, err := strconv.ParseFloat(salaryStr, 64)
	if err != nil {
		http.Error(w, "Invalid salary value", http.StatusBadRequest)
		return
	}
	consumers := services.GetConsumersBySalary(salary)
	json.NewEncoder(w).Encode(consumers)
}
