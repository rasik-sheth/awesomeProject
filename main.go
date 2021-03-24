package main

import "fmt"
import "bufio"
import "os"
import "strings"
import "strconv"

var i, j int

func BubbleSorta(sli []int) {
	n := len(sli)
	for i = 0; i < n-1; i++ {
		for j = 0; j < n-i-1; j++ {
			if sli[j] > sli[j+1] {
				Swap(sli, j)
			}
		}
	}
}

func Swap(sli []int, index int) {
	sli[index], sli[index+1] = sli[index+1], sli[index]
}

func main() {
	fmt.Printf("Enter integers splitted by space: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	integers_str := strings.Split(input, " ")
	integers := make([]int, len(integers_str))

	for i, v := range integers_str {
		integers[i], _ = strconv.Atoi(v)
	}

	BubbleSorta(integers)
	fmt.Printf("result: ")
	for _, v := range integers {
		int_str := strconv.Itoa(v)
		fmt.Printf(int_str)
		fmt.Printf(" ")
	}

}
