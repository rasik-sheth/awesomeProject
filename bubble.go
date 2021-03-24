package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func swap(intSlice []int, a int) {
	x := intSlice[a]
	intSlice[a] = intSlice[a+1]
	intSlice[a+1] = x
}

func BubbleSort(intSlice []int) {
	for i := 0; i < len(intSlice); i++ {
		for j := 0; j < len(intSlice)-1; j++ {
			if intSlice[j] > intSlice[j+1] {
				swap(intSlice, j)
			}
		}
	}
}

func main() {
	fmt.Print("\nEnter Integers, seperated by space, press enter to sort them: ")
	in := bufio.NewReader(os.Stdin)
	line, err := in.ReadString('\n')
	curatedLine := strings.Trim(line, "\n\r")

	var intSlice []int
	if err == nil {
		splitStr := strings.Split(curatedLine, " ")

		for _, s := range splitStr {
			i, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal("Incorrect input, only integers are allowed %s", s)
			}
			intSlice = append(intSlice, i)
		}

		BubbleSort(intSlice)
		fmt.Print("Sorted Array: ", intSlice)
	}

}
