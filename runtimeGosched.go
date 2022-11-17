package main

import (
	"fmt"
	"runtime"
)

func show(msg string) {
	for i := 0; i < 2; i++ {
		fmt.Printf("msg: %v\n", msg)
	}
}

func main() {
	go show("go")
	for i := 0; i < 2; i++ {
		runtime.Gosched() //让出时间片 可以注释这句看效果
		fmt.Println("main...")
	}
	fmt.Print("end...")
}
