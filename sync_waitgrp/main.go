package main

import (
	"fmt"
	"sync"
)

func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done() // signal that go-routines is done

	fmt.Printf("start the task %d\n", i)

	fmt.Printf("End the task %d\n", i)

}
func main() {

	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	// wait for all the workers to finish the task
	wg.Wait()
	fmt.Println("The task is completed")
}
