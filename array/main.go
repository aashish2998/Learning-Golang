package main

import "fmt"

func main() {

	fmt.Println("Learning Go lang")

	var name [5]string

	name[0] = " AJsedsg "
	name[1] = "Akwe"

	fmt.Println("print the array", name)

	var numbers = [5]int{1, 2, 3, 4, 5}

	fmt.Println(" The numbers array  is ", numbers)
	//fmt.Printf("%T", numbers)

	fmt.Println(" The numbers array's length  is ", len(numbers))

	fmt.Println("print the array", name[1])

	var price [5]string
	fmt.Println("the price is", price)
	fmt.Printf("the price is %q", price)

}
