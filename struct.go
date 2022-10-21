package main

import "fmt"

type PPerson struct {
	id   int
	name string
}

type Dog struct {
	name string
}

type Ren struct {
	name string
	dog  Dog
}

func main() {
	var p *PPerson
	p = &PPerson{}
	p.id = 1
	p.name = "111"
	fmt.Printf("p: %p\n", p)
	fmt.Printf("p: %v\n", *p)

	xx := PPerson{
		id:   2,
		name: "222",
	}
	fmt.Printf("xx: %v\n", xx)

	var yy = new(PPerson)
	yy.id = 3
	yy.name = "333"
	fmt.Printf("yy: %p\n", yy)
	fmt.Printf("yy: %v\n", yy)

	r1 := Ren{}
	r1.name = "r1"
	r1.dog.name = "kate"
	fmt.Printf("r1: %v\n", r1)

}
