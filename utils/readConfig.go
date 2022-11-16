package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/team-triage/triage/types"
)

func readConfig(configFilePath string) (config types.TriageConfig, kafkaConfigs map[string]string) {

	config = types.TriageConfig{}
	kafkaConfigs = make(map[string]string)

	file, err := os.Open(configFilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open file: %s", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if !strings.HasPrefix(line, "#") && len(line) != 0 {
			kv := strings.Split(line, "=")
			parameter := strings.TrimSpace(kv[0])
			value := strings.TrimSpace(kv[1])
			if parameter == "topic.name" {
				config.TopicName = value
				continue
			} else if parameter == "authentication.token" {
				config.AuthenticationToken = value
				continue
			} else if parameter == "num.of.partitions" {
				continue
			} else {
				kafkaConfigs[parameter] = value
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Failed to read file: %s", err)
		os.Exit(1)
	}

	return config, kafkaConfigs
}
