FROM golang:1.22.4-alpine3.20 as builder
ADD . /opt/src
RUN apk add git build-base
ENV CGO_ENABLED=1
RUN cd /opt/src && go build -o /opt/bin/fortune_teller


FROM alpine:3.20
ENV KAFKA_EVENTHUB_ENDPOINT="kafka:9092"
ENV KAFKA_EVENTHUB_CONNECTION_STRING="kafka:9092"
ENV KAFKA_EVENTHUB_TOPIC_NAME="test"
ENV AOAI_RESOURCE_NAME="test"
ENV AOAI_DEPLOYMENT_NAME="test"
ENV AOAI_API_VERSION="v1"
ENV AOAI_API_KEY="test"

COPY --from=builder /opt/bin/fortune_teller /opt/bin/fortune_teller
CMD ["/opt/bin/fortune_teller"]
