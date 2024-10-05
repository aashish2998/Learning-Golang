package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "apple,orange,grapes"
	data := strings.Split(str, ",")
	fmt.Println("Print the data value", data)

	cnt := "one two five three four two three"
	num := strings.Count(cnt, "four")
	fmt.Println("Print the num value", num)

	space := "   Hello ! brother "
	sol := strings.TrimSpace(space)
	fmt.Println("Print the trimmed value:", sol)

	str1 := "Aashish"
	str2 := "kushwaha"
	res := strings.Join([]string{str1, str2}, " ")
	fmt.Println("joined the word: ", res)

}
