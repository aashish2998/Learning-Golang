package main

import (
	"fmt"
	"os"
)

func main() {

	// file, err := os.Create("example.txt")

	// if err != nil {
	// 	fmt.Println("unable to print the file", err)
	// 	return
	// }

	// defer file.Close()

	// content := " Yup yup king"

	// _, errors := io.WriteString(file, content+"\n") // byte,errors:=
	// if errors != nil {
	// 	fmt.Println("print the error", errors)
	// 	return
	// }

	// fmt.Println("successfully created a file")

	//method 1
	// file, err := os.Open("example.txt")
	// if err != nil {
	// 	fmt.Println("unable to print the file", err)
	// 	return
	// }

	// defer file.Close()

	// // create  a buffer to read file content
	// buffer := make([]byte, 1024)

	// // Read the file content into buffer
	// for {
	// 	n, err := file.Read(buffer)
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		fmt.Println("Error while reading the file", err)
	// 		return
	// 	}

	// 	// process the read content
	// 	fmt.Println(string(buffer[:n]))
	// }

	// shortcut function to read and print

	//Read the entire file into byte slice

	content, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Print the error", err)
		return
	}

	fmt.Println(string(content))

}
