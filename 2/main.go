package main

import "fmt"

var msg string

func init(){
	msg = "from init()"
}

func main() {

	fmt.Println(msg)

	inc := increment()
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
}

func increment() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func increment2() int {
	count := 0
	count++
	return count
}
