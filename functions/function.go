package main

import "fmt"

func add(a, b, c int) int {
	return a + b + c
}

func main() {
	fmt.Println(add(18, 11, 26))

	a, b := swap("hello", "world")
	fmt.Println(a, b)
	/*if a == "world" {
		fmt.Println("The swap function worked!")
	}
	*/
}
