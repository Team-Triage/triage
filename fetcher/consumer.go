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

func Fetch(topic string) {
	var wg sync.WaitGroup
	c := makeConsumer()
	go consume(c, topic)
	wg.Add(1)
	go committer(c)
	wg.Add(1)
	wg.Wait()
}

func makeConsumer() *kafka.Consumer {
	configFile := "dev/tmp/devConfig.properties" // os.Args[1] hard-coding the relative file path for now, since we're running from main.go
	conf := ReadConfig(configFile)
	conf["group.id"] = "kafka-go-getting-started" // need to make this an environmental variable so all instances of a given deployment share the same group.id
	conf["auto.offset.reset"] = "earliest"        // REQUIRES ADDITIONAL READING policy for when triage first connects to Kafka
	conf["enable.auto.commit"] = "false"          // turned off for manual committing (see consumer.Commit() or consumer.CommitMessage())

	c, err := kafka.NewConsumer(&conf)

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
			ev, err := c.ReadMessage(100 * time.Millisecond)
			if err != nil {
				// Errors are informational and automatically handled by the consumer
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
