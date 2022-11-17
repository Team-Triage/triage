package main

import (
	"fmt"
	"sync"

	"github.com/team-triage/triage/commitCalculator"
	"github.com/team-triage/triage/consumerManager"
	"github.com/team-triage/triage/dispatch"
	"github.com/team-triage/triage/fetcher"
	"github.com/team-triage/triage/filter"
	"github.com/team-triage/triage/reaper"
	"github.com/team-triage/triage/utils"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Triage starting!")

	config := utils.GetConfig()
	wg.Add(1)
	go fetcher.Fetch(config.TopicName, config.KafkaConfigMap)

	wg.Add(1)
	go dispatch.Dispatch()

	wg.Add(1)
	go filter.Filter()

	wg.Add(1)
	go reaper.Reap(config.DeadLetterTableName)

	consumerManager.SetToken(config.AuthenticationToken)
	wg.Add(1)
	go consumerManager.StartHttpServer()

	wg.Add(1)
	go commitCalculator.Calculate()
	wg.Wait()
}
