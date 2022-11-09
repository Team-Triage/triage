package commitCalculator

import (
	"log"
	"time"

	"github.com/team-triage/triage/channels/commits"
	"github.com/team-triage/triage/data/commitTable"
	// "github.com/go-co-op/gocron"
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
			commitTable.Delete(maxValidOffset) // okay because the method uses the underlying method on the SafeCommitHash
		} else {
			log.Fatalln("COMMIT CALCULATOR: Could not retrieve kafka message from commitHash")
		}

	}
}

func getMessageToCommit() int {
	offsets := commitTable.CommitHash.GetOffsets()
	maxValidOffset := -1
	for _, offset := range offsets {
		// iterate over range of sorted offsets, starting at the lowest
		if entry, ok := commitTable.CommitHash.Read(offset); ok {
			// if entry exists
			if entry.Value != true {
				// if entry has not been acknowledged, we're at the highest possible offset that we can commit
				break
			}
			// if entry has been acknowledged, keep going
			maxValidOffset = offset
		} else {
			// if entry does not exist, there are no more offsets in the commitHash
			break
		}
	}

	return maxValidOffset
}
