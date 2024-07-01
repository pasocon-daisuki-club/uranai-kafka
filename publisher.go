package uranai

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"github.com/IBM/sarama"
	"log"
	"time"
)

type Publisher interface {
	Publish(ctx context.Context, resultSet *ResultSet) error
}

type SaramaPublisher struct {
	c         sarama.SyncProducer
	topicName string
}

func NewSaramaPublisher(c sarama.SyncProducer, topicName string) *SaramaPublisher {
	return &SaramaPublisher{
		c:         c,
		topicName: topicName,
	}
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
			Topic: s.topicName,
			Value: sarama.StringEncoder(jsonBytes),
		}
		partition, offset, err := s.c.SendMessage(message)
		if err != nil {
			return err
		} else {
			log.Printf("> message sent to partition %d at offset %d:\n %s", partition, offset, string(jsonBytes))
		}
	}
	return nil
}

func NewEventHubSaramaConfig(connString string) *sarama.Config {
	config := sarama.NewConfig()
	config.Net.DialTimeout = 10 * time.Second
	config.Net.SASL.Enable = true
	config.Net.SASL.User = "$ConnectionString"
	config.Net.SASL.Password = connString
	config.Net.SASL.Mechanism = "PLAIN"

	config.Net.TLS.Enable = true
	config.Net.TLS.Config = &tls.Config{
		InsecureSkipVerify: true,
		ClientAuth:         0,
	}
	config.Producer.Return.Successes = true
	return config
}
