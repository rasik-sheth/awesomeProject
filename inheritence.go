//Demonstrate Go Inheritenc
package main

import (
	"fmt"
)

type Pet interface {
	Play()
}

type pet struct {
	name string
}

type Dog interface {
	Pet
	Breed() string
}

type dog struct {
	pet
	breed string
}

func NewPet(name string) *pet {
	p := &pet{
		name: name,
	}
	return p
}

func NewDog(name, breed string) *dog {
	d := &dog{
		pet:   pet{name: name},
		breed: breed,
	}
	return d
}

func (p *pet) Play() {
	fmt.Println("I am a rockstar")
}

func Play(p Pet) {
	p.Play()
}

func main() {
	d := NewDog("spot", "pointer")
	Play(d)
}
