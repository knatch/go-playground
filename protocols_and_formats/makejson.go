/*
Write a program which prompts the user to first enter a name, and then enter an address. Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively. Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.
*/

package main

import (
	"fmt"
	"encoding/json"
)

func main() {
	var nameInput string
	var addrInput string
	var person = make(map[string]string)

	fmt.Printf("Enter your name: ")
	fmt.Scan(&nameInput)

	fmt.Printf("Enter your address: ")
	fmt.Scan(&addrInput)

	person["name"] = nameInput
	person["address"] = addrInput

	personObj, err := json.Marshal(person)
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Printf("JSON object: %s", string(personObj))
}