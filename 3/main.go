package main

import "fmt"

func main() {
	message := "I work"

	fmt.Println(message)

	changeMessage(&message)

	fmt.Println(message)
}

func changeMessage(message *string) {
	*message += " (из функции print Message())"
}
