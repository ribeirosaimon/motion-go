package akafka

import (
	"fmt"

	"github.com/IBM/sarama"
	"github.com/magiconair/properties"
)

type KafkaMotion struct{}

var kafkamotion *KafkaMotion

func NewKafkaMotion() *KafkaMotion {
	if kafkamotion == nil {
		kafkamotion = &KafkaMotion{}
	}
	return kafkamotion
}

func (k *KafkaMotion) CreateConn(prop *properties.Properties) <-chan *sarama.ConsumerMessage {
	host := prop.GetString("kafka.host", "localhost")
	port := prop.GetString("kafka.port", "9092")

	consumer, err := sarama.NewConsumer([]string{fmt.Sprintf("%s:%s", host, port)}, nil)
	if err != nil {
		panic(err)
	}

	partitionConsumer, _ := consumer.ConsumePartition("motion", 0, sarama.OffsetNewest)

	return partitionConsumer.Messages()
}
