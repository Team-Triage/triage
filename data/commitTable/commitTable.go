package commitTable

var CommitHash map[int]bool = make(map[int]bool) // map where keys are offsets and 'ack/nack' is boolean
