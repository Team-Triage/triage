package main

import (
	"fmt"
	"sync"

	"github.com/team-triage/triage/dev/tmp"
	"github.com/team-triage/triage/dispatch"
	"github.com/team-triage/triage/fetcher"
	"github.com/team-triage/triage/filter"
)

// "github.com/confluentinc/confluent-kafka-go/kafka"

const TOPIC string = "triage-test-topic"

var wg sync.WaitGroup

func main() {
	fmt.Println("Triage firing up!!!")
	wg.Add(4)
	go fetcher.Consume(TOPIC)
	// go tmp.DummyDispatch()
	go dispatch.Dispatch()
	go filter.Filter()
	go tmp.Receiver()
	wg.Wait()
}
