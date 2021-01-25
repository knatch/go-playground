package main

import (
	"fmt"
	"strings"
)

func main() {
	var strInput string
	var isFound bool // default to false

	fmt.Scan(&strInput)
	str := strings.ToLower(strInput)
	
	if string(str[0]) == "i" && string(str[len(str)-1]) == "n" && strings.Contains(str, "a") {
		isFound = true
	}
	
	if isFound {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}