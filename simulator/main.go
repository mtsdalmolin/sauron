package main

import (
	"fmt"
	"log"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
	kafkaApp "github.com/mtsdalmolin/sauron-simulator/application/kafka"
	"github.com/mtsdalmolin/sauron-simulator/infra/kafka"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	msgChannel := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChannel)

	// start new thread
	go consumer.Consume()

	for message := range msgChannel {
		fmt.Println("Produced on topic (" + string(*message.TopicPartition.Topic) + "): " + string(message.Value))
		go kafkaApp.Produce(message)
	}
}
