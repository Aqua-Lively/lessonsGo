package main

import "fmt"

func main() {

	defer hanldePanic()

	// defer printMessage()

	// fmt.Println("main()")
	// fmt.Println("main() 2")
	// fmt.Println("main() 3")

	messages := []string{
		"message 1",
		"message 2",
		"message 3",
		"message 4",
		"message 5",
	}
	messages[5] = "messages 6"

	fmt.Println(messages)
	// panic("Help")
	// fmt.Println(messages)

}

func printMessage() {
	fmt.Println("printMessage()")
}

func hanldePanic() {
	if r := recover(); r != nil {
		fmt.Println(r)
	}

	fmt.Println("Program work")
}
