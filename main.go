package main

import (
	"fmt"
	"sync"
	"triage/dev/tmp"
	"triage/dispatch"
	"triage/fetcher"
	"triage/filter"
)

// "github.com/confluentinc/confluent-kafka-go/kafka"

const TOPIC string = "triage-test-topic"

var wg sync.WaitGroup

func main() {

	fmt.Println("Triage firing up!!!")
	wg.Add(1)
	go fetcher.Consume(TOPIC)
	// go tmp.DummyDispatch()
	go dispatch.Dispatch()
	go filter.Filter()
	go tmp.Receiver()
	wg.Wait()
}
