package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Print("Enter float: ")
    var input string
    fmt.Scanln(&input)
    fl, _ := strconv.ParseFloat(input, 64)
    fmt.Printf("%.6f", fl)
}