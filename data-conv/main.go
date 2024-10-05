package main

import (
	"fmt"
	"strconv"
)

func main() {

	var num int = 2

	fmt.Println("Print the value of ", num)
	fmt.Printf("print the type of %T\n", num)

	var data float64 = float64(num)

	fmt.Println("Print the value of ", data)
	fmt.Printf("print the type of %T\n", data)

	str := 545
	ans := strconv.Itoa(str)

	fmt.Println("Print the value of ", ans)
	fmt.Printf("print the type of %T\n", ans)

	strng_num := "1236"

	sol, _ := strconv.Atoi(strng_num)

	fmt.Println("Print the value of ", sol)
	fmt.Printf("print the type of %T\n", sol)

	num_string := "123.6"

	solution, _ := strconv.ParseFloat(num_string, 64) // the string is not getting convert to the exacct value so we used parsefloat to return the int  value as in the num_string variable the data type is float inside the string , so float can't be converted to string it will give zero

	fmt.Println("Print the value of ", solution)
	fmt.Printf("print the type of %T\n", solution)
}
