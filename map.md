# map

map 存的是 key:value 对，key有序
map[key] 返回两个值，第一个为value，第二个为key是否存在，如果key不存在，返回的value是空值，第二个是false

## map声明和分配内存
```go
var a map[string]int     //声明一个map
a = make(map[string]int) //分配内存
fmt.Printf("a: %v\n", a) //a: map[]
```

## map的初始化
```go
var b = map[string]int{
    "aaa": 1,
    "bbb": 2,
    "ccc": 3,
}
fmt.Printf("b: %v\n", b) //b: map[aaa:1 bbb:2 ccc:3]

var c map[string]int
c = make(map[string]int)
c["ee"] = 6
c["dd"] = 7
fmt.Printf("c: %v\n", c) //c: map[dd:7 ee:6]
```

## 判断键值是否存在
```go
key := "ee"
v, ok := c[key]
fmt.Printf("v: %v\n", v)                  //v: 6
fmt.Printf("ok: %v\n =========== \n", ok) //ok: true
key = "ff"
v, ok = c[key]
fmt.Printf("v: %v\n", v)                  //v: 0
fmt.Printf("ok: %v\n =========== \n", ok) //ok: false
```

## map遍历
```go
//遍历key
for k := range c {
    fmt.Printf("k: %v\n", k)
}

//遍历key和value
for k, val := range b {
    fmt.Printf("%v:%v\n", k, val)
}
```


[本节示例](https://github.com/onlyone2019/golang_learn/blob/master/map.go)
