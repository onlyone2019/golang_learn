# init函数

init函数优先于主函数执行
每个包 每个文件中可以有多个init 但执行顺序不确定 init间最好不要有依赖
init函数不能带参数和返回值

```go
func init() {
	fmt.Print("init...\n")
}

func main() {
	fmt.Print("main...\n")
}

init...
main...
```