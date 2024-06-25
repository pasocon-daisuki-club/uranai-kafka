IMAGE_NAME = "fortune-teller:local"

build: Dockerfile
	docker build --platform linux/arm64 -t $(IMAGE_NAME) .

run-aci: build
	# already retired just for example
	docker context use uranai-kafka-app
	docker run \
	-e KAFKA_EVENTHUB_ENDPOINT=uranai-kafka-namespace.servicebus.windows.net:9093 \
	-e KAFKA_EVENTHUB_CONNECTION_STRING=$(KAFKA_EVENTHUB_CONNECTION_STRING) \
	-e KAFKA_EVENTHUB_TOPIC_NAME=uranai-kafka \
	-e AOAI_RESOURCE_NAME=uranai-kafka-aoai \
	-e AOAI_DEPLOYMENT_NAME=gpt-4o \
	-e AOAI_API_VERSION=2024-02-01 \
	-e AOAI_API_KEY=$(AOAI_API_KEY) \
	ghcr.io/piroyoung/uranai-kafka:v0.0.5

