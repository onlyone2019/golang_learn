# Ticker 周期执行

创建一个ticker ： `ticker := time.NewTicker(time.Second)`

下面这段代码将会隔1秒执行一次
```golang
for _ = range ticker.C {
    fmt.Printf("%v\n", time.Now())
    fmt.Println("ticker ...")
}
```

[本节样例](https://github.com/onlyone2019/golang_learn/blob/master/ticker.md)