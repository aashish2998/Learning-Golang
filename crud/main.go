package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Todo struct {
	UserID    int    `json:"userid"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func performGet() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("there is an error", err)
		return
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {

		fmt.Println("The code is showing the error", res.Status)
		return

	}

	// data, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println("Show the error", err)
	// 	return
	// }
	// fmt.Println("Print the data value", string(data))

	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("print the error", err)
		return
	}

	fmt.Println("print the todo ", todo)

}

func performPost() {

	// todo := Todo{
	// 	UserID:    23,
	// 	Title:     "Mango Man",
	// 	Completed: true,
	// }

	// // convert struct to json
	// jsonData, err := json.Marshal(todo)
	// if err != nil {
	// 	fmt.Println("print the error", err)
	// 	return
	// }

	// // convert json to string

	// jsonString := string(jsonData)

	// // True, they both represent the same data but in different forms:
	// // jsonString: A simple string in Go, just a plain text value.
	// // jsonReader: An io.Reader interface that allows reading data as a stream. It's more flexible, especially for handling larger datasets or reading data piece-by-piece rather than loading it all into memory at once.
	// // The key difference lies in usage:
	// // Use jsonString when you need direct access to the data.
	// // Use jsonReader when you need to read the data in a more controlled and memory-efficient manner.

	// //convert string data to reader

	// jsonReader := strings.NewReader(jsonString)

	// myUrl := "https://jsonplaceholder.typicode.com/todos/1"

	// //send post request

	// res, err := http.Post(myUrl, "application/json", jsonReader)
	// if err != nil {
	// 	fmt.Println("error response request", err)
	// 	return
	// }

	// defer res.Body.Close()

	// data, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println("print the data error")
	// 	return
	// }

	// fmt.Println("Response", string(data))

	// //fmt.Println("respo", res.Status)

	todo := Todo{
		UserID:    23,
		Title:     "Mango Man",
		Completed: true,
	}

	// Convert struct to JSON
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	//convert to json data to string
	jsonString := string(jsonData)

	// Convert JSON data to reader
	jsonReader := strings.NewReader(jsonString)

	myUrl := "https://jsonplaceholder.typicode.com/todos"

	// Send POST request
	res, err := http.Post(myUrl, "application/json", jsonReader)
	if err != nil {
		fmt.Println("Error making POST request:", err)
		return
	}
	defer res.Body.Close()

	// Read response body
	// data, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println("Error reading response body:", err)
	// 	return
	// }

	data, _ := ioutil.ReadAll(res.Body)

	fmt.Println("Response:", string(data))
}

func performUpdateRequest() {
	todo := Todo{
		UserID:    2311,
		Title:     "Bad Boys",
		Completed: false,
	}

	// Convert struct to JSON
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	// convert json data to string

	jsonString := string(jsonData)
	// Convert JSON data to reader
	jsonReader := strings.NewReader(jsonString)

	const myurl = "https://jsonplaceholder.typicode.com/todos/1"

	//create PUT request
	req, err := http.NewRequest(http.MethodPut, myurl, jsonReader)
	if err != nil {
		fmt.Println("Error creating put request", err)
		return
	}

	req.Header.Set("Content-type", "application/json")

	//send the request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("error while sending the client request")
		return
	}

	defer res.Body.Close()

	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println(" print the res", string(data))
	fmt.Println("Response status ", res.Status)
}

func perforDeleteRequest() {

	const myurl = "https://jsonplaceholder.typicode.com/todos/1"
	//create PUT request
	req, err := http.NewRequest(http.MethodDelete, myurl, nil)
	if err != nil {
		fmt.Println("Error creating put request", err)
		return
	}

	//send the request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("error while sending the client request")
		return
	}

	defer res.Body.Close()
	fmt.Println("Response status ", res.Status)

}

func main() {
	fmt.Println("learning crud")
	//performGet()
	//performPost()
	//performUpdateRequest()
	perforDeleteRequest()

}
