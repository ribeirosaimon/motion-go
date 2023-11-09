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
	kafkaHost   string
}

var kafkaProducer *MotionKafka

func NewMotionKafka(configurations properties.Properties) {
	host := configurations.GetString("kafka.host", "")
	port := configurations.GetString("kafka.port", "")
	topic := configurations.GetString("kafka.topic", "")

	kafkaHost := fmt.Sprintf("%s:%s", host, port)

	producer, err := getProducer(kafkaHost)
	if err != nil {
		kafkaDefault := "localhost:9092"
		syncProducer, _ := getProducer(kafkaDefault)
		kafkaProducer = &MotionKafka{producer: syncProducer, kafkaHost: kafkaDefault, senderTopic: "default"}
	}
	kafkaProducer = &MotionKafka{producer: producer, kafkaHost: kafkaHost, senderTopic: topic}
}

func (k *MotionKafka) openConn() error {
	syncProducer, err := getProducer(k.kafkaHost)
	k.producer = syncProducer
	if err != nil {
		return err
	}
	return nil
}

func GetMotionKafka() *MotionKafka {
	return kafkaProducer
}

func (k *MotionKafka) SendMessage(message interface{}) error {
	if err := k.openConn(); err != nil {
		return err
	}
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
