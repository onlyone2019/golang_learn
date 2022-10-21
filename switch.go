package main

import "fmt"

func f1() {
	a := 100
	switch a {
	case 100:
		fmt.Printf("a: %v\n", a)
	case 200:
		fmt.Printf("a: %v\n", a)
	default:
		fmt.Printf("other")
	}
}

func f2() {
	a := 400
	switch {
	case a <= 100:
		fmt.Printf("a: %v\n", a)
	case a > 100:
		fmt.Printf("a: %v\n", a)
	default:
		fmt.Printf("other")
	}
}

func f3() {
	a := 4
	switch a {
	case 1, 2, 3:
		fmt.Printf("a: %v\n", a)
	case 4:
		fmt.Printf("a: %v\n", a)
	default:
		fmt.Printf("other")
	}
}

func f4() {
	a := 2
	switch a {
	case 1, 2, 3:
		fmt.Printf("1,2,3 a: %v\n", a)
		fallthrough
	case 4:
		fmt.Printf("4 a: %v\n", a)
	case 5:
		fmt.Printf("5 a: %v\n", a)
	default:
		fmt.Printf("other")
	}
}

func main() {
	f4()
}
