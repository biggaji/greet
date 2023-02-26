package greet

import (
	"errors"
	"fmt"
)

// Insert user's name into message and return a string
func SayHi(name string) (string, error) {
	if name == "" {
		return "", errors.New("name must be provided")
	}

	message := fmt.Sprintf("Hey %v. Welcome to Go land!", name)
	return message, nil
}
