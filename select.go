/*
select 语句

2022-11-12
*/

package main

import (
	"fmt"
	"time"
)

var chanInt = make(chan int)
var chanStr = make(chan string)

func write() {
	chanInt <- 9
	chanStr <- "hello"
}

func main() {
	go write()
	time.Sleep(time.Second)
	for i := 0; i < 5; i++ {
		select {
		case c := <-chanInt:
			// fmt.Printf("c: %v\n", c)
			fmt.Printf("chanInt: %v\n", c)
		case c := <-chanStr:
			fmt.Printf("chanStr: %s\n", c)
		default:
			fmt.Println("deault")
		}
	}
}
