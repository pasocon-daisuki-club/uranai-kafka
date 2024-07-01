package uranai

import (
	"context"
	"crypto/tls"
	"github.com/IBM/sarama"
	"os"
	"testing"
	"time"
)

func TestSaramaPublisher_Publish(t *testing.T) {
	config := sarama.NewConfig()
	config.Net.DialTimeout = 10 * time.Second
	config.Net.SASL.Enable = true
	config.Net.SASL.User = "$ConnectionString"
	config.Net.SASL.Password = os.Getenv("KAFKA_EVENTHUB_CONNECTION_STRING")
	config.Net.SASL.Mechanism = "PLAIN"

	config.Net.TLS.Enable = true
	config.Net.TLS.Config = &tls.Config{
		InsecureSkipVerify: true,
		ClientAuth:         0,
	}
	config.Producer.Return.Successes = true

	addr := os.Getenv("KAFKA_EVENTHUB_ENDPOINT")
	producer, err := sarama.NewSyncProducer([]string{addr}, config)
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	type fields struct {
		c         sarama.SyncProducer
		topicName string
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
				c:         producer,
				topicName: "test",
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
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SaramaPublisher{
				c:         tt.fields.c,
				topicName: tt.fields.topicName,
			}
			if err := s.Publish(tt.args.ctx, tt.args.resultSet); (err != nil) != tt.wantErr {
				t.Errorf("Publish() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
