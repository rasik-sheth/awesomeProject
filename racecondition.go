package main

/*
The program is simple.

Contains 2 subroutines,
#1 Adds current value of 'z' with 1 and stores the value back in z
#2 Multiples current value of 'z' with 2 and stores the value back in z.

Both routines execute concurrently.
1. The outcome is printed after both the goroutines are complete.
2. Since both the goroutines are modifying the space integer through pointer, the outcome is dependent on time of goroutine execution.
3. Race condition is evident as outcome of the program is not consistent instead heavily depends of the order of execution
*/

import (
	"fmt"
	"sync"
	"time"
)

func multiplier(wg *sync.WaitGroup, i *int, j int) {
	defer wg.Done()
	for v := 0; v < 25; v++ {
		*i = *i * j
		time.Sleep(time.Millisecond * 20)
		//fmt.Printf(" %d", *i)
	}
	time.Sleep(time.Second)
}

func adder(wg *sync.WaitGroup, i *int, j int) {
	defer wg.Done()
	for v := 0; v < 25; v++ {
		*i = *i + j
		time.Sleep(time.Millisecond * 20)
		//fmt.Printf("\nFinal Value of i: %d", *i)
	}
	time.Sleep(time.Second)
}

func main() {
	var z *int
	z = new(int)
	*z = 2
	var wg sync.WaitGroup

	fmt.Printf("\nInitial Value of z: %d", *z)
	wg.Add(1)
	go multiplier(&wg, z, 2)
	wg.Add(1)
	go adder(&wg, z, 1)
	wg.Wait()

	fmt.Printf("\nFinal Value of z[value will every iteration]: %d", *z)

}
