package main

import "fmt"

func main() {
	var a map[string]int     //声明一个map
	a = make(map[string]int) //分配内存
	fmt.Printf("a: %v\n", a) //a: map[]

	//声明时即初始化
	var b = map[string]int{
		"aaa": 1,
		"bbb": 2,
		"ccc": 3,
	}
	fmt.Printf("b: %v\n", b) //b: map[aaa:1 bbb:2 ccc:3]

	var c map[string]int
	c = make(map[string]int)
	c["ee"] = 6
	c["dd"] = 7
	fmt.Printf("c: %v\n", c) //c: map[dd:7 ee:6]

	//判断键值是否存在
	key := "ee"
	v, ok := c[key]
	fmt.Printf("v: %v\n", v)                  //v: 6
	fmt.Printf("ok: %v\n =========== \n", ok) //ok: true
	key = "ff"
	v, ok = c[key]
	fmt.Printf("v: %v\n", v)                  //v: 0
	fmt.Printf("ok: %v\n =========== \n", ok) //ok: false

	//map遍历
	for k := range c {
		fmt.Printf("k: %v\n", k)
	}

	for k, val := range b {
		fmt.Printf("%v:%v\n", k, val)
	}
}
