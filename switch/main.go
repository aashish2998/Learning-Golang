package main

import "fmt"

func main() {

	day := 22

	switch day {

	case 1, 30, 22:
		fmt.Println("Its Monday ")

	case 2:
		fmt.Println("Its Wednesday")

	case 3:
		fmt.Println("Fun day")

	default:
		fmt.Println("Its a weekend")
	}

	month := "December"

	switch month {

	case "January , February ,March":
		fmt.Println("Winter is here")

	case " April, May, June":
		fmt.Println("Winter is gone")

	case " July,August,September":
		fmt.Println("Winter is yet to come")

	default:
		fmt.Println("Winter is coming")
	}

	temp := 20

	switch {
	case temp < 0:
		fmt.Println("its cold")

	case temp > 0 && temp < 15:
		fmt.Println("mild cold")

	case temp > 15 && temp < 30:
		fmt.Println("Summer weather ")

	default:
		fmt.Println("Very hot weather")

	}
}
