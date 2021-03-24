package mainx

import (
	"fmt"
	"sort"
	"strconv"
)

func slice() {
	sl := make([]int, 0, 3)
	for {
		fmt.Println("\nEnter integer number or x to exit ")
		var st string
		fmt.Scan(&st)

		if st == "x" {
			break
		} else {
			val, err := strconv.Atoi(st)
			if err == nil {
				//use append instead, append takes care of resizing if needed
				sl = append(sl, val)
				sort.Ints(sl)

				fmt.Print("Sorted:  ")
				for _, v := range sl {
					fmt.Printf(" %d", v)
				}

			} else {
				fmt.Println("Not a valid number")
			}
		}
	}
}
