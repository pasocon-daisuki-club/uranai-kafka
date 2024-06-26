FROM golang:1.22.4-bookworm as builder
ADD . /opt/src
RUN apt update
RUN apt install -y git build-essential librdkafka-dev
ENV CGO_ENABLED=1
RUN cd /opt/src && go build -o /opt/bin/app


FROM debian:bookworm
ENV KAFKA_EVENTHUB_ENDPOINT="kafka:9093"
ENV KAFKA_EVENTHUB_CONNECTION_STRING="conn-str"
ENV KAFKA_EVENTHUB_TOPIC_NAME="test"
ENV AOAI_RESOURCE_NAME="test"
ENV AOAI_DEPLOYMENT_NAME="test"
ENV AOAI_API_VERSION="v1"
ENV AOAI_API_KEY="your-key"
ENV AOAI_TEMPERATURE="0.7"

COPY --from=builder /opt/bin/app /opt/bin/app
RUN chmod +x /opt/bin/app
CMD ["/opt/bin/app"]
