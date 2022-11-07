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
		fmt.Printf("Filter received message at offset %v\n", ack.Offset)
		if ack.Status == 1 { // if ack, simply updated commitHash
			commitTable.CommitHash[ack.Offset] = true
		} else {
			deadLetters.AppendMessage(ack)
		}

	}
}
