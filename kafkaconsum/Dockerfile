FROM golang

RUN go get github.com/Shopify/sarama
RUN go get github.com/wvanbergen/kafka/consumergroup

ENV ZOOKEEPER_HOST zookeeper
ENV ZOOKEEPER_PORT 2181
ENV TOPIC test

ADD . /app
WORKDIR /app


RUN go build -o consumer *.go

ENTRYPOINT ["./consumer"]