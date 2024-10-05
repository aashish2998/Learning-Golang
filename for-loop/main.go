package main

import "fmt"

func main() {

	for i := 0; i < 4; i++ {
		fmt.Println("Print the number", i)
	}

	count := 0
	for {
		fmt.Println("Infinite loop")
		count++
		if count == 1 {
			break
		}
	}
	number := []int{1, 2, 5, 6}
	for index, val := range number {
		fmt.Printf("index of the number %d, val of the number %d\n", index, val)
	}

	data := "hello,world"
	for index, char := range data {
		fmt.Printf("index of the data %d, char of the data %c\n", index, char)
	}
}
