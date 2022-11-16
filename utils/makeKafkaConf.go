package utils

import "github.com/confluentinc/confluent-kafka-go/kafka"

func makeKafkaConf(conf map[string]string) (kafkaConf kafka.ConfigMap) {
	kafkaConf = make(kafka.ConfigMap)
	keys := []string{}
	for k := range conf {
		keys = append(keys, k)
	}

	for i := range keys {
		kafkaConf[keys[i]] = conf[keys[i]]
	}

	return kafkaConf
}
