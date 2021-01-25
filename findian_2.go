package main

import (
	"fmt"
	"strings"
)

func main()  {
	var s string
	fmt.Print("Your string: ")
	_, err := fmt.Scan(&s)
	if (err!=nil) {
		panic(err)
	}
	s = strings.ToLower(s)
	s = strings.Trim(s, "\n ")
	if (strings.HasPrefix(s, "i") && strings.Contains(s, "a") && strings.HasSuffix(s, "n")) {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}