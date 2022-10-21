package main

import "fmt"

func k() {

}

func main() {

	//1
	a := 1
	b := true
	c := "hello"
	fmt.Printf("%T\n", a) //int
	fmt.Printf("%T\n", b) //bool
	fmt.Printf("%T\n", c) //string

	//2
	d := 3.14
	p := &d
	fmt.Printf("%T\n", p) //float64

	//3
	e := [3]int{1, 2, 3}
	fmt.Printf("%T\n", e) //[3]int

	//4
	f := []int{1, 2}
	fmt.Printf("%T\n", f) //[]int

	//5
	fmt.Printf("%T\n", k) //func()

}
