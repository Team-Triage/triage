package toDispatch

import "github.com/confluentinc/confluent-kafka-go/kafka"

var c chan *kafka.Message = make(chan *kafka.Message)

func GetMessage() *kafka.Message {
	msg := <-c
	return msg
}

func AppendMessage(msg *kafka.Message) {
	c <- msg
}
