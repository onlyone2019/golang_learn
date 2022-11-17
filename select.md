# select 

select是golang中的一个控制结构，类似switch，用于处理异步I/O操作。

select会监听case语句中的channel操作，当case中的读写操作为非阻塞时，才会触发相应的动作。所以，select中的case语句必须是一个channel操作

如果有多个case可以执行，select会公平地选取一个执行，其他不会执行

如果没有可运行的case语句，且有default语句，那么就会执行default语句

如果没有可运行的case语句，且没有default语句，那么select将会阻塞，直到某个case通信可以执行。

[本节样例](https://github.com/onlyone2019/golang_learn/blob/master/select.md)