package uranai

import (
	"context"
	"encoding/json"
	"github.com/IBM/sarama"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"log"
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

type SaramaPublisher struct {
	c sarama.SyncProducer
}

func (s SaramaPublisher) Publish(ctx context.Context, resultSet *ResultSet) error {
	for _, result := range resultSet.Results {
		// Convert result to JSON
		jsonBytes, err := json.Marshal(result)
		if err != nil {
			return err
		}
		// Publish to Kafka
		message := &sarama.ProducerMessage{
			Topic: "uranai-kafka", //FIXME: Hardcoded topic
			Value: sarama.StringEncoder(jsonBytes),
		}
		partition, offset, err := s.c.SendMessage(message)
		if err != nil {
			return err
		} else {
			log.Printf("> message sent to partition %d at offset %d\n", partition, offset)
		}
	}
	return nil
}
