IMAGE_NAME = "uranai-kafka:local"

build: Dockerfile .dockerignore
	docker build -t $(IMAGE_NAME) .

run: build
	# already retired just for example
	docker run \
	-it \
	-e KAFKA_EVENTHUB_ENDPOINT=uranai-kafka-namespace.servicebus.windows.net:9093 \
	-e KAFKA_EVENTHUB_CONNECTION_STRING=$(KAFKA_EVENTHUB_CONNECTION_STRING) \
	-e KAFKA_EVENTHUB_TOPIC_NAME=uranai-kafka \
	-e AOAI_RESOURCE_NAME=uranai-kafka-aoai \
	-e AOAI_DEPLOYMENT_NAME=gpt-4o \
	-e AOAI_API_VERSION=2024-02-01 \
	-e AOAI_API_KEY=$(AOAI_API_KEY) \
	uranai-kafka:local

