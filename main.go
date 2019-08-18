package main

import (
	"github.com/RonnanSouza/kafka_environment/consumer"
	"github.com/RonnanSouza/kafka_environment/producer"
	"os"
)

func main()  {
	if os.Args[1] == "consumer" {
		consumer.StartConsumer()
	} else if os.Args[1] == "producer" {
		producer.StartProducer()
	}
}
