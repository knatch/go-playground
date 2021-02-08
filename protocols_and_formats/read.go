/*
Write a program which reads information from a file and represents it in a slice of structs. Assume that there is a text file which contains a series of names. Each line of the text file has a first name and a last name, in that order, separated by a single space on the line. 

Your program will define a name struct which has two fields, fname for the first name, and lname for the last name. Each field will be a string of size 20 (characters).

Your program should prompt the user for the name of the text file. Your program will successively read each line of the text file and create a struct which contains the first and last names found in the file. Each struct created will be added to a slice, and after all lines have been read from the file, your program will have a slice containing one struct for each line in the file. After reading all lines from the file, your program should iterate through your slice of structs and print the first and last names found in each struct.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct  {
	fname string
	lname string
}

func main() {
	var people = make([]Name, 0)
	inputScanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Enter file to read: ")
	var fileName string
	if inputScanner.Scan() {
		fileName = inputScanner.Text()
	}

	// fmt.Println(fileName)

	f, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	defer f.Close()
	

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		readName := scanner.Text()
		
		tempName := strings.Split(readName, " ")
		if len(tempName) >= 2 {
			var p Name
			p.fname = tempName[0]
			p.lname = tempName[1]
			
			people = append(people, p)
		}
	}
	
	if err := scanner.Err(); err != nil {
		fmt.Printf("%v\n", err)
	}

	fmt.Println("List of names: ", people)
}