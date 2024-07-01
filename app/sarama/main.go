package main

import (
	"context"
	"github.com/IBM/sarama"
	"github.com/piroyoung/uranai-kafka"
	"net/http"
	"os"
	"strconv"
)

func main() {
	// kafka parameters
	hostName := os.Getenv("KAFKA_EVENTHUB_ENDPOINT")
	connString := os.Getenv("KAFKA_EVENTHUB_CONNECTION_STRING")
	topicName := os.Getenv("KAFKA_EVENTHUB_TOPIC_NAME")

	// aoai parameters
	resourceName := os.Getenv("AOAI_RESOURCE_NAME")
	deploymentName := os.Getenv("AOAI_DEPLOYMENT_NAME")
	apiVersion := os.Getenv("AOAI_API_VERSION")
	accessToken := os.Getenv("AOAI_API_KEY")
	temperature, err := strconv.ParseFloat(os.Getenv("AOAI_TEMPERATURE"), 32)
	if err != nil {
		panic(err)
	}
	config := uranai.NewEventHubSaramaConfig(connString)
	producer, err := sarama.NewSyncProducer([]string{hostName}, config)
	if err != nil {
		panic(err)
	}
	defer producer.Close()
	publisher := uranai.NewSaramaPublisher(producer, topicName)

	httpClient := &http.Client{}
	client := uranai.NewClient(httpClient, resourceName, deploymentName, apiVersion, accessToken)
	fortuneTeller := uranai.NewFortuneTeller(client, float32(temperature))

	batch := uranai.NewBatch(fortuneTeller, publisher)
	ctx := context.Background()

	err = batch.Run(ctx)
	if err != nil {
		panic(err)
	}
}
