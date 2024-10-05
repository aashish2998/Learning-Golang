package main

import "fmt"

func modifyByReference(num *int) {
	*num = *num + 10

}

func main() {

	//var num int = 2

	num := 2

	//var ptr *int = &num
	ptr := &num

	fmt.Println("the value of the number ", num)
	fmt.Println("the address of the ptr ", ptr)
	fmt.Println("Pointer contains the value  ", *ptr)

	var pointer *int
	if pointer == nil {
		fmt.Println("The pointer is an error")
	}

	value := 40
	modifyByReference(&value)
	fmt.Println("Print the value", value)
}
