package types

import "github.com/confluentinc/confluent-kafka-go/kafka"

type Acknowledgement struct {
	Status int
	Offset int
	Event  *kafka.Message
}
