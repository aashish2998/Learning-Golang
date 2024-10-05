package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name" `
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_adult"`
}

func main() {

	fmt.Println("Json")

	person := Person{Name: "Aashish", Age: 26, IsAdult: true}
	fmt.Println("Print the person details ", person)

	// convert person into  json encoding {marshalling}

	jsonData, err := json.Marshal(person)

	if err != nil {
		fmt.Println("the jsondata error")
	}
	fmt.Printf("type of json data %T\n", jsonData)
	fmt.Println("the jsondata deails", string(jsonData))

	//Decoding(UnMarshalling)

	var personData Person
	err = json.Unmarshal(jsonData, &personData)
	if err != nil {
		fmt.Println("print the unmarshal error", err)
		return
	}

	fmt.Println("print the personData", personData)
}
