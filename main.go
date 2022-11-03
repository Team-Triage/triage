package main

import (
	"fmt"
	"triage/dev/tmp"
	"triage/fetcher"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

const TOPIC string = "triage-test-topic"

func main() {
	fmt.Println("Triage firing up!!!")

	toDispatch := make(chan *kafka.Message)

	go fetcher.Consume(toDispatch, TOPIC)
	go tmp.Receiver(toDispatch)
	fmt.Scanln()
}
