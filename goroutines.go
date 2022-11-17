package main

import (
	"fmt"
	"time"
)

func task(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("msg: %v\n", msg)
	}
}

func main() {
	go task("hello")
	task("go")
	fmt.Printf("main\n")
	time.Sleep(time.Microsecond * 1000)
}
