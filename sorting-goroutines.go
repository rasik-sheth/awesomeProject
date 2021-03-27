package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
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

func merger(mainSlice []int, addSlice []int) []int {
	if len(mainSlice) == 0 {
		mainSlice = addSlice
		return mainSlice
	} else {
		merged := []int{}
		pointera, pointerb := 0, 0

		for {
			if pointera < len(mainSlice) && pointerb < len(addSlice) {
				if mainSlice[pointera] < addSlice[pointerb] {
					merged = append(merged, mainSlice[pointera])
					pointera++
				} else {
					merged = append(merged, addSlice[pointerb])
					pointerb++
				}
			} else {
				break
			}
		}
		//fmt.Print("\nMerged Sorted Arrays: ", merged)
		for {
			if pointera < len(mainSlice) {
				merged = append(merged, mainSlice[pointera])
				pointera++
			} else {
				break
			}
		}

		for {
			if pointerb < len(addSlice) {
				merged = append(merged, addSlice[pointerb])
				pointerb++
			} else {
				break
			}
		}
		return merged
	}

}

func sortArr(wg *sync.WaitGroup, a []int) {
	defer wg.Done()
	BubbleSort(a)
}

func buildArrays(slice []int, chunkSize int) [4][]int {
	var chunks [4][]int
	cnt := 0
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize

		// necessary check to avoid slicing beyond
		// slice capacity
		if end > len(slice) {
			end = len(slice)
		}

		chunks[cnt] = slice[i:end]
		cnt++
	}

	return chunks
}

func main() {
	testInput := []int{
		5, 8, 1, -1, -12, 89, 11, 23, 22, 48, 99, 98, 76, 13, 9,
		4, 9, 13, 15, 7, 90, 91, 700, 699, 600, 300, 305,
		401, 901, 123, 151, 17, 110, 222, 700, 204, 599, 300, 305,
		504, 811, 111, -123, -152, 809, 1011, 2113, 202, 498, 919, 908, 716, 123, 922,
	}

	fmt.Print("\nEnter Integers, seperated by space, press enter to sort them, enter d to use default input: ")
	in := bufio.NewReader(os.Stdin)
	line, err := in.ReadString('\n')
	curatedLine := strings.Trim(line, "\n\r")
	var intSlice []int

	if curatedLine == "d" {
		intSlice = testInput
	} else if err == nil {
		splitStr := strings.Split(curatedLine, " ")

		for _, s := range splitStr {
			i, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal("Incorrect input, only integers are allowed", s)
			}
			intSlice = append(intSlice, i)
		}
	}

	chunkLength := 0
	if len(intSlice)%4 == 0 {
		chunkLength = len(intSlice) / 4
	} else {
		chunkLength = len(intSlice)/4 + 1
	}

	allArrs := buildArrays(intSlice, chunkLength)
	for j := 0; j < len(allArrs); j++ {
		fmt.Print("\nIndividual arrays: ", allArrs[j])
		fmt.Print("\n Lenth: ", len(allArrs[j]))
	}

	var wg sync.WaitGroup
	for j := 0; j < len(allArrs); j++ {
		wg.Add(1)
		go sortArr(&wg, allArrs[j])
	}
	wg.Wait()

	allMerged := []int{}
	for j := 0; j < len(allArrs); j++ {
		allMerged = merger(allMerged, allArrs[j])
	}
	fmt.Print("\n\nMerged Sorted Arrays: ", allMerged)
	fmt.Print(" Length: ", len(allMerged))
}
