/*
超小项目二：
命令行读用户输入 然后输出

小生不才
2022-11-17
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("what is your name?")
	reader := bufio.NewReader(os.Stdin) //从标准输入中读
	text, _ := reader.ReadString('\n')  //以回车结束输入 并且回车会读进text
	fmt.Printf("your name is: %v\n", text)
}
