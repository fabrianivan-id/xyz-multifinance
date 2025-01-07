package repositories

import (
	"sync"
	"myapp/models"
)

var (
	consumers = make(map[string]models.Consumer)
	mu        sync.Mutex
)

func AddConsumer(consumer models.Consumer) {
	mu.Lock()
	consumers[consumer.NIK] = consumer
	mu.Unlock()
}