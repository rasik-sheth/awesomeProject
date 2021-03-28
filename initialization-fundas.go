package main

import "fmt"

type PersonInfo struct {
	name, addr string
}

func printN(p PersonInfo) {
	fmt.Printf("\nprintP %s %s ", p.name, p.addr)
}

func printP(p *PersonInfo) {
	fmt.Printf("\nprintP %s %s ", p.name, p.addr)
}

func main() {
	//
	personInfo := PersonInfo{}
	fmt.Printf("\nPrint default inside %s %s ", personInfo.name, personInfo.addr)

	personInfo1 := PersonInfo{name: "Rasik", addr: "MM"}
	fmt.Printf("\nPrint default inside %s %s ", personInfo1.name, personInfo1.addr)

	var personInfo2 PersonInfo
	fmt.Printf("\nPrint default inside %s %s ", personInfo2.name, personInfo2.addr)

	printN(personInfo1)
	printP(&personInfo1)

	personInfo3 := new(PersonInfo)
	personInfo3.addr = "118"
	personInfo3.name = "viraaj"
	fmt.Printf("\nPrint default inside %s %s ", personInfo3.name, personInfo3.addr)
	printP(personInfo3)

}
