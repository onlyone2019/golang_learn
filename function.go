package main

import "fmt"

func f1(a ...int) {
	for _, v := range a {
		fmt.Printf("v: %v\n", v)
	}
}

func f2(name string, a ...int) {
	fmt.Printf("v: %v\n", name)
	for _, v := range a {
		fmt.Printf("v: %v\n", v)
	}
}

func sum(a int, b int) int {
	return a + b
}

func minu(a int, b int) int {
	return a - b
}

func calculate(f int) func(int, int) int {
	switch f {
	case 1:
		return sum
	case 2:
		return minu
	default:
		return minu
	}
}

func sayhello(name string) {
	fmt.Printf("hello %v!\n", name)
}

func saybye(name string) {
	fmt.Printf("Bye %v~\n", name)
}

func say(name string, function func(string)) {
	function(name)
}

func bibao() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

func main() {
	// f2("111", 1, 2, 3, 4, 5)
	// var f func(int, int) int
	// f = calculate(1)
	// s := f(1, 2)
	// fmt.Printf("s: %v\n", s)
	// f = calculate(2)
	// s = f(2, 1)
	// fmt.Printf("s: %v\n", s)

	// say("TT", sayhello)
	// say("TT", saybye)

	fff := bibao()
	fmt.Printf("fff(1): %v\n", fff(1)) //fff(1): 1
	fmt.Printf("fff(2): %v\n", fff(2)) //fff(2): 3
	fff = bibao()
	fmt.Printf("fff(1): %v\n", fff(1)) //fff(1): 1
}
