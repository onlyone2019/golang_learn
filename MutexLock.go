/*
2022-11-12
线程间共享变量的加锁  Mutex锁 样例

理论上，这个程序执行完后 i 的值还得是 3，如果不加锁，那结果就未知了

*/

package main

import (
	"fmt"
	"sync"
	"time"
)

var i int = 3

var lock sync.Mutex

func a() {
	lock.Lock()
	i++
	time.Sleep(time.Millisecond * 10)
	fmt.Printf("i ++ : %v\n", i)
	lock.Unlock()
}

func b() {
	lock.Lock()
	i--
	fmt.Printf("i --: %v\n", i)
	lock.Unlock()
}

func main() {
	for t := 0; t < 20; t++ {
		go a()
		go b()
	}
	time.Sleep(time.Second)
	fmt.Printf("end i: %v\n", i)
}
