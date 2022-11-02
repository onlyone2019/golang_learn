# switch 语句

## 标准写法 
不用显式加break 会自动break，default也可以不要
```go
a := 100
switch a {
case 100:
    fmt.Printf("a: %v\n", a)
case 200:
    fmt.Printf("a: %v\n", a)
default:
    fmt.Printf("other")
}
```

## case 里面可以加表达式
switch 后面什么条件也不写 默认为真
```go
a := 400
switch {
case a <= 100:
    fmt.Printf("a: %v\n", a)
case a > 100:
    fmt.Printf("a: %v\n", a)
default:
    fmt.Printf("other")
}
```

## case里可以有多个值
```go
a := 4
switch a {
case 1, 2, 3:
    fmt.Printf("a: %v\n", a)
case 4:
    fmt.Printf("a: %v\n", a)
default:
    fmt.Printf("other")
}
```

## 使用 fullthrough 关键字执行多个case
一个fullthrough只会保证下一个case也执行，如果下下一个也需要执行则还需要加
```go
a := 2
switch a {
case 1, 2, 3:
    fmt.Printf("1,2,3 a: %v\n", a)
    fallthrough
case 4:
    fmt.Printf("4 a: %v\n", a)
case 5:
    fmt.Printf("5 a: %v\n", a)
default:
    fmt.Printf("other")
}
```

[本节示例](https://github.com/onlyone2019/golang_learn/blob/master/switch.go)