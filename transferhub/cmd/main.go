package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	fmt.Println("hello word")
	Consume()
}

func Consume() {
	kafkaConsumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "host.docker.internal:9094",
		"group.id":          "motion-go",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		panic(err)
	}
	kafkaConsumer.SubscribeTopics([]string{"teste"}, nil)
	for {
		msg, err := kafkaConsumer.ReadMessage(-1)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(msg)

	}
}
