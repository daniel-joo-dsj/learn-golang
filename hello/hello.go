package main

import (
	"fmt"
	"log"

	"github.com/daniel-joo-dsj/learn-golang/greetings"
)


func main() {
	// Set log properties
	log.SetPrefix("greetings: ")
    log.SetFlags(0)

	// Request a greeting message
	var message string
	var err error
	message, err = greetings.Hello("Daniel")

	if err != nil {
		log.Fatal(err)
	}

	// If no error, print message
	fmt.Println(message)
}