/*
扫描TCP端口
*/

package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

var p sync.WaitGroup

func main() {
	start := time.Now()
	for i := 1; i < 600; i++ {
		p.Add(1)
		go func(j int) {
			defer p.Done()
			address := fmt.Sprintf("192.168.0.168:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				fmt.Printf("%s 关闭了！\n", address)
				return
			}
			conn.Close()
			fmt.Printf("%s 打开了！！！\n", address)
		}(i)
	}
	p.Wait()
	elapsed := time.Since(start) / 1e9
	fmt.Printf("elapsed: %d second\n", elapsed)
}
