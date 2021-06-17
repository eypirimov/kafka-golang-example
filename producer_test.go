package main

import (
	"kafka-golang-example/producer"
	"testing"
)

func TestProducer(t *testing.T) {
	err:=producer.Producer()
	if err != nil {
		t.Errorf("Producer Error %v",err)
	}
}
