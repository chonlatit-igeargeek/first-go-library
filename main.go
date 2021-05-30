package greetings

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) {
	// Return a greeting that embeds the name in a message.
	fmt.Println(fmt.Sprintf("Hi, %v. Welcome!", name))
}
