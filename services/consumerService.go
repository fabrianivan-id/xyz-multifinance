package services

import (
	"myapp/models"
	"myapp/repositories"
)

func CreateConsumer(consumer models.Consumer) {
	repositories.AddConsumer(consumer)
}