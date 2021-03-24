package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
}

type Bird struct {
}

type Snake struct {
}

func (a Cow) Eat() {
	fmt.Println("grass")
}

func (a Cow) Move() {
	fmt.Println("walk")
}

func (a Cow) Speak() {
	fmt.Println("moo")
}

func (a Bird) Speak() {
	fmt.Println("peep")
}

func (a Bird) Eat() {
	fmt.Println("worms")
}

func (a Bird) Move() {
	fmt.Println("fly")
}

func (a Snake) Speak() {
	fmt.Println("hsss")
}

func (a Snake) Eat() {
	fmt.Println("mice")
}

func (a Snake) Move() {
	fmt.Println("slither")
}

func newAnimal(animalType string, name string) {
	switch animalType {
	case "cow":
		allAnimals[name] = Cow{}
		fmt.Println("Cow Added")
	case "bird":
		allAnimals[name] = Bird{}
		fmt.Println("Bird Added")
	case "snake":
		allAnimals[name] = Snake{}
		fmt.Println("Snake Added")
	default:
		fmt.Println("Unknown animal, not supported at this time")
	}
}

func queryAnimal(name string, action string) {
	animal := allAnimals[name]
	switch action {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	default:
		fmt.Println("Unknown actions")
	}
}

var allAnimals = make(map[string]Animal)

func main() {

	for {
		fmt.Print("\nEnter New Animal [newanimal <name> <type>] or Query Animal [query <name> <action>] : ")
		in := bufio.NewReader(os.Stdin)
		line, err := in.ReadString('\n')
		curatedLine := strings.Trim(line, "\n\r")

		if err == nil {
			splitStr := strings.Split(curatedLine, " ")

			if len(splitStr) == 3 {

				if splitStr[0] == "newanimal" {
					newAnimal(splitStr[2], splitStr[1])
				} else if splitStr[0] == "query" {
					queryAnimal(splitStr[1], splitStr[2])
				} else {
					fmt.Println("Unsupported command")
				}
			} else {
				fmt.Println("3 arguments needed")
			}
		}
	}
}
