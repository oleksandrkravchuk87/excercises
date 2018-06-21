FROM golang

RUN go get github.com/Shopify/sarama


ENV KAFKA_HOST localhost
ENV KAFKA_PORT 9092
ENV KAFKA_TOPIC test

ADD . /app
WORKDIR /app


RUN go build -o produzer *.go

ENTRYPOINT ["./produzer"]