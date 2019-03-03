package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type name struct {
	firstName string
	lastName  string
}

func main() {
	slice := make([]name, 0, 3)

	fmt.Println("Enter file name: ")
	reader := bufio.NewReader(os.Stdin)
	fileName, _ := reader.ReadString('\n')
	fileName = strings.TrimSuffix(fileName, "\n")

	file, err := os.Open(fileName)
	if err != nil {
		log.Println(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		nameArray := strings.Split(scanner.Text(), " ")
		slice = append(slice, name{firstName: nameArray[0], lastName: nameArray[1]})
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	for _, name := range slice {
		fmt.Printf("First Name: %s Last Name: %s\n", name.firstName, name.lastName)
	}
}
