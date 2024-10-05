package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {

	defer fmt.Println("coding is fun")
	data := add(9, 6)
	defer fmt.Println("Print the data value :", data) // if there are 2 differ keyword uses then it will follow LIFO Condition
	fmt.Println("code the sol")

}
