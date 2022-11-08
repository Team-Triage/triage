package filter

import (
	"fmt"

	"github.com/team-triage/triage/channels/acknowledgements"
	"github.com/team-triage/triage/channels/deadLetters"
	"github.com/team-triage/triage/data/commitTable"
)

func Filter() {
	for {
		ack := acknowledgements.GetMessage()
		fmt.Printf("FILTER: Received message at offset %v with status %v\n", ack.Offset, ack.Status)
		if ack.Status == 1 { // if ack, simply updated commitHash
			if entry, ok := commitTable.CommitHash.Read(ack.Offset); ok {
				entry.Value = true
				commitTable.CommitHash.Write(ack.Offset, entry)
			}
		} else {
			deadLetters.AppendMessage(ack)
		}

	}
}
