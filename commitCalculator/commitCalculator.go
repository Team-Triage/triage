package commitCalculator

import (
	"sort"
	"time"

	"github.com/team-triage/triage/channels/commits"
	"github.com/team-triage/triage/data/commitTable"

	// "github.com/go-co-op/gocron"
	"golang.org/x/exp/maps"
)

/*
create 'cron' job that every X seconds
	get max valid offset
	if max valid offset is -1, go to next loop (no valid offsets)

	send the message from the commitHash[offset].Event to the "commits" channel
		we should have a go routine that listens on that channel, and calls consumer.CommitMessage() on said message

	Iterate through the keys of the commitHash and delete all values that are equal to or lower than the committed offset
*/

func Calculate() {
	for {
		time.Sleep(time.Second * 5)
		maxValidOffset := getMessageToCommit()
		if maxValidOffset == -1 {
			continue
		}

		commits.AppendMessage(maxValidOffset)
		commitTable.Delete(maxValidOffset)
	}
}

func getMessageToCommit() int {
	offsets := maps.Keys(commitTable.CommitHash)
	sort.Ints(offsets)

	maxValidOffset := -1
	for _, offset := range offsets {
		if commitTable.CommitHash[offset].Value == true {
			maxValidOffset = offset
		} else {
			break
		}
		/* look at current offset
		if value in commitHash is true
			reassign maxValidOffset to offset
		else break
		*/
	}

	return maxValidOffset
}
