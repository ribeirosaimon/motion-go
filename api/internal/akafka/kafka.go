package akafka

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Producer() {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "host.docker.internal:9094",
		"group.id":          "motion-go",
		"auto.offset.reset": "earliest",
	})
	topic := "teste"
	partition := kafka.TopicPartition{
		Topic:     &topic,
		Partition: kafka.PartitionAny,
	}

	if err != nil {
		panic(err)
	}
	for i := 0; i < 10; i++ {
		value := fmt.Sprintf("teste mensagem %d", i)
		err := producer.Produce(&kafka.Message{
			TopicPartition: partition,
			Value:          []byte(value),
		}, nil)

		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(value)
	}
	producer.Flush(15 * 1000)
	producer.Close()
}
