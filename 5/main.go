package main

import "fmt"

func main() {

	// matrix := make([][]int, 10)
	// counter := 0
	// for x := 0; x < 10; x++ {
	// 	matrix[x] = make([]int, 10)
	// 	for y := 0; y < 10; y++ {
	// 		counter++
	// 		matrix[x][y] = counter
	// 	}
	// 	fmt.Println(matrix[x])
	// }

	// x := 0

	// for true {
	// 	x++
	// 	fmt.Println(x)
	// }

	messages := []string{
		"message 1",
		"message 2",
		"message 3",
		"message 4",
		"message 5",
	}

	for _, message := range messages {
		fmt.Println(message)
	}
}
