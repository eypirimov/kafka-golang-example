package config

import (
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func KafkaProducerConfig() *kafka.ConfigMap {
	return &kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092"}
}

func KafkaConsumerConfig() *kafka.ConfigMap {
	return &kafka.ConfigMap{"bootstrap.servers": "localhost:9092",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest"}
}

func KafkaProducerMessage(msg string, topic string) *kafka.Message {
	return &kafka.Message{TopicPartition: kafka.TopicPartition{Topic: &topic,
		Partition: kafka.PartitionAny}, Value: []byte(msg)}
}
