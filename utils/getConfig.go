package utils

import "github.com/team-triage/triage/types"

func GetConfig() types.TriageConfig {
	configFilePath := "config.properties"
	config, kafkaConfigs := readConfig(configFilePath)
	config.KafkaConfigMap = makeKafkaConf(kafkaConfigs)
	return config
}
