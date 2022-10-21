# struct

go没有面向对象的思想

## 基本用法
```go
type PPerson struct {
	id   int
	name string
}

var p *PPerson
p = &PPerson{}
p.id = 1
p.name = "111"
fmt.Printf("p: %p\n", p)  //p: 0xc000092060
fmt.Printf("p: %v\n", *p) //p: {1 111}

xx := PPerson{
    id:   2,
    name: "222",
}
fmt.Printf("xx: %v\n", xx) //xx: {2 222}

var yy = new(PPerson)
yy.id = 3
yy.name = "333"
fmt.Printf("yy: %p\n", yy) //yy: 0xc0000920a8
fmt.Printf("yy: %v\n", yy) //yy: &{3 333}
```

## 结构体可以嵌套
```go
type Dog struct {
	name string
}

type Ren struct {
	name string
	dog  Dog
}

r1 := Ren{}
r1.name = "r1"
r1.dog.name = "kate"
fmt.Printf("r1: %v\n", r1) //r1: {r1 {kate}}
```