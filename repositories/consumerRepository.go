package repositories

import (
	"myapp/models"
	"sync"
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
