package main

import (
	"fmt"
	"triage/dev/tmp"
	"triage/fetcher"
	"triage/filter"
	// "github.com/confluentinc/confluent-kafka-go/kafka"
)

const TOPIC string = "triage-test-topic"

func main() {
	fmt.Println("Triage firing up!!!")
	go fetcher.Consume(TOPIC)
	go tmp.DummyDispatch()
	go filter.Filter()
	go tmp.Receiver()
	fmt.Scanln()
}
