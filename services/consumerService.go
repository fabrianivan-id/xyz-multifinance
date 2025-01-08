package services

import (
	"myapp/models"
	"myapp/repositories"
)

func CreateConsumer(consumer models.Consumer) {
	repositories.AddConsumer(consumer)
}

func GetCreditLimit(nik string, salary float64) map[int]float64 {
	return repositories.GetCreditLimit(nik, salary)
}

func GetCreditScore(nik string) int {
	return repositories.GetCreditScore(nik)
}

func GetConsumer(nik string) (models.Consumer, bool) {
	return repositories.GetConsumerByNIK(nik)
}

func UpdateConsumer(consumer models.Consumer) bool {
	return repositories.UpdateConsumer(consumer)
}

func DeleteConsumer(nik string) bool {
	return repositories.DeleteConsumer(nik)
}

func GetConsumers() []models.Consumer {
	return repositories.GetAllConsumers()
}

func GetConsumersByStatus(status string) []models.Consumer {
	return repositories.GetConsumersByStatus(status)
}

func GetConsumersByCreditScore(score int) []models.Consumer {
	return repositories.GetConsumersByCreditScore(score)
}

func GetConsumersBySalary(salary float64) []models.Consumer {
	return repositories.GetConsumersBySalary(salary)
}

func GetConsumersByCreditLimit(amount float64) []models.Consumer {
	return repositories.GetConsumersByLimit(amount)
}

func GetConsumersByStatusAndCreditScore(status string, score int) []models.Consumer {
	return repositories.GetConsumersByStatusAndCreditScore(status, score)
}

func GetConsumersByStatusAndSalary(status string, salary float64) []models.Consumer {
	return repositories.GetConsumersByStatusAndSalary(status, salary)
}
