package main

import (
	"fmt"
	"net/url"
)

func main() {

	fmt.Println("Learning url")
	myurl := "https://www.geeksforgeeks.org/top-algorithms-and-data-structures-for-competitive-programming/"
	//fmt.Printf("print the url %T\n", myurl)

	parsedUrl, err := url.Parse(myurl)
	//fmt.Printf("print the url %T\n", parsedUrl)

	if err != nil {
		fmt.Println("Print the url error", err)
		return
	}
	//fmt.Println("Print the parsedurl :", parsedUrl)

	fmt.Println("Scheme:", parsedUrl.Scheme)
	fmt.Println("Host :", parsedUrl.Host)
	fmt.Println("Path:", parsedUrl.Path)
	fmt.Println("RawQuery:", parsedUrl.RawQuery)

	// modifying the parsed url
	parsedUrl.Path = "/newPath"
	parsedUrl.RawQuery = "username=aashish29"

	newUrl := parsedUrl.String()

	fmt.Println("print the new url :", newUrl)

}
