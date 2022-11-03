package channel

import (

)

var C chan string = make(chan string)

func Pinger(c chan string) {
		c <- "ping"
		// message := <- c
		// return c
} 

func Printer(c chan string) string {
	msg := <- c
	return msg
} 


