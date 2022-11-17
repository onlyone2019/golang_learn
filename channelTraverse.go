/*
channel的遍历

2022-11-12
*/

package main

import (
	"fmt"
	"time"
)

var channel2 = make(chan int)

func write() {
	for i := 0; i < 2; i++ {
		channel2 <- i
	}
	close(channel2) // 注意这里要及时关闭 不然遍历时会报错 fatal error: all goroutines are asleep - deadlock!
}

func main() {
	go write()
	time.Sleep(time.Microsecond)
	// method 1
	// for {
	// 	data, ok := <-channel2
	// 	if ok {
	// 		fmt.Printf("data: %v\n", data)
	// 	} else {
	// 		break
	// 	}
	// }

	// method 2
	for v := range channel2 {
		fmt.Printf("v: %v\n", v)
	}

	v := <-channel2
	fmt.Printf("v: %v\n", v) //这里读到的值为0
	//如果通道里面没有值了，接着读 读出来的值就是默认值，比如int类型就是0
	//但是如果通道没关闭，接着读就会出问题。
}
