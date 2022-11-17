/*
ticker

2022-11-17 小生不才
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)

	// for _ = range ticker.C {
	// 	fmt.Printf("time.Now(): %v\n", time.Now())
	// 	fmt.Println("ticker ...")
	// }

	tickerChannel := make(chan int)

	go func() {
		for _ = range ticker.C {
			select {
			case tickerChannel <- 1:
			case tickerChannel <- 2:
			case tickerChannel <- 3:
			}
		}
	}()

	sum := 0
	for v := range tickerChannel {
		fmt.Printf("收到: %v\n", v)
		sum += v
		if sum >= 15 {
			ticker.Stop()
			break
		}
	}

}
