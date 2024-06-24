IMAGE_NAME = "fortune-teller:local"

build: Dockerfile
	docker build -t $(IMAGE_NAME) .
