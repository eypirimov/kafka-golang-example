package main

import (
	"kafka-golang-example/consumer"
	"testing"
)

func TestConsumer(t *testing.T) {
	err := consumer.Consumer()

	if err != nil {
		t.Errorf("Consumer Error %v",err)
	}
}
