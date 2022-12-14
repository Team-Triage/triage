package reaper

import (
	"fmt"

	"github.com/team-triage/triage/channels/deadLetters"
	"github.com/team-triage/triage/data/commitTable"
)

func Reap(deadLetterTableName string) {
	sess := makeDynamoSession()
	dynamoClient := makeDynamoClient(sess)

	for {
		ack := deadLetters.GetMessage()

		fmt.Printf("REAPER: Got a dead letter: %v \n", string(ack.Event.Value))

		err := writeDeadLetter(ack.Event, dynamoClient, deadLetterTableName)

		if err != nil {
			fmt.Println("REAPER: DynamoDB not available!")
		}

		if entry, ok := commitTable.CommitHash.Read(ack.Offset); ok {
			entry.Value = true
			commitTable.CommitHash.Write(ack.Offset, entry)
		}
	}
}
