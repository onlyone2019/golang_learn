// interface definition
package main

import "fmt"

type USB interface {
	read()
	write()
}

type Computer struct {
}

type Mobile struct {
}

func (c Computer) read() { //接收者类型也可以是指针类型
	fmt.Printf("computer read...\n")
}

func (c Computer) write() {
	fmt.Printf("computer write...\n")
}

func (m Mobile) read() {
	fmt.Printf("mobile read...\n")
}

func (m Mobile) write() {
	fmt.Printf("mobile write...\n")
}

func main() {
	var com Computer
	com.read()
	com.write()
	var mob Mobile
	mob.read()
	mob.write()
}
