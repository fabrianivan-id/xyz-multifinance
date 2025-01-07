package repositories

import (
	"sync"
	"myapp/models"
)

var (
	transactions = make([]models.Transaction, 0)
	mu           sync.Mutex
)

func AddTransaction(transaction models.Transaction) {
	mu.Lock()
	transactions = append(transactions, transaction)
	mu.Unlock()
}