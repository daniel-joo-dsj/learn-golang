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

	// Create an array of names
	names := []string{"Daniel", "Christian", "Mary"}

	// Request greeting messages for the names
	var messages map[string]string
	var err error
	messages, err = greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	// If no error, print messages
	fmt.Println(messages)
}