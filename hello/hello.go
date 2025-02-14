package main

import (
	"fmt"

	"github.com/daniel-joo-dsj/learn-golang/greetings"
)


func main() {
	var message string
	message = greetings.Hello("Daniel")
	fmt.Println(message)
}