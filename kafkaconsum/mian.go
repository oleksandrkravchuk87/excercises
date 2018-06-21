package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Shopify/sarama"
	"github.com/wvanbergen/kafka/consumergroup"
)

const (
	zookeeperConn = "zookeeper:2181"
	topic         = "test"
	cgroup        = "zgroup"
)

func main() {
	sarama.Logger = log.New(os.Stdout, "", log.Ltime)

	cg, err := initConsumer()
	if err != nil {
		log.Fatal(err)
	}
	defer cg.Close()
	consume(cg)
}

func initConsumer() (*consumergroup.ConsumerGroup, error) {
	config := consumergroup.NewConfig()
	config.Offsets.Initial = sarama.OffsetOldest
	config.Offsets.ProcessingTimeout = 10 * time.Second

	cg, err := consumergroup.JoinConsumerGroup(cgroup, []string{topic}, []string{zookeeperConn}, config)
	if err != nil {
		return nil, err
	}
	return cg, err
}

func consume(cg *consumergroup.ConsumerGroup) {
	for {
		select {
		case msg := <-cg.Messages():
			if msg.Topic != topic {
				continue
			}
			fmt.Println("Topic: ", msg.Topic)
			fmt.Println("Value: ", string(msg.Value))

			err := cg.CommitUpto(msg)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
