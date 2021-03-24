package main

import (
	"fmt"
	"log"
	"strconv"
)

func getFloatInput() float64 {
	var st string
	fmt.Scan(&st)
	val, err := strconv.ParseFloat(st, 64)
	if err != nil {
		log.Fatal("Unexpected input")
	}
	return val
}

func GenDispaceFn(acc float64, initvel float64, initdis float64) func(float64) float64 {
	fn := func(t float64) float64 {
		return (0.5 * acc * t * t) + (initvel * t) + initdis
	}
	return fn
}

func main() {
	fmt.Print("Enter accelaration: ")
	acc := getFloatInput()
	fmt.Print("Enter initial velocity: ")
	initvel := getFloatInput()
	fmt.Print("Enter initial displacement: ")
	initdis := getFloatInput()

	fn := GenDispaceFn(acc, initvel, initdis)

	fmt.Print("Enter time: ")
	time := getFloatInput()
	fmt.Printf("Displacement = %f", fn(time))
}
