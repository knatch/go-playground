/*
Write a program which prompts the user to enter integers and stores the integers in a sorted slice. The program should be written as a loop. Before entering the loop, the program should create an empty integer slice of size (length) 3. During each pass through the loop, the program prompts the user to enter an integer to be added to the slice. The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order. The slice must grow in size to accommodate any number of integers which the user decides to enter. The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.
*/

package main

import (
	"fmt"
	"strconv"
	"sort"
	"strings"
)

func main() {
	var input string
	var exit bool = false
	var counter int = 0
	inputArr := make([]int, 3)

	for exit == false {
		fmt.Printf("Enter your number: ")
		fmt.Scan(&input)

		// getting cannot use type untyped rune as type string error
		// TODO: find a better way to parse user input
		str := strings.ToUpper(input)
		if str[0] == 'X' {
			exit = true
			continue;
		}

		// parse string to int
		i, err := strconv.Atoi(input)
		if err != nil {
			// handle error
			// fmt.Println(err)
			fmt.Println("Invalid number")
		} else {
			if counter < 3 {
				inputArr[counter] = i
				counter++
			} else {
				inputArr = append(inputArr, i)
				sort.Ints(inputArr)
			}
			fmt.Println("Sorted: ", inputArr)
		}
	}

	fmt.Println("exit.")	
}