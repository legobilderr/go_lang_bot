package greetings

import (
	"math/rand"
	"time"
)

func RandomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %s. Welcome!",
		"Great to see you, %s!",
		"Hail, %s! Well met!",
	}
	rand.Seed(time.Now().UnixNano())

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}

func test() {

}
