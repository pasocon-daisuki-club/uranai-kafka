package uranai

import (
	"context"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"os"
	"testing"
)

func TestPublisher_Publish(t *testing.T) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_EVENTHUB_ENDPOINT"),
		"sasl.mechanisms":   "PLAIN",
		"security.protocol": "SASL_SSL",
		"sasl.username":     "$ConnectionString",
		"sasl.password":     os.Getenv("KAFKA_EVENTHUB_CONNECTION_STRING"),
	})
	if err != nil {
		panic(err)
	}
	defer p.Close()
	topicName := "test"
	partition := kafka.TopicPartition{Topic: &topicName, Partition: kafka.PartitionAny}

	type fields struct {
		c         *kafka.Producer
		partition *kafka.TopicPartition
	}
	type args struct {
		ctx       context.Context
		resultSet *ResultSet
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "valid case",
			fields: fields{
				c:         p,
				partition: &partition,
			},
			args: args{
				ctx: context.Background(),
				resultSet: &ResultSet{
					Results: []Result{
						{
							Rank:         1,
							Name:         "test",
							LuckyItem:    "test",
							LuckyColor:   "test",
							LuckyService: "test",
							CareerLuck:   1,
							LoveLuck:     1,
							HealthLuck:   1,
							Description:  "test",
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &ConfluentPublisher{
				c:         tt.fields.c,
				partition: tt.fields.partition,
			}
			if err := p.Publish(tt.args.ctx, tt.args.resultSet); (err != nil) != tt.wantErr {
				t.Errorf("Publish() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
