package greet

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Insert user's name into message and return a string
func SayHi(name string) (string, error) {
	if name == "" {
		return name, errors.New("name must be provided")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	predefinedMessages := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return predefinedMessages[rand.Intn(len(predefinedMessages))]
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string) //make a map of messages --> map[name] = message

	for _, name := range names {
		message, err := SayHi(name)

		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}
