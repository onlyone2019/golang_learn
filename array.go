package main

import "fmt"

func main() {

	var b []int
	for i := 1; i <= 3; i++ {
		b = append(b, i)
	}
	b = append(b, 1, 2, 3)
	fmt.Printf("b: %v\n", b)

	b = append(b[:2], b[3:]...)
	fmt.Printf("b: %v\n", b) //b: [1 2 1 2 3]

	a := b //a与b内存地址相同
	a[0] = 100
	fmt.Printf("b: %v\n", b) //b: [100 2 1 2 3]

	var c = make([]int, 7)
	copy(c, b)
	c[1] = 99
	fmt.Printf("c: %v\n", c) //c: [100 99 1 2 3]
	fmt.Printf("b: %v\n", b) //b: [100 2 1 2 3]
}
