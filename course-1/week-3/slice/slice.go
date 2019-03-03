package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const prompt = "Enter an integer to add to slice, or 'X' to stop loop:"

func main() {
	slice := make([]int, 3)
	i := 1

	for {
		fmt.Println(prompt)

		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')

		if err != nil {
			log.Println("Unknown input. " + prompt)
			continue
		}

		input = strings.ToLower(strings.TrimSuffix(input, "\n"))
		if input == "x" {
			break
		}

		integer, err := strconv.Atoi(input)
		if err != nil {
			continue
		}

		// if we are at the end
		if i < 3 {
			slice[len(slice)-1] = integer
			i++
		} else {
			slice = append(slice, integer)
		}

		sort.Ints(slice)
		fmt.Println(slice)
	}
}
