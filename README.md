# Kafka Broker Environment
This repository has the code to run the example developed in this [article](https://medium.com/@ronnansouza/setting-up-a-kafka-broker-using-docker-creating-a-producer-and-consumer-group-with-multiple-384b724cd324?sk=4f828cdc1adeec088e9e67f35dbb0c8f)

## Install
`go get github.com/RonnanSouza/kafka_environment`

## Set up kafka broker
`docker-compose up -d`

## Producer
`go run main.go producer`

## Consumer
`go run main.go consumer` 

**Note:** *To start multiples consumers, run this command in different terminals concurrently*.
