# waitGroup 实现协程同步

正常情况下，主程序中创建的线程可能还没执行完，主程序就退出了。

可以使用waitGroup 等所有协程执行完了 主协程才退出。

其基本处理逻辑是，创建一个协程则将组协程计数器增1，每执行完一个协程则将计数器减1，主协程前使用wait()等待所有协程执行完毕在往下执行。

[本节示例](https://github.com/onlyone2019/golang_learn/blob/master/waitGroup.go)