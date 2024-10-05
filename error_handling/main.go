package main

import "fmt"

// func divide(a, b float64) (float64, error) {
// 	if b == 0 {
// 		return 0, fmt.Errorf("the denominator is 0")
// 	}
// 	return a / b, nil
// }

func divide(a, b float64) (float64, string) {
	if b == 0 {
		return 0, "the denominator is 0"
	}
	return a / b, "nil"
}
func main() {
	fmt.Println(" Learning error handling")

	// 	ans, err := divide(100, 0)
	// 	if err != nil {
	// 		fmt.Println("Error Handling")
	// 	}

	// 	fmt.Println("the ans is : ", ans)
	// }

	ans, _ := divide(100, 4)
	fmt.Println("the ans is : ", ans)
}
