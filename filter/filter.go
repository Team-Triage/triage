package filter

import (
	"fmt"
	ackChannel "triage/channels/acknowledgements"
	"triage/data/commitTable"
)

func Filter() {
	for {
		ack := ackChannel.GetMessage()
		fmt.Printf("Filter received message at offset %v\n", ack.Offset)
		if ack.Status == 1 { // if ack, simply updated commitHash
			commitTable.CommitHash[ack.Offset] = true
		} else {
			fmt.Println("Sending to reaper!") // will replace with sending to reaper channel
			// send to reaper
			// nack
		}

	}
}
