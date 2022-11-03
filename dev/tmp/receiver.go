package tmp

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Receiver(toDispatch chan *kafka.Message) {
	for msg := range toDispatch {
		fmt.Println(msg)
	}
}
