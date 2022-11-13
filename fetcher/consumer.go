package fetcher

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/team-triage/triage/channels/commits"
	"github.com/team-triage/triage/channels/messages"
	"github.com/team-triage/triage/data/commitTable"
	"github.com/team-triage/triage/types"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Fetch(topic string, kafkaConf kafka.ConfigMap) {
	var wg sync.WaitGroup
	c := makeConsumer(kafkaConf)
	go consume(c, topic)
	wg.Add(1)
	go committer(c)
	wg.Add(1)
	wg.Wait()
}

func makeConsumer(kafkaConf kafka.ConfigMap) *kafka.Consumer {
	kafkaConf["group.id"] = "team-triage"       // need to make this an environmental variable so all instances of a given deployment share the same group.id
	kafkaConf["auto.offset.reset"] = "earliest" // REQUIRES ADDITIONAL READING policy for when triage first connects to Kafka
	kafkaConf["enable.auto.commit"] = "false"   // turned off for manual committing (see consumer.Commit() or consumer.CommitMessage())

	c, err := kafka.NewConsumer(&kafkaConf)

	if err != nil {
		fmt.Printf("Failed to create consumer: %s", err)
		os.Exit(1)
	}
	return c
}

func consume(c *kafka.Consumer, topic string) {
	err := c.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		fmt.Printf("Failed to subscribe: %s\n", err)
		os.Exit(1)
	}
	// Set up a channel for handling Ctrl-C, etc
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	// Process messages
	fmt.Println("FETCHER: Consumer running!")
	run := true
	for run {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev, err := c.ReadMessage(100 * time.Second)
			if err != nil {
				fmt.Printf("FETCHER: Error from ReadMessage %v!\n", err)
				continue
			}
			messages.AppendMessage(ev) // writing event to channel
			commitStore := types.CommitStore{Value: false, Message: ev}
			commitTable.CommitHash.Write(int(ev.TopicPartition.Offset), commitStore)
		}
	}
	c.Close()
}

func committer(c *kafka.Consumer) {
	for {
		msg := commits.GetMessage()
		fmt.Printf("FETCHER: Going to commit offset: %v\n", msg.TopicPartition.Offset)
		c.CommitMessage(msg)
		fmt.Println("FETCHER: Committed successfully!")
	}
}
