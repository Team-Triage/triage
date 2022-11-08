package reaper

import (
	"fmt"

	"github.com/team-triage/triage/channels/deadLetters"
	"github.com/team-triage/triage/data/commitTable"
)

func Reap() {
	for {
		ack := deadLetters.GetMessage()
		fmt.Printf("REAPER: Got a dead letter: %v \n", string(ack.Event.Value))
		// ^ is an abstraction for writing to DynamoDB
		// AFTER response from DynamoDB
		if entry, ok := commitTable.CommitHash.Read(ack.Offset); ok {
			entry.Value = true
			commitTable.CommitHash.Write(ack.Offset, entry)
		}
	}
}
