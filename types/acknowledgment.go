package types

import "github.com/confluentinc/confluent-kafka-go/kafka"

type Acknowledgment struct {
	event  *kafka.Message
	status int
}
