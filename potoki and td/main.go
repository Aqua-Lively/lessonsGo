package main

import "fmt"

func main() {
	s := []int{1, -5, 24, 64, 99}

	c := make(chan int)

	go sumMass(s[:len(s)/2], c)
	go sumMass(s[len(s)/2:], c)

	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}

func sumMass(s []int, c chan int) {
	sum := 0
	for _, elem := range s {
		sum += elem
	}

	c <- sum

}
