package toFilter

import (
	"triage/types"
)

var c chan *types.Acknowledgment = make(chan *types.Acknowledgment)

func GetMessage() *types.Acknowledgment {
	msg := <-c
	return msg
}

func AppendMessage(msg *types.Acknowledgment) {
	c <- msg
}
