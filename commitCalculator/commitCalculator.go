package commitCalculator

import (
	"log"
	"time"

	"github.com/team-triage/triage/channels/commits"
	"github.com/team-triage/triage/data/commitTable"
)

func Calculate() {
	for {
		time.Sleep(time.Second * 5)
		maxValidOffset := getMessageToCommit()
		if maxValidOffset == -1 {
			continue
		}
		if kafkaMessage, ok := commitTable.CommitHash.Read(maxValidOffset); ok {
			commits.AppendMessage(kafkaMessage.Message)
			commitTable.Delete(maxValidOffset)
		} else {
			log.Fatalln("COMMIT CALCULATOR: Could not retrieve kafka message from commitHash")
		}

	}
}

func getMessageToCommit() int {
	offsets := commitTable.CommitHash.GetOffsets()
	maxValidOffset := -1
	for _, offset := range offsets {
		if entry, ok := commitTable.CommitHash.Read(offset); ok {
			if entry.Value != true {
				break
			}
			maxValidOffset = offset
		} else {
			break
		}
	}

	return maxValidOffset
}
