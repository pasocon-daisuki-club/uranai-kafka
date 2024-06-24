package uranai

import (
	"context"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Publisher struct {
	c *kafka.Producer
}

func (c *Publisher) Publish(ctx context.Context, result *ResultSet) error {
	return nil
}
