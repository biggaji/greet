package greet

import "fmt"

// Insert user's name into message and return a string
func SayHi(name string) string {
	message := fmt.Sprintf("Hey %v. Welcome to Go land!", name)
	return message
}
