package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	userInfo := make(map[string]string)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln("Unable to read name")
	}
	userInfo["name"] = strings.Replace(name, "\n", "", -1)

	fmt.Println("Enter address: ")
	address, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln("Unable to read address")
	}
	userInfo["address"] = strings.Replace(address, "\n", "", -1)

	jsonObj, err := json.Marshal(userInfo)
	if err != nil {
		log.Fatalln("Unable to create json object")
	}

	fmt.Printf("\n%s\n", jsonObj)
}
