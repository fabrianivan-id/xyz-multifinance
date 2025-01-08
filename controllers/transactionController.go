package controllers

import (
	"encoding/json"
	"myapp/models"
	"myapp/services"
	"net/http"
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

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	transactions := services.GetTransactions()
	json.NewEncoder(w).Encode(transactions)
}

func GetTransactionByContractNumber(w http.ResponseWriter, r *http.Request) {
	contractNumber := r.URL.Query().Get("contract_number")
	transaction, exists := services.GetTransactionByContractNumber(contractNumber)
	if !exists {
		http.Error(w, "Transaction not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(transaction)
}

func UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	var transaction models.Transaction
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !services.UpdateTransaction(transaction) {
		http.Error(w, "Transaction not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	contractNumber := r.URL.Query().Get("contract_number")
	if !services.DeleteTransaction(contractNumber) {
		http.Error(w, "Transaction not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetTransactionsByNIK(w http.ResponseWriter, r *http.Request) {
	nik := r.URL.Query().Get("nik")
	transactions := services.GetTransactionsByNIK(nik)
	json.NewEncoder(w).Encode(transactions)
}

func GetTransactionsByStatus(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	transactions := services.GetTransactionsByStatus(status)
	json.NewEncoder(w).Encode(transactions)
}

func GetTransactionsByDueDate(w http.ResponseWriter, r *http.Request) {
	dueDate := r.URL.Query().Get("due_date")
	transactions := services.GetTransactionsByDueDate(dueDate)
	json.NewEncoder(w).Encode(transactions)
}

func GetTransactionsByAssetName(w http.ResponseWriter, r *http.Request) {
	assetName := r.URL.Query().Get("asset_name")
	transactions := services.GetTransactionsByAssetName(assetName)
	json.NewEncoder(w).Encode(transactions)
}
