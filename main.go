package main

import (
	"fmt"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/kafka"
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
	kafkaConf := kafka.ConfigMap{}
	fmt.Println("Triage firing up!!!")

	path := "config.properties"
	config := utils.ReadConfig(path)
	kafkaConf["bootstrap.servers"] = config["bootstrap.servers"]
	kafkaConf["security.protocol"] = config["security.protocol"]
	kafkaConf["sasl.mechanisms"] = config["sasl.mechanisms"]
	kafkaConf["sasl.username"] = config["sasl.username"]
	kafkaConf["sasl.password"] = config["sasl.password"]
	kafkaConf["session.timeout.ms"] = config["session.timeout.ms"]
	fmt.Println(kafkaConf)
	topic := config["kafka.topic"]

	wg.Add(7)
	go fetcher.Fetch(topic, kafkaConf)
	go dispatch.Dispatch()
	go filter.Filter()
	go reaper.Reap()
	go consumerManager.Start()
	go commitCalculator.Calculate()
	// go tmp.Receiver()
	wg.Wait()
}
