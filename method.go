package main

import "fmt"

type Person struct {
	name string
}

func (pp Person) showname1() {
	pp.name = "111"
	fmt.Printf("pp: %v\n", pp)
}

func (pp *Person) showname2() {
	pp.name = "222"
	fmt.Printf("pp: %v\n", pp)
}

type myint int

func (xx myint) show() {
	fmt.Printf("xx: %v\n", xx)
}

func main() {
	var p = Person{
		name: "name",
	}
	p.showname1()
	fmt.Printf("p: %v\n ===============\n", p)
	p.showname2()
	fmt.Printf("p: %v\n", p)

	var x myint
	x = 1
	x.show()
}
