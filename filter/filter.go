package filter

import (
	"fmt"
	filterChan "triage/channels/toFilter"
	"triage/data/commitTable"
)

func Filter() {
	for {
		ack := filterChan.GetMessage()
		fmt.Printf("Filter received message at offset %v\n", ack.Offset)
		if ack.Status >= 1 {
			commitTable.CommitHash[ack.Offset] = true
		} else {
			fmt.Println("Sending to reaper!") // will replace with sending to reaper channel
			// send to reaper
			// nack
		}

	}
}
