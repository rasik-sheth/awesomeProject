package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) Eat() string {
	if a.food == "" {
		return "unknown animal"
	}
	return a.food
}

func (a Animal) Move() string {
	if a.food == "" {
		return "unknown animal"
	}
	return a.locomotion
}

func (a Animal) Speak() string {
	if a.food == "" {
		return "unknown animal"
	}
	return a.noise
}

func main() {
	animalMap := make(map[string]Animal)
	animalMap["cat"] = Animal{food: "grass", locomotion: "walk", noise: "moo"}
	animalMap["bird"] = Animal{food: "worms", locomotion: "fly", noise: "peep"}
	animalMap["snake"] = Animal{food: "mice", locomotion: "slither", noise: "hsss"}

	fmt.Print("\nEnter Animal and Action: ")
	in := bufio.NewReader(os.Stdin)
	line, err := in.ReadString('\n')
	curatedLine := strings.Trim(line, "\n\r")

	if err == nil {
		splitStr := strings.Split(curatedLine, " ")

		if len(splitStr) == 2 {
			animal := splitStr[0]
			aml := animalMap[animal]
			switch action := splitStr[1]; {
			case action == "eat":
				fmt.Println(aml.Eat())
			case action == "move":
				fmt.Println(aml.Move())
			case action == "speak":
				fmt.Println(aml.Speak())
			default:
				fmt.Println("Unknown actions")
			}
		} else {
			fmt.Println("Not right set of arguments")
		}
	}
}
