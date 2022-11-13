package reaper

import (
	"fmt"

	"github.com/team-triage/triage/channels/deadLetters"
	"github.com/team-triage/triage/data/commitTable"
)

func Reap() {
	sess := makeDynamoSession()
	svc := makeDynamoClient(sess)
	// create dynamo session
	// create dynamo client
	// pass to writer function
	for {
		ack := deadLetters.GetMessage()

		fmt.Printf("REAPER: Got a dead letter: %v \n", string(ack.Event.Value))

		err := writeDeadLetter(ack.Event, svc)

		// below clause to be removed pending deployed DynamoDB instance
		if err != nil {
			fmt.Println("REAPER: DynamoDB not available!")
		}

		// for err != nil {
		// 	err = writeDeadLetter(ack.Event, svc)
		// } // if we get an error, keep trying to send to dynamo

		if entry, ok := commitTable.CommitHash.Read(ack.Offset); ok {
			entry.Value = true
			commitTable.CommitHash.Write(ack.Offset, entry)
		}
	}
}
