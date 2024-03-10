package main

import (
	"errors"
	"fmt"
)

func main() {
	// messages := [3]string{"1", "2", "3"}
	// messages := [3]string{}
	// messages[0] = "5"

	// messages := []string{"1", "2", "3"}

	// printMessages(messages)

	// fmt.Println(messages)

	// var messages []string

	messages := make([]string, 5)

	messages = append(messages, "6")
	fmt.Println(messages)
	fmt.Println(len(messages))
	fmt.Println(cap(messages))
}

func printMessages(messages []string) error {
	if len(messages) == 0 {
		return errors.New("empty array")
	}
	messages[1] = "5"
	fmt.Println(messages)

	return nil
}
