package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func nestedSort(waitGroup *sync.WaitGroup, integers []int) {
	fmt.Println(integers)
	sort.Ints(integers)
	waitGroup.Done()
}

func mergeSortedArrays(smallSize int, numbers []int) (sortedIntegers []int) {
	c1 := numbers[:1*smallSize]
	c2 := numbers[1*smallSize : 2*smallSize]
	c3 := numbers[2*smallSize : 3*smallSize]
	c4 := numbers[3*smallSize:]

	for k := 0; k < len(numbers); k++ {
		i := 0
		j := 0
		if len(c1) != 0 {
			i = c1[0]
			j = 1
		}
		if len(c2) != 0 {
			if j == 0 {
				i = c2[0]
				j = 2
			} else {
				if c2[0] < i {
					i = c2[0]
					j = 2
				}
			}
		}
		if len(c3) != 0 {
			if j == 0 {
				i = c3[0]
				j = 3
			} else {
				if c3[0] < i {
					i = c3[0]
					j = 3
				}
			}
		}
		if len(c4) != 0 {
			if j == 0 {
				i = c4[0]
				j = 4
			} else {
				if c4[0] < i {
					i = c4[0]
					j = 4
				}
			}
		}
		sortedIntegers = append(sortedIntegers, i)
		switch j {
		case 1:
			c1 = c1[1:]
		case 2:
			c2 = c2[1:]
		case 3:
			c3 = c3[1:]
		case 4:
			c4 = c4[1:]
		}
	}

	return
}

func main() {
	// take input
	fmt.Println("Enter some integers that you want to sort:")
	reader := bufio.NewReader(os.Stdin)
	array, _, _ := reader.ReadLine()
	numberStrings := strings.Split(string(array), " ")
	var numbers []int
	for _, str := range numberStrings {
		num, _ := strconv.Atoi(str)
		numbers = append(numbers, num)
	}

	smallSize := len(numbers) / 4
	var waitGroup sync.WaitGroup
	for c := 0; c < 4; c++ {
		waitGroup.Add(1)
		if c != 3 {
			go nestedSort(&waitGroup, numbers[c*smallSize:(c+1)*smallSize])
		} else {
			// last part
			go nestedSort(&waitGroup, numbers[c*smallSize:])
		}
	}
	waitGroup.Wait()

	sortedArray := mergeSortedArrays(smallSize, numbers)
	fmt.Println(sortedArray)
}
