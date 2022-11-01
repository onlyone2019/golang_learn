// 接口嵌套
package main

import "fmt"

type Fly interface {
	fly()
}

type Swim interface {
	swim()
}

type FlyFish interface {
	Fly
	Swim
}

type Fish struct {
}

func (fish Fish) fly() {
	fmt.Println("fly")
}

func (fish Fish) swim() {
	fmt.Println("swim")
}

func main() {
	var ff FlyFish
	ff = Fish{}
	ff.fly()
	ff.swim()
}
