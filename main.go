package main

import (
	"fmt"
	"triage/dev/tmp"
	"triage/fetcher"
	// "github.com/confluentinc/confluent-kafka-go/kafka"
)

const TOPIC string = "triage-test-topic"

func main() {
	fmt.Println("Triage firing up!!!")
	go fetcher.Consume(TOPIC)
	// go filter.Start()
	go tmp.Receiver()
	fmt.Scanln()
}
