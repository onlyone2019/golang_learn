package main

import "fmt"

func f1() {
	for i := 1; i < 10; i++ {
		fmt.Printf("i: %v\n", i)
	}
}

func f2() {
	i := 1
	for i < 10 {
		fmt.Printf("i: %v\n", i)
		i++
	}
	for ; i < 10; i++ {
		fmt.Printf("i: %v\n", i)
	}
}

func f3() {
	var a = [5]int{5, 4, 3, 2, 1}
	for i, x := range a {
		fmt.Printf("i=%v x=%v\n", i, x)
	}
}

func main() {
	f3()
}
