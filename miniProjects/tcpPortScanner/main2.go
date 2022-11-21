/*
goroutine 池并发实现TCP端口扫描
*/

package main

import (
	"fmt"
	"net"
	"sync"
)

func worker(ports chan int, wg *sync.WaitGroup) {
	for p := range ports {
		address := fmt.Sprintf("192.168.0.168:%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			fmt.Printf("%s 关闭了\n", address)
		} else {
			fmt.Printf("%s 打开了\n", address)
			conn.Close()
		}
		wg.Done()
	}
}

func main() {
	var wg sync.WaitGroup

	ports := make(chan int, 100)

	for i := 0; i < cap(ports); i++ { //相当于雇佣了100个工人
		go worker(ports, &wg)
	}

	for i := 1; i < 1024; i++ { //这里一共1024个任务
		wg.Add(1)
		ports <- i
	}
	wg.Wait()
	close(ports)
}
