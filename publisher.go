package uranai

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Publisher struct {
	c *kafka.Producer
}

func (c *Publisher) publish(result *ResultSet) error {
	return nil
}
