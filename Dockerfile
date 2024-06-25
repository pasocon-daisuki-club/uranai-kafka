FROM golang:1.22.4-alpine3.20 as builder
ADD . /opt/src
RUN apk add git build-base
ENV CGO_ENABLED=1
ENV GOOS=linux
ENV GOARCH=amd64
RUN cd /opt/src && go build -o /opt/bin/app


FROM alpine:3.20
ENV KAFKA_EVENTHUB_ENDPOINT="kafka:9092"
ENV KAFKA_EVENTHUB_CONNECTION_STRING="kafka:9092"
ENV KAFKA_EVENTHUB_TOPIC_NAME="test"
ENV AOAI_RESOURCE_NAME="test"
ENV AOAI_DEPLOYMENT_NAME="test"
ENV AOAI_API_VERSION="v1"
ENV AOAI_API_KEY="test"

COPY --from=builder /opt/bin/app /opt/bin/app
RUN chmod +x /opt/bin/fortune_teller
CMD ["/opt/bin/app"]
