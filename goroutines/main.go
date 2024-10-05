package main

import (
	"fmt"
	"time"
)

func sup() {
	fmt.Println("Hewwww")
	time.Sleep(800 * time.Millisecond)
	fmt.Println("Animal Kingdom")

}

func saySup() {
	fmt.Println("Hi")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("END")
}

func main() {
	fmt.Println("learn Go routines")
	//defer fmt.Println("learn Goat routines")

	go sup()
	go saySup()
	time.Sleep(900 * time.Millisecond)
}
