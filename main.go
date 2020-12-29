package main

import (
	"flag"
	"fmt"

	"github.com/RonnanSouza/kafka_environment/consumer"
	"github.com/RonnanSouza/kafka_environment/producer"
)

func main() {

	// Flags
	consumerFlagValue := flag.Bool("c", false, "    Use this flag to start a Kafka Consumer")
	producerFlagValue := flag.Bool("p", false, "    Use this flag to start a Kafka Producer")
	stringFlagValue := flag.String("a", "", "    Use this flag with either \"consumer\" or \"producer\"")

	// Flag Processing
	flag.Parse()

	// Decision Time
	if *producerFlagValue == true {
		producer.StartProducer()
	} else if *consumerFlagValue == true {
		consumer.StartConsumer()
	} else if *stringFlagValue == "consumer" {
		consumer.StartConsumer()
	} else if *stringFlagValue == "producer" {
		producer.StartProducer()
	} else {
		fmt.Print("Usage: \n -c     Use this flag to start a Kafka Consumer\n -p     Use this flag to start a Kafka Producer\n -a     Use this flag with either \"consumer\" or \"producer\"\n")
	}
}
