package services

import (
	"myapp/models"
	"myapp/repositories"
)

func CreateTransaction(transaction models.Transaction) {
	repositories.AddTransaction(transaction)
}