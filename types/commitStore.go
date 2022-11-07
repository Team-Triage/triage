package types

import "github.com/confluentinc/confluent-kafka-go/kafka"

type CommitStore struct {
	Value   bool
	Message *kafka.Message
}
