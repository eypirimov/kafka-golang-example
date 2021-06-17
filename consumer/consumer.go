package consumer

import (
	"fmt"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"kafka-golang-example/config"
)

func Consumer() error{
	c, err := kafka.NewConsumer(config.KafkaConsumerConfig())

	if err != nil {
		panic(err)
	}

	topic:="myTopic"
	c.SubscribeTopics([]string{topic}, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
			return nil // this return is only for test 
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			return err
		}
	}

	c.Close()
	return nil
}
