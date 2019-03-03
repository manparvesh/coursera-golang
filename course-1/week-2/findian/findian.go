package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("Enter a String:")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln("Unable to read string ")
	}

	// remove newlines
	input = strings.Replace(input, "\n", "", -1)
	input = strings.Replace(input, "\r", "", -1)
	// change string to lowercase
	input = strings.ToLower(input)

	check1 := strings.Index(input, "i") == 0
	check2 := strings.Contains(input, "a")
	check3 := strings.LastIndex(input, "n") == (len(input) - 1)

	if check1 && check2 && check3 {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
