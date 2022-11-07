package commitTable

import "github.com/team-triage/triage/types"

var CommitHash map[int]types.CommitStore = make(map[int]types.CommitStore) // map where keys are offsets and 'ack/nack' is boolean
