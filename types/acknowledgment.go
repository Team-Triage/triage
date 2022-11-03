package types

import "github.com/confluentinc/confluent-kafka-go/kafka"

type Acknowledgment struct {
	Status int
	Offset int
	Event  *kafka.Message
}
