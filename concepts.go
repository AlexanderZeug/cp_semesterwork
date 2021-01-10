package main

import "fmt"

type Animal struct { // <1>
	age int // <3>
}

func NewAnimal(age int) Animal {
	a := Animal{}
	a.age = age
	return a
}

func (a Animal) SetAge(age int) { // <4>
	a.age = age
}

func (a Animal) GetAge() int {
	return a.age
}

type IMakeSound interface { // <5>
	makeSound()
}

type Dog struct {
	Animal // <2>
	colour string
}

func NewDog(age int, colour string) Dog {
	d := Dog{}
	d.age = age
	d.colour = colour
	return d
}

func (d Dog) SetColour(colour string) {
	d.colour = colour
}

func (d Dog) GetColour() string {
	return d.colour
}

func (d Dog) MakeSound() { // <6>
	fmt.Println("Wuff")
}

func main() { // <7>
	var laika = NewDog(13, "black")
	laika.MakeSound()
	fmt.Println(laika.GetAge())
	fmt.Println(laika.GetColour())
}

/*
output:
 Wuff
 13
 black
*/