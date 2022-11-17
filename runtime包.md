# runtime包(并发编程)

## runtime.Gosched() 让出时间片

当前协程写了这句 `runtime.Gosched()` 后，会放弃安排给它的时间片

[示例](https://github.com/onlyone2019/golang_learn/blob/master/runtimeGosched.go)

## runtime.Goexit() 停止协程

当前协程写了这句 `runtime.Goexit()` 后，会停止这个协程

[示例](https://github.com/onlyone2019/golang_learn/blob/master/runtimeGoexit.go)

## runtime.GOMAXPOCS(n) 设置运行线程的CPU个数

如标题，设置运行线程的CPU个数

[示例](https://github.com/onlyone2019/golang_learn/blob/master/runtimeGOMAXPOCS.go)