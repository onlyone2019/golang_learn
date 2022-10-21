package main

import "fmt"

func f1() {
	fmt.Printf("f1 : step1\n")
	defer fmt.Printf("f1 : step2\n")
	defer fmt.Printf("f1 : step3\n")
	fmt.Printf("f1 : step5\n")
}

func f2() {
	fmt.Printf("f2 : step1\n")
	defer fmt.Printf("f2 : step2\n")
	defer fmt.Printf("f2 : step3\n")
	fmt.Printf("f2 : step5\n")
}

func main() {
	fmt.Printf("step1\n")
	defer fmt.Printf("step2\n")
	defer fmt.Printf("step3\n")
	defer fmt.Printf("step4\n")
	fmt.Printf("step5\n")
	f1()
	f2()
}
