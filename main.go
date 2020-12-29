package main

import (
	"flag"
	"fmt"

	"github.com/RonnanSouza/kafka_environment/consumer"
	"github.com/RonnanSouza/kafka_environment/producer"
)

func main() {

	// Flags
	consumerFlagValue := flag.Bool("c", false, "Use this flag to start a Kafka Consumer")
	producerFlagValue := flag.Bool("p", false, "Use this flag to start a Kafka Producer")

	// Flag Processing
	flag.Parse()

	// Decision Time
	if *producerFlagValue == true {
		producer.StartProducer()
	} else if *consumerFlagValue == true {
		consumer.StartConsumer()
	} else {
		fmt.Print("Usage: \n -c     Use this flag to start a Kafka Consumer\n -p     Use this flag to start a Kafka Producer\n")
	}
}
