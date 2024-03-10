package main

import "fmt"

func main() {
	// users := map[string]int{}
	// users := map[string]int{
	// 	"Vasya": 15,
	// 	"Petya": 23,
	// 	"Vova":  48,
	// }
	// age, exists := users["Petya"]
	// if exists {
	// 	fmt.Println("Petya", age)
	// } else {
	// 	fmt.Println("Net v spiske")
	// }

	// fmt.Println(users["Vasya"], age)
	var users map[string]int
	users["Vasya"] = 15

	delete(users, "Vasya")

	for key, value := range users {
		fmt.Println(key, value)
	}
}
