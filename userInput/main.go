package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("What's up Buddy")

	// name := ""
	// fmt.Scan(&name)

	// fmt.Println("Hey i am cool", name)  this i one way to input the value but it doesn't work after the first blank space so that's why were are using bufio to be able to input and print the complete variable value

	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Println("I am cool", name)

}
