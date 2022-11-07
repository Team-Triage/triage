package fetcher

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/team-triage/triage/channels/messages"
	"github.com/team-triage/triage/data/commitTable"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Consume(topic string) {
	// commented next 5 lines because we're running this from main.go, not the command line
	// if len(os.Args) != 2 {
	// 	fmt.Fprintf(os.Stderr, "Usage: %s <config-file-path>\n",
	// 		os.Args[0])
	// 	os.Exit(1)
	// }

	configFile := "dev/tmp/devConfig.properties" // os.Args[1] hard-coding the relative file path for now, since we're running from main.go
	conf := ReadConfig(configFile)
	conf["group.id"] = "kafka-go-getting-started" // need to make this an environmental variable so all instances of a given deployment share the same group.id
	conf["auto.offset.reset"] = "earliest"        // policy for when triage "dies"
	conf["enable.auto.commit"] = "false"          // turned off for manual committing (see consumer.Commit() or consumer.CommitMessage())

	c, err := kafka.NewConsumer(&conf)

	if err != nil {
		fmt.Printf("Failed to create consumer: %s", err)
		os.Exit(1)
	}

	// topic := "purchases" -- replaced with parameter
	err = c.SubscribeTopics([]string{topic}, nil)
	// Set up a channel for handling Ctrl-C, etc
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	// Process messages
	fmt.Println("Consumer running!")
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
			commitTable.CommitHash[int(ev.TopicPartition.Offset)] = false
		}
	}

	c.Close()
}
