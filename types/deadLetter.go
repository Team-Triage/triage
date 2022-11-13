package types

type DeadLetter struct {
	UUID      string
	TIMESTAMP string
	Topic     string
	Partition int
	Offset    int
	Key       string
	String    string
	Headers   string
}
