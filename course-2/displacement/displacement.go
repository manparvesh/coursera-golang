package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

/**
Let us assume the following formula for displacement s as a function of time t, acceleration a, initial velocity vo, and initial displacement so.

s =½ a t2 + vot + so

Write a program which first prompts the user to enter values for acceleration, initial velocity, and initial displacement. Then the program should prompt the user to enter a value for time and the program should compute the displacement after the entered time.

You will need to define and use a function called GenDisplaceFn() which takes three float64 arguments, acceleration a, initial velocity vo, and initial displacement so. GenDisplaceFn() should return a function which computes displacement as a function of time, assuming the given values acceleration, initial velocity, and initial displacement. The function returned by GenDisplaceFn() should take one float64 argument t, representing time, and return one float64 argument which is the displacement travelled after time t.

For example, let’s say that I want to assume the following values for acceleration, initial velocity, and initial displacement: a = 10, vo = 2, so = 1. I can use the following statement to call GenDisplaceFn() to generate a function fn which will compute displacement as a function of time.

fn := GenDisplaceFn(10, 2, 1)

Then I can use the following statement to print the displacement after 3 seconds.

fmt.Println(fn(3))

And I can use the following statement to print the displacement after 5 seconds.

fmt.Println(fn(5))
*/

// GenDisplaceFn generates function based on the given values
func GenDisplaceFn(acceleration, initialVelocity, initialDisplacement float64) func(float64) float64 {
	function := func(time float64) float64 {
		return 0.5*acceleration*math.Pow(time, 2) + initialVelocity*time + initialDisplacement
	}
	return function
}

func getFloatInput() float64 {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	floatValue, _ := strconv.ParseFloat(input, 64)
	return floatValue
}

func promptAndGetValue(name string) float64 {
	fmt.Println("Enter value for " + name + ": ")
	return getFloatInput()
}

func main() {
	acceleration := promptAndGetValue("acceleration")
	initialVelocity := promptAndGetValue("initialVelocity")
	initialDisplacement := promptAndGetValue("initialDisplacement")
	time := promptAndGetValue("time")

	function := GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)
	fmt.Println(function(time))
}
