package main

import (
	"fmt"
)

func main() {

	//numbers := []int{1, 2, 5, 6}
	// numbers = append(numbers, 5, 8, 9, 4, 50, 1, 5, 6)
	// fmt.Println("Print the numbers", numbers)
	// sort.Ints(numbers)
	// fmt.Println("Print the numbers", numbers)
	// fmt.Printf("Print the data type of numbers %T\n", numbers)
	// fmt.Println("Print the data type of numbers ", len(numbers))

	numbers := make([]int, 3, 5)
	numbers = append(numbers, 1, 5, 6)

	fmt.Println(" Slice", numbers)
	fmt.Println(" len of the slice", len(numbers))
	fmt.Println("cap of the slice", cap(numbers))
}
