package main

import "fmt"

func main() {

	studentGrades := make(map[string]int)

	studentGrades["Aashish"] = 100
	studentGrades["Aj"] = 98
	studentGrades["Ak"] = 93
	studentGrades["Alpha"] = 91
	studentGrades["Beta"] = 88

	fmt.Println("Print the marks of Aj:", studentGrades["Aj"])
	studentGrades["Aj"] = 100

	fmt.Println("Print the marks of Aj:", studentGrades["Aj"])

	delete(studentGrades, "Aj")
	fmt.Println("Print the marks of Aj:", studentGrades["Aj"])

	//checking if the key exists

	grades, exist := studentGrades["David"]
	fmt.Println("Print the grades of david :", grades)
	fmt.Println("Print does David exists:", exist)

	for index, value := range studentGrades {
		fmt.Printf("key is %s, the value is %d\n", index, value)
	}

	person := map[string]int{
		"Alice": 90,
		"Bob":   85,
		"Alpha": 80,
	}

	for index, value := range person {
		fmt.Printf("----------Key is %s, the value is %d\n", index, value)
	}
}
