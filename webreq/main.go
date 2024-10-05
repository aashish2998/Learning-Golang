package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	fmt.Println(" Learning web service")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error while getting GET response", err)
		return
	}

	defer res.Body.Close()
	fmt.Printf("print the data type of res : %T", res)
	// fmt.Println("print the response", res)

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("print the error", err)
		return
	}

	fmt.Println("print the response", string(data))
}
