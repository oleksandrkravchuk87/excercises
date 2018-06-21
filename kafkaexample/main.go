package main

import (
	"log"
	"os"

	"github.com/Shopify/sarama"

	"bufio"
	"fmt"
)

const (
	kafkaConn = "kafka:9092"
	topic     = "test"
)

func main() {
	producer, err := initProducer()
	if err != err {
		log.Fatal(err)
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("enter a message")
		msg, _ := reader.ReadString('\n')

		publish(msg, producer)
	}

}

func initProducer() (sarama.SyncProducer, error) {
	sarama.Logger = log.New(os.Stdout, "", log.Ltime)

	config := sarama.NewConfig()
	config.Producer.Retry.Max = 5
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	prd, err := sarama.NewSyncProducer([]string{kafkaConn}, config)
	return prd, err
}

func publish(message string, producer sarama.SyncProducer) {
	msg := &sarama.ProducerMessage{Topic: topic,
		Value: sarama.StringEncoder(message),
	}
	p, o, err := producer.SendMessage(msg)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Partition: ", p)
	fmt.Println("Offset: ", o)
}
