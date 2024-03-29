package main

import (
	"fmt"
	"log"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"

	"github.com/joho/godotenv"
	kafka2 "github.com/vitor7001/simulator/application/kafka"
	"github.com/vitor7001/simulator/infra/kafka"
)

func init() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("error loading .env file")
	}

}

//{"clientId": "1", "routeId": "1"}
func main() {

	msgChan := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)

	go consumer.Consume()

	for msg := range msgChan {
		fmt.Println(string(msg.Value))
		go kafka2.Produce(msg)
	}

}
