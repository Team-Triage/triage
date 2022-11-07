package reaper

import (
	"fmt"

	"github.com/team-triage/triage/channels/deadLetters"
)

func Reap() {
	for {
		ack := deadLetters.GetMessage()
		fmt.Printf("Got a dead letter: %v \n", string(ack.Event.Value))
		// ^ is an abstraction for writing to DynamoDB
	}
}
