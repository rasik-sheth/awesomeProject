package mainx

import "fmt"

func fA() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
func closure() {
	fB := fA()
	fc := fA()
	fmt.Print(fB())
	fmt.Print(fB())
	fmt.Print(fc())
	fmt.Print(fc())
}
func fz(int) string {
	return ""
}
