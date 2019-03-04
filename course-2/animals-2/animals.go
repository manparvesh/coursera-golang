package main

import "fmt"
import "bufio"
import "os"
import "strings"

/*
Write a program which allows the user to create a set of animals and to get information about those animals. Each animal has a name and can be either a cow, bird, or snake. With each command, the user can either create a new animal of one of the three types, or the user can request information about an animal that he/she has already created. Each animal has a unique name, defined by the user. Note that the user can define animals of a chosen type, but the types of animals are restricted to either cow, bird, or snake. The following table contains the three types of animals and their associated data.

Animal	Food eaten	Locomtion method	Spoken sound
cow	grass	walk	moo
bird	worms	fly	peep
snake	mice	slither	hsss
Your program should present the user with a prompt, “>”, to indicate that the user can type a request. Your program should accept one command at a time from the user, print out a response, and print out a new prompt on a new line. Your program should continue in this loop forever. Every command from the user must be either a “newanimal” command or a “query” command.

Each “newanimal” command must be a single line containing three strings. The first string is “newanimal”. The second string is an arbitrary string which will be the name of the new animal. The third string is the type of the new animal, either “cow”, “bird”, or “snake”. Your program should process each newanimal command by creating the new animal and printing “Created it!” on the screen.

Each “query” command must be a single line containing 3 strings. The first string is “query”. The second string is the name of the animal. The third string is the name of the information requested about the animal, either “eat”, “move”, or “speak”. Your program should process each query command by printing out the requested data.

Define an interface type called Animal which describes the methods of an animal. Specifically, the Animal interface should contain the methods Eat(), Move(), and Speak(), which take no arguments and return no values. The Eat() method should print the animal’s food, the Move() method should print the animal’s locomotion, and the Speak() method should print the animal’s spoken sound. Define three types Cow, Bird, and Snake. For each of these three types, define methods Eat(), Move(), and Speak() so that the types Cow, Bird, and Snake all satisfy the Animal interface. When the user creates an animal, create an object of the appropriate type. Your program should call the appropriate method when the user issues a query command.
*/

var slice []Animal

type Animal interface {
	Name() string
	Eat()
	Move()
	Speak()
}

type Cow struct {
	name       string
	food       string
	locomotion string
	sound      string
}

func (x Cow) Name() string {
	return x.name
}

func (x Cow) Eat() {
	fmt.Println(x.food)
}

func (x Cow) Move() {
	fmt.Println(x.locomotion)
}

func (x Cow) Speak() {
	fmt.Println(x.sound)
}

type Bird struct {
	name       string
	food       string
	locomotion string
	sound      string
}

func (x Bird) Name() string {
	return x.name
}

func (x Bird) Eat() {
	fmt.Println(x.food)
}

func (x Bird) Move() {
	fmt.Println(x.locomotion)
}

func (x Bird) Speak() {
	fmt.Println(x.sound)
}

type Snake struct {
	name       string
	food       string
	locomotion string
	sound      string
}

func (x Snake) Name() string {
	return x.name
}

func (x Snake) Eat() {
	fmt.Println(x.food)
}

func (x Snake) Move() {
	fmt.Println(x.locomotion)
}

func (x Snake) Speak() {
	fmt.Println(x.sound)
}

func getNewAnimal(name string, type1 string) Animal {
	var a1 Animal
	switch type1 {
	case "cow":
		cow := Cow{name, "grass", "walk", "moo"}
		a1 = cow
		return a1
	case "bird":
		bird := Bird{name, "worms", "fly", "peep"}
		a1 = bird
		return a1
	case "snake":
		snake := Snake{name, "mice", "slither", "hss"}
		a1 = snake
		return a1
	default:
		return nil
	}
}

func performAction(name string, info string) {
	for _, v := range slice {
		if name == v.Name() {
			switch info {
			case "eat":
				v.Eat()
				continue
			case "move":
				v.Move()
				continue
			case "speak":
				v.Speak()
				continue
			default:
				fmt.Println("No such action defined. Use eat, speak or move.")
			}
		}

	}
}

func main() {
	inputReader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf(">")
		input, _ := inputReader.ReadString('\n')
		input = strings.ToLower(strings.TrimSuffix(input, "\n"))
		command := strings.Split(input, " ")

		if command[0] == "newanimal" {
			animal := getNewAnimal(command[1], command[2])
			fmt.Println(animal)
			if animal != nil {
				fmt.Println("Created it!")
				slice = append(slice, animal)
				fmt.Println(slice)

			}

		} else if command[0] == "query" {
			performAction(command[1], command[2])
		} else {
			fmt.Println("First word needs to be either 'newanimal' or 'query'")
			continue
		}

	}
}
