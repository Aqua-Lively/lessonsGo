package main

import "fmt"

func main() {
	fmt.Println(findMin(1, 2, 3, 4, 5, 6, 0))
}

func findMin(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}

	min := numbers[0]

	for _, i := range numbers {
		if numbers[i] < min {
			min = i
		}
	}

	return min
}
