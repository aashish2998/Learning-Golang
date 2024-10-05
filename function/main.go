package main

import "fmt"

func add(a, b int) (result int) {
	result = a + b

	return
}
func simpleFunction() {
	fmt.Println("This is the sub function")
}
func main() {

	fmt.Println("Learing function in Go Lang")
	simpleFunction()

	ans := add(3, 8)

	fmt.Println("The sum is", ans)
}
