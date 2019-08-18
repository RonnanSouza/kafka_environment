# Kafka Broker Environment
This repository has the code to run the example developed in this [article](google.com)

## Install
`go get github.com/RonnanSouza/kafka_environment`

## Set up kafka broker
`docker-compose up -d`

## Producer
`go run main.go producer`

## Consumer
`go run main.go consumer` 

**Note:** *To start multiples consumers, run this command in different terminals concurrently*.
