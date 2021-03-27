package main

import (
	"fmt"
	"time"
)

func main() {

	const iterationsCount = 10
	var a int

	go func() {
		for i := 0; i < iterationsCount; i++ {
			a = 1
			time.Sleep(1 * time.Millisecond)
			a = 0
			time.Sleep(1 * time.Millisecond)
		}
		fmt.Println("Goroutine #1 completed")
	}()

	go func() {
		for i := 0; i < iterationsCount; i++ {
			if a == 1 {
				fmt.Printf("Observed a == 1 at iteration #%v!\n", i)
			}
			time.Sleep(1 * time.Millisecond)
		}
		fmt.Println("Goroutine #2 completed")
	}()

	time.Sleep(1 * time.Second)
}
