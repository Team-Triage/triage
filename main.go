package main

import (
	"fmt"
	"sync"

	"github.com/team-triage/triage/commitCalculator"
	"github.com/team-triage/triage/consumerManager"
	"github.com/team-triage/triage/dev/tmp"
	"github.com/team-triage/triage/dispatch"
	"github.com/team-triage/triage/fetcher"
	"github.com/team-triage/triage/filter"
	"github.com/team-triage/triage/reaper"
)

// "github.com/confluentinc/confluent-kafka-go/kafka"

const TOPIC string = "triage-test-topic"

var wg sync.WaitGroup

func main() {
	fmt.Println("Triage firing up!!!")
	wg.Add(7)
	go fetcher.Consume(TOPIC)
	go dispatch.Dispatch()
	go filter.Filter()
	go reaper.Reap()
	go consumerManager.Start()
	go commitCalculator.Calculate()
	go tmp.Receiver()
	wg.Wait()
}
