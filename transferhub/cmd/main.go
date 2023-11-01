package main

import (
	"fmt"

	"github.com/IBM/sarama"
)

func main() {
	Consume()
}

func Consume() {
	consumer, _ := sarama.NewConsumer([]string{"localhost:9092"}, nil)

	partitionConsumer, _ := consumer.ConsumePartition("motion", 0, sarama.OffsetNewest)

	channel := partitionConsumer.Messages()

	for {
		msg := <-channel
		fmt.Println(string(msg.Value))
	}

}
