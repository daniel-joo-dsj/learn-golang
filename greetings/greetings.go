package greetings

import (
	"fmt"
	"errors"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func randomFormat() string {
	rand.Seed(time.Now().UnixNano()) // need new seed to make rand.Intn return random number

	formats := []string{
		"Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}