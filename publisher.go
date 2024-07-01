package uranai

import (
	"context"
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Publisher interface {
	Publish(ctx context.Context, resultSet *ResultSet) error
}

type ConfluentPublisher struct {
	c         *kafka.Producer
	partition *kafka.TopicPartition
}

func NewConfluentPublisher(c *kafka.Producer, partition *kafka.TopicPartition) *ConfluentPublisher {
	return &ConfluentPublisher{c: c, partition: partition}
}

func (p *ConfluentPublisher) Publish(ctx context.Context, resultSet *ResultSet) error {

	for _, result := range resultSet.Results {
		// Convert result to JSON
		jsonBytes, err := json.Marshal(result)
		if err != nil {
			return err
		}
		// Publish to Kafka
		message := &kafka.Message{
			TopicPartition: *p.partition,
			Value:          jsonBytes,
		}
		err = p.c.Produce(message, nil)
		if err != nil {
			return err
		}
	}
	p.c.Flush(15 * 1000)
	return nil
}
