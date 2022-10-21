package main

import (
	"fmt"
)

func main() {
	/*
		const PI1 float64 = 3.1415
		const PI2 = 3.14

		fmt.Printf("PI1: %v\n", PI1) //3.1415
		fmt.Printf("PI2: %v\n", PI2) //3.14

		const (
			a = 123
			b = 456
		)
		fmt.Printf("a: %v\n", a) //123
		fmt.Printf("b: %v\n", b) //456

		const (
			c = 1
			d
			e
		) // 都赋一样的值
		fmt.Printf("c: %v\n", c) //1
		fmt.Printf("d: %v\n", d) //1
		fmt.Printf("e: %v\n", e) //1
	*/

	//1
	const a = iota
	const b = iota
	fmt.Printf("a: %v\n", a) //0
	fmt.Printf("b: %v\n", b) //0

	//2
	const (
		c = iota
		d = iota
	)
	fmt.Printf("c: %v\n", c) //0
	fmt.Printf("d: %v\n", d) //1

	//3
	const (
		e = iota
		_
		g = iota
	)
	fmt.Printf("e: %v\n", e) //0
	fmt.Printf("g: %v\n", g) //2

	//4
	const (
		h = iota
		i = 100
		j = iota
	)
	fmt.Printf("h: %v\n", h) //0
	fmt.Printf("i: %v\n", i) //100
	fmt.Printf("j: %v\n", j) //2
}
