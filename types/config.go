package types

import "github.com/confluentinc/confluent-kafka-go/kafka"

type TriageConfig struct {
	TopicName           string
	DeadLetterTableName string
	AuthenticationToken string
	KafkaConfigMap      kafka.ConfigMap
}
