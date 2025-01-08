package repositories

import (
	"myapp/models"
	"sync"
	"time"
)

var (
	consumers  = make(map[string]models.Consumer)
	consumerMu sync.Mutex
)

func AddConsumer(consumer models.Consumer) {
	consumerMu.Lock()
	consumers[consumer.NIK] = consumer
	consumerMu.Unlock()
}

func GetCreditLimit(nik string, salary float64) map[int]float64 {
	consumerMu.Lock()
	defer consumerMu.Unlock()

	consumer, exists := consumers[nik]
	if !exists {
		return nil
	}

	// Calculate risk factor based on salary
	riskFactor := 1.0
	if salary <= 3000000 {
		riskFactor = 0.4
	} else if salary <= 5000000 {
		riskFactor = 0.6
	} else if salary <= 12000000 {
		riskFactor = 0.75
	} else {
		riskFactor = 0.9
	}

	// Adjust credit limit based on risk factor
	adjustedLimit := make(map[int]float64)
	for tenor, limit := range consumer.Limit {
		adjustedLimit[tenor] = limit * riskFactor
	}

	return adjustedLimit
}

func GetCreditScore(nik string) int {
	consumerMu.Lock()
	defer consumerMu.Unlock()

	_, exists := consumers[nik]
	if !exists {
		return 0
	}

	// Calculate credit score based on the number of transactions
	creditScore := 0
	for _, transaction := range transactions {
		if transaction.NIK == nik {
			creditScore++
		}
	}

	return creditScore
}

func CheckCreditLimit(nik string, amount float64) []float64 {
	consumerMu.Lock()
	defer consumerMu.Unlock()

	consumer, exists := consumers[nik]
	if !exists {
		return []float64{}
	}

	// Calculate the total OTR used monthly before admin fee for transactions with status "Done"
	totalOTRUsed := 0.0
	for _, transaction := range transactions {
		if transaction.Date.Month() == time.Now().Month() && transaction.Date.Year() == time.Now().Year() && transaction.Status == "Done" {
			totalOTRUsed += transaction.OTR
		}
	}

	// Check if the remaining limit is sufficient for the amount
	var remainingLimit []float64
	for tenor, limit := range consumer.Limit {
		if tenor == 6 {
			remainingLimit = append(remainingLimit, limit-totalOTRUsed)
		}

		if tenor == 3 {
			remainingLimit = append(remainingLimit, limit-totalOTRUsed/2)
		}

		if tenor == 2 {
			remainingLimit = append(remainingLimit, limit-totalOTRUsed/3)
		}

		if tenor == 1 {
			remainingLimit = append(remainingLimit, limit-totalOTRUsed/6)
		}

	}

	return remainingLimit
}

func GetConsumerByNIK(nik string) (models.Consumer, bool) {
	consumerMu.Lock()
	defer consumerMu.Unlock()

	consumer, exists := consumers[nik]
	return consumer, exists
}

func UpdateConsumer(consumer models.Consumer) bool {
	consumerMu.Lock()
	defer consumerMu.Unlock()

	_, exists := consumers[consumer.NIK]
	if exists {
		consumers[consumer.NIK] = consumer
	}

	return exists
}

func DeleteConsumer(nik string) bool {
	consumerMu.Lock()
	defer consumerMu.Unlock()

	_, exists := consumers[nik]
	if exists {
		delete(consumers, nik)
	}

	return exists
}

func GetAllConsumers() []models.Consumer {
	consumerMu.Lock()
	defer consumerMu.Unlock()

	result := make([]models.Consumer, 0)
	for _, consumer := range consumers {
		result = append(result, consumer)
	}

	return result
}

func GetConsumersByStatus(status string) []models.Consumer {
	consumerMu.Lock()
	defer consumerMu.Unlock()

	result := make([]models.Consumer, 0)
	for _, consumer := range consumers {
		if consumer.Status == status {
			result = append(result, consumer)
		}
	}

	return result
}

func GetConsumersByCreditScore(score int) []models.Consumer {
	consumerMu.Lock()
	defer consumerMu.Unlock()

	result := make([]models.Consumer, 0)
	for _, consumer := range consumers {
		if consumer.CreditScore == score {
			result = append(result, consumer)
		}
	}

	return result
}

func GetConsumersByCreditScoreAndStatus(score int, status string) []models.Consumer {
	consumerMu.Lock()
	defer consumerMu.Unlock()

	result := make([]models.Consumer, 0)
	for _, consumer := range consumers {
		if consumer.CreditScore == score && consumer.Status == status {
			result = append(result, consumer)
		}
	}

	return result
}

func GetConsumersBySalary(salary float64) []models.Consumer {
	consumerMu.Lock()
	defer consumerMu.Unlock()

	result := make([]models.Consumer, 0)
	for _, consumer := range consumers {
		if consumer.Salary == salary {
			result = append(result, consumer)
		}
	}

	return result
}

func GetConsumersByLimit(amount float64) []models.Consumer {
	consumerMu.Lock()
	defer consumerMu.Unlock()

	result := make([]models.Consumer, 0)
	for _, consumer := range consumers {
		for _, limit := range consumer.Limit {
			if limit == amount {
				result = append(result, consumer)
				break
			}
		}
	}

	return result
}

func GetConsumersByStatusAndCreditScore(status string, score int) []models.Consumer {
	consumerMu.Lock()
	defer consumerMu.Unlock()

	result := make([]models.Consumer, 0)
	for _, consumer := range consumers {
		if consumer.Status == status && consumer.CreditScore == score {
			result = append(result, consumer)
		}
	}

	return result
}

func GetConsumersByStatusAndSalary(status string, salary float64) []models.Consumer {
	consumerMu.Lock()
	defer consumerMu.Unlock()

	result := make([]models.Consumer, 0)
	for _, consumer := range consumers {
		if consumer.Status == status && consumer.Salary == salary {
			result = append(result, consumer)
		}
	}

	return result
}

func GetConsumersByStatusAndLimit(status string, tenor int, amount float64) []models.Consumer {
	consumerMu.Lock()
	defer consumerMu.Unlock()

	result := make([]models.Consumer, 0)
	for _, consumer := range consumers {
		if consumer.Status == status && consumer.Limit[tenor] == amount {
			result = append(result, consumer)
		}
	}

	return result
}
