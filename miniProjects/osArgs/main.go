/*
做一个简单地命令行参数显示工具

第一个golang小项目
小生不才
2022-11-17
*/

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	/*
		//方法一
		var s, sep string
		for i := 1; i < len(os.Args); i++ {
			s += sep + os.Args[i]
			sep = " "
		}
		fmt.Printf("%v\n", s)
	*/

	/*
		//方法二
		var s, sep string
		for _, v := range os.Args[1:] {
			s += sep + v
			sep = " "
		}
		fmt.Printf("%v\n", s)
	*/

	//方法三
	fmt.Println(strings.Join(os.Args[1:], " "))
}
