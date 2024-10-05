package main

import (
	"fmt"
	"time"
)

func main() {

	currentTime := time.Now()
	fmt.Println("Print the time ", currentTime)
	//fmt.Printf("The current time type is  %T\n", currentTime)

	//waqt := currentTime.Format("02-01-2006, Monday")   - here we are declaring date and day
	//waqt := currentTime.Format("02-01-2006, 15:04:05")
	waqt := currentTime.Format("02-01-2006, 15:04PM")

	fmt.Println("waqt :", waqt)

	layout_str := "2006-01-02"
	date_time := "2026-09-25" // converting string to time format
	formatted_str, _ := time.Parse(layout_str, date_time)
	fmt.Println("", formatted_str)

	// add 1 day in current time format
	new_date := currentTime.Add(24 * time.Hour)
	fmt.Println("print new date", new_date)
	newFormat := new_date.Format("02-01-2006, Monday")
	fmt.Println("print the newFormat:", newFormat)

}
