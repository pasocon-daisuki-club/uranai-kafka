IMAGE_NAME = "fortune-teller:local"

build: Dockerfile
	docker build -t $(IMAGE_NAME) .

run: build
	docker run $(IMAGE_NAME)
