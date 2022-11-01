// 多个类型实现同一个接口 类似多态
package main

import "fmt"

type Pet interface {
	eat()
}

type Dog struct {
}

type Cat struct {
}

func (d Dog) eat() {
	fmt.Printf("Dog eat...\n")
}

func (c Cat) eat() {
	fmt.Printf("Cat eat...\n")
}

func main() {
	// 方式1
	// var cat Cat
	// var dog Dog
	// cat = Cat{}
	// dog = Dog{}
	// cat.eat()
	// dog.eat()

	//方式2
	var pet Pet
	pet = Dog{}
	pet.eat()
	pet = Cat{}
	pet.eat()
}
