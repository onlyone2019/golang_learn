//接口的OCP原则

package main

import "fmt"

type Pet interface {
	eat()
}

type Dog struct {
}

type Cat struct {
}

func (d Dog) eat() {
	fmt.Printf("Dog eat...\n")
}

func (c Cat) eat() {
	fmt.Printf("Cat eat...\n")
}

type Person struct {
}

func (p Person) care(pet Pet) {
	pet.eat()
}

func main() {
	pp := Person{}
	var pet Pet
	pet = Dog{}
	pp.care(pet)

	pet = Cat{}
	pp.care(pet)
}
