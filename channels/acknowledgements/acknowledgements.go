package acknowledgements

import (
	"triage/types"
)

var c chan *types.Acknowledgement = make(chan *types.Acknowledgement)

func GetMessage() *types.Acknowledgement {
	msg := <-c
	return msg
}

func AppendMessage(msg *types.Acknowledgement) {
	c <- msg
}
