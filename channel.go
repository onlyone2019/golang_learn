package main

import (
	"fmt"
	"math/rand"
	"time"
)

var channel = make(chan int) //无缓冲的通道

func send() {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(10)
	fmt.Printf("send value: %v\n", value)
	// time.Sleep(time.Second * 5)
	channel <- value
}

func main() {
	// 从通道接收数据
	defer close(channel) // 主程序结束时 关闭通道
	go send()            //启动一个协程
	fmt.Printf("wait...\n")
	value := <-channel // 这里会阻塞  等到有值在接收
	fmt.Printf("receive value: %v\n", value)
	fmt.Printf("end ... \n")
}
