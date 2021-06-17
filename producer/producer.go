package producer

import (
	"fmt"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	k "kafka-golang-example/config"
)

func Producer() error {
	p, err := kafka.NewProducer(k.KafkaProducerConfig())
	if err != nil {
		panic(err)
		return err
	}
	defer p.Close()

	// Delivery report handler for produced messages
	go deliveryMessage(p)
	// Produce messages to topic (asynchronously)
	topic := "myTopic"
	for _, word := range []string{"Welcome to the kafka golang"} {
		p.Produce(k.KafkaProducerMessage(word,topic),nil)
	}

	// Wait for message deliveries before shutting down
	p.Flush(15 * 1000)
	return nil
}

func deliveryMessage(p *kafka.Producer)  {
	for e := range p.Events() {
		switch ev := e.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
			} else {
				fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
			}
		}
	}
}
