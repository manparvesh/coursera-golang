package main

import (
	"fmt"
)

func main() {
	fmt.Print("Enter float: ")
	var input float64
	_, err := fmt.Scanln(&input)
	if err != nil {
		panic("Can't read float")
	}
	fmt.Printf("%d", int(input))
}
