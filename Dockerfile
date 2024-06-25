FROM golang:1.22.4-alpine3.20 as builder
ADD . /opt/src
RUN apk add git build-base librdkafka-dev pkgconf
ENV CGO_ENABLED=1
RUN cd /opt/src && go build -tags musl -o /opt/bin/app


FROM alpine:3.20
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
