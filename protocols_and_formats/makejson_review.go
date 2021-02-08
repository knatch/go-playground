package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

/* // Person Struct
type person struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

func newPerson(name string, address string) *person {
	p := person{Name: name}
	p.Address = address
	return &p
} */

func main() {
	inputScanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Please input name\n: ")
	var userInputName string
	if inputScanner.Scan() {
		userInputName = inputScanner.Text()
	}

	fmt.Printf("Please input address\n: ")
	var userInputAddress string
	if inputScanner.Scan() {
		userInputAddress = inputScanner.Text()
	}

	//person := newPerson(userInputName, userInputAddress)

	//Making a map, for the project. Although Struct is probably better
	personMap := make(map[string]string)

	personMap["address"] = userInputAddress
	personMap["name"] = userInputName

	jsonOutput, _ := json.Marshal(personMap)

	os.Stdout.Write(jsonOutput)
	println("\n")

}