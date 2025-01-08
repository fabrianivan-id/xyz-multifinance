package services

import (
	"myapp/models"
	"myapp/repositories"
)

func CreateTransaction(transaction models.Transaction) {
	repositories.AddTransaction(transaction)
}

func GetTransactions() []models.Transaction {
	return repositories.GetTransactions()
}

func GetTransactionByContractNumber(contractNumber string) (models.Transaction, bool) {
	return repositories.GetTransactionByContractNumber(contractNumber)
}

func UpdateTransaction(transaction models.Transaction) bool {
	return repositories.UpdateTransaction(transaction)
}

func DeleteTransaction(contractNumber string) bool {
	return repositories.DeleteTransaction(contractNumber)
}

func GetTransactionsByNIK(nik string) []models.Transaction {
	return repositories.GetTransactionsByNIK(nik)
}

func GetTransactionsByStatus(status string) []models.Transaction {
	return repositories.GetTransactionsByStatus(status)
}
