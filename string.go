package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	//1 字符串赋值 多行字符串用反引号括起来
	s1 := "hello"
	s2 := `
	line 1
	line 2
	line 3
	`
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s2: %v\n", s2)

	//2 字符串拼接
	s2 = "world"
	fmt.Printf("%v\n", s1+s2)

	fmt.Printf("strings.Join([]string{s1, s2}, \"+\"): %v\n", strings.Join([]string{s1, s2}, "+"))

	var buffer bytes.Buffer
	buffer.WriteString("11111")
	buffer.WriteString(",")
	buffer.WriteString("22222")
	fmt.Printf("%s\n", buffer.String())

	//3 字符串格式化输出
	tmp := fmt.Sprintf("%s,%s", s1, s2)
	fmt.Printf("%v\n", tmp)

	//4 字符串切片操作
	s1 = "Hello World"
	fmt.Printf("s1[3]: %c\n", s1[3])
	fmt.Printf("s1[3:6]: %v\n", s1[3:6])
	fmt.Printf("s1[3:]: %v\n", s1[3:])
	fmt.Printf("s1[:4]: %v\n", s1[:4])

	//5 字符串的一些函数
	fmt.Printf("len(s1): %v\n", len(s1))
	fmt.Printf("strings.Split(s1, \" \"): %v\n", strings.Split(s1, " "))           //按空格切分 输出一个数组
	fmt.Printf("strings.Contains(s1, \"llo\"): %v\n", strings.Contains(s1, "llo")) //是否包含某个字符 包含就是true
	fmt.Printf("strings.ToLower(s1): %v\n", strings.ToLower(s1))                   //变小写
	fmt.Printf("strings.ToUpper(s1): %v\n", strings.ToUpper(s1))
	fmt.Printf("strings.HasPrefix(s1, \"Hell\"): %v\n", strings.HasPrefix(s1, "Hell")) //判断前缀是不是这个字符串
	fmt.Printf("strings.HasPrefix(s1, \"llo\"): %v\n", strings.HasPrefix(s1, "llo"))
	fmt.Printf("strings.HasSuffix(\"World\"): %v\n", strings.HasSuffix(s1, "World")) //判断后缀是不是这个字符串
	fmt.Printf("strings.HasSuffix(s1, \"orl\"): %v\n", strings.HasSuffix(s1, "orl"))
	fmt.Printf("strings.HasSuffix(s1, \"orld\"): %v\n", strings.HasSuffix(s1, "orld"))
	fmt.Printf("strings.Index(s1, \"orl\"): %v\n", strings.Index(s1, "orl"))     //查找字符串位置
	fmt.Printf("strings.LastIndex(s1, \"l\"): %v\n", strings.LastIndex(s1, "l")) //查找最后出现的位置
}
