package main

import "fmt"

type Person struct {
	firstname string
	lastname  string
	age       int
}

type Contact struct {
	Email string
	Phone string
}

type Address struct {
	House int
	Area  string
	State string
}

type Employee struct {
	Person_Details Person
	Person_Contact Contact
	Person_Address Address
}

func main() {
	//1st method
	var person Person
	//fmt.Println("Print the person details  ", person)
	person.firstname = "Ak"
	person.lastname = "K"
	person.age = 26
	//fmt.Println("Print the person details  ", person)

	//2nd Method
	person1 := Person{
		firstname: "Bob",
		lastname:  "builder",
		age:       60,
	}
	fmt.Println("Person the details of person1\n", person1)

	//3rd Method

	var person2 = new(Person)
	person2.firstname = "Khh"
	person2.lastname = "kfdnn"
	person2.age = 40

	//fmt.Println("person2: ", person2)

	var employee1 Employee
	employee1.Person_Details = Person{
		firstname: "Aashish",
		lastname:  "Kushwaha",
		age:       26,
	}

	employee1.Person_Contact.Email = "ndn@gmail.com"
	employee1.Person_Contact.Phone = "987456321"

	employee1.Person_Address = Address{
		House: 17,
		Area:  "Varanasi",
		State: "UP",
	}
	fmt.Println("employee1", employee1)
}
