package commits

var c chan int = make(chan int)

func GetMessage() int {
	msg := <-c
	return msg
}

func AppendMessage(msg int) {
	c <- msg
}
