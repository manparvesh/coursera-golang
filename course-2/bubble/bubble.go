package bubble

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/**
Write a Bubble Sort program in Go. The program should prompt the user to type in a sequence of up to 10 integers. The program should print the integers out on one line, in sorted order, from least to greatest. Use your favorite search tool to find a description of how the bubble sort algorithm works.

As part of this program, you should write a function called BubbleSort() which takes a slice of integers as an argument and returns nothing. The BubbleSort() function should modify the slice so that the elements are in sorted order.

A recurring operation in the bubble sort algorithm is the Swap operation which swaps the position of two adjacent elements in the slice. You should write a Swap() function which performs this operation. Your Swap() function should take two arguments, a slice of integers and an index value i which indicates a position in the slice. The Swap() function should return nothing, but it should swap the contents of the slice in position i with the contents in position i+1.
*/

func swap(ar []int, i int) {
	if ar[i] > ar[i+1] {
		temp := ar[i]
		ar[i] = ar[i+1]
		ar[i+1] = temp
	}
}

func BubbleSort(ar []int) {
	end := len(ar)
	for i := 0; i < len(ar)-1; i++ {
		log.Println("Sort number: ", i+1)

		for index := range ar[:end-1] {
			swap(ar[:end], index)
		}
		end--
	}
}

func main() {
	fmt.Println("Enter a sequence of up to 10 integers separated by spaces:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	slice := make([]string, 0)
	slice = strings.Split(input, " ")

	nums := make([]int, len(slice))
	for i, ss := range slice {
		nums[i], _ = strconv.Atoi(ss)
	}

	BubbleSort(nums)

	fmt.Println("Sorted sequence:")
	fmt.Println(nums)
}
