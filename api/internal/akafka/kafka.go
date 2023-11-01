package akafka

import (
	"encoding/json"
	"fmt"

	"github.com/IBM/sarama"
	"github.com/magiconair/properties"
)

type MotionKafka struct {
	producer    sarama.SyncProducer
	senderTopic string
}

var kafkaProducer *MotionKafka

func NewMotionKafka(configurations properties.Properties) *MotionKafka {
	if kafkaProducer != nil {
		return kafkaProducer
	}

	host := configurations.GetString("kafka.host", "")
	port := configurations.GetString("kafka.port", "")
	topic := configurations.GetString("kafka.topic", "")

	kafkaHost := fmt.Sprintf("%s:%s", host, port)

	producer, err := getProducer(kafkaHost)
	if err != nil {
		syncProducer, _ := getProducer("localhost:9092")
		kafkaProducer = &MotionKafka{producer: syncProducer, senderTopic: "default"}
	}
	kafkaProducer = &MotionKafka{producer: producer, senderTopic: topic}
	return kafkaProducer
}

func GetMotionKafka() *MotionKafka {
	return kafkaProducer
}

func (k *MotionKafka) SendMessage(message interface{}) error {
	defer k.producer.Close()
	jsonData, err := json.Marshal(message)
	if err != nil {
		return err
	}
	msg := getMessage(jsonData, k.senderTopic)
	_, _, err = k.producer.SendMessage(msg)
	if err != nil {
		return err
	}
	return nil
}

func getProducer(kafkaHost string) (sarama.SyncProducer, error) {

	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	return sarama.NewSyncProducer([]string{kafkaHost}, config)
}

func getMessage(msg []byte, topic string) *sarama.ProducerMessage {
	return &sarama.ProducerMessage{
		Topic:     topic,
		Partition: -1,
		Value:     sarama.ByteEncoder(msg),
	}
}
