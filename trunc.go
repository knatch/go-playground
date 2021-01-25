package main

import (
	"fmt"
	"math"
)

func main() {
	var floatInput float64
	var truncated int

	fmt.Printf("Enter your number: ")
	fmt.Scan(&floatInput)

	truncated = int(math.Trunc(floatInput))
	fmt.Println(truncated)
}