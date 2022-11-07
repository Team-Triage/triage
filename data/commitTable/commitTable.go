package commitTable

import "github.com/team-triage/triage/types"

var CommitHash map[int]types.CommitStore = make(map[int]types.CommitStore) // map where keys are offsets and 'ack/nack' is boolean

func Delete(offset int) {
	currentOffset := offset
	for {
		if _, ok := CommitHash[offset]; ok {
			delete(CommitHash, offset)
			currentOffset--
		} else {
			break
		}
	}
}
