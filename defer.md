# defer 语句

延迟处理
遇到defer 其后的语句将会被压栈 待，函数返回时再一次弹栈执行
所以，后defer的先执行
一般用于数据库连接关闭、文件关闭等操作

```go
fmt.Printf("step1\n")
defer fmt.Printf("step2\n")
defer fmt.Printf("step3\n")
defer fmt.Printf("step4\n")
fmt.Printf("step5\n")

step1
step5
step4
step3
step2
```

```go
	fmt.Printf("f1 : step1\n")
	defer fmt.Printf("f1 : step2\n")
	defer fmt.Printf("f1 : step3\n")
	fmt.Printf("f1 : step5\n")
}

func f2() {
	fmt.Printf("f2 : step1\n")
	defer fmt.Printf("f2 : step2\n")
	defer fmt.Printf("f2 : step3\n")
	fmt.Printf("f2 : step5\n")
}
f1()
f2()

f1 : step1
f1 : step5
f1 : step3
f1 : step2
f2 : step1
f2 : step5
f2 : step3
f2 : step2
```

