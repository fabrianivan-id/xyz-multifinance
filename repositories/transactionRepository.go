package repositories

import (
	"myapp/models"
	"sync"
)

var (
	transactions = make([]models.Transaction, 0)
	mu           sync.Mutex
)

func AddTransaction(transaction models.Transaction) {
	mu.Lock()
	defer mu.Unlock()

	// Calculate the fee as 0.01% of the OTR (On The Road) price
	fee := transaction.OTR * 0.0001
	insuranceFee := transaction.OTR * 0.0005
	transaction.AdminFee = fee + insuranceFee + 15.000

	// Calculate the daily rate of 0.015% and apply it to the tenor to get the monthly payment
	dailyRate := 0.00015
	monthlyRate := dailyRate * 30
	transaction.Tenor = int(transaction.OTR*monthlyRate*float64(transaction.Tenor)) + 15.000

	// Add the transaction to the list
	transactions = append(transactions, transaction)
}

func GetTransactions() []models.Transaction {
	mu.Lock()
	defer mu.Unlock()

	return transactions
}

func GetTransactionByContractNumber(contractNumber string) (models.Transaction, bool) {
	mu.Lock()
	defer mu.Unlock()

	for _, transaction := range transactions {
		if transaction.ContractNumber == contractNumber {
			return transaction, true
		}
	}

	return models.Transaction{}, false
}

func UpdateTransaction(transaction models.Transaction) bool {
	mu.Lock()
	defer mu.Unlock()

	for i, t := range transactions {
		if t.ContractNumber == transaction.ContractNumber {
			transactions[i] = transaction
			return true
		}
	}

	return false
}

func DeleteTransaction(contractNumber string) bool {
	mu.Lock()
	defer mu.Unlock()

	for i, transaction := range transactions {
		if transaction.ContractNumber == contractNumber {
			transactions = append(transactions[:i], transactions[i+1:]...)
			return true
		}
	}

	return false
}

func GetTransactionsByNIK(nik string) []models.Transaction {
	mu.Lock()
	defer mu.Unlock()

	result := make([]models.Transaction, 0)
	for _, transaction := range transactions {
		if transaction.NIK == nik {
			result = append(result, transaction)
		}
	}

	return result
}

func GetTransactionsByAssetName(assetName string) []models.Transaction {
	mu.Lock()
	defer mu.Unlock()

	result := make([]models.Transaction, 0)
	for _, transaction := range transactions {
		if transaction.AssetName == assetName {
			result = append(result, transaction)
		}
	}

	return result
}

func GetCreditUsed(nik string) float64 {
	mu.Lock()
	defer mu.Unlock()

	var creditUsed float64
	for _, transaction := range transactions {
		if transaction.NIK == nik {
			creditUsed += transaction.OTR
		}
	}

	return creditUsed
}

func GetTransactionsByStatus(status string) []models.Transaction {
	mu.Lock()
	defer mu.Unlock()

	result := make([]models.Transaction, 0)
	for _, transaction := range transactions {
		if transaction.Status == status {
			result = append(result, transaction)
		}
	}

	return result
}
