package main

import (
	"fmt"
	"strconv"
)

func trunc() {
	for {
		fmt.Println("\nEnter float number or x to exit ")
		var st string
		fmt.Scan(&st)

		if( st == "x"){
			break;
		} else{
			val, err := strconv.ParseFloat(st,64)
			if(err == nil){
				fmt.Printf("Truncated number is %d", int(val))
			} else{
				fmt.Println("Not a valid number or a float")
			}
		}
	}
}
