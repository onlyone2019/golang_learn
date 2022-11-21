/*
goroutine 池并发实现TCP端口扫描 强化版

主进程接收端口信息 排序后统一输出
*/

package main

import (
	"fmt"
	"net"
	"sort"
)

func worker_(ports chan int, result chan int) {
	for p := range ports {
		address := fmt.Sprintf("192.168.0.168:%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			result <- p
			continue
		}
		result <- -1 // 开着的端口就给一个-1 ， 排序后这几个-1会排在前面
		conn.Close()
	}
}

func main() {
	ports := make(chan int, 100)
	result := make(chan int)
	defer close(ports)
	defer close(result)

	var status []int

	for i := 0; i < cap(ports); i++ {
		go worker_(ports, result) //启100个worker 接收处理ports中的数据
	}

	go func() {
		for i := 1; i < 1024; i++ { //往ports里面丢数据
			ports <- i
		}
	}()

	for i := 1; i < 1024; i++ {
		x := <-result //阻塞地从result中拿1024个数据
		status = append(status, x)
	}

	sort.Ints(status)
	for _, v := range status {
		fmt.Printf("%v\n", v)
	}
}
