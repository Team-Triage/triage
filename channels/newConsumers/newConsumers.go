package newConsumers

var c chan string = make(chan string)

func GetMessage() string {
	msg := <-c
	return msg
}

func AppendMessage(msg string) {
	c <- msg
}
