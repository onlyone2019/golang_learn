package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	x := Person{name: "zhangsan", age: 22}
	fmt.Printf("%v\n", x)  //{zhangsan 22}
	fmt.Printf("%#v\n", x) //main.Person{name:"zhangsan", age:22} 打印得更详细
	fmt.Printf("%T\n", x)  //main.Person

	p := &x
	fmt.Printf("%p\n", p) //打印p指向的位置

	/*
		%c 字符
		%d 十进制
		%b 二进制
		%s 字符串
	*/
}
