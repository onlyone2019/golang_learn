# 常用命令

## mod init
初始化项目 `go mod init project_name`

项目最开始时需要在命令行执行一下`init`命令  会生成`go.mod`文件


## build
生成一个可执行文件，如：`go build main.go` 会生成一个 `go.exe`

同样，自己构建的package有修改后，需要切换到包文件夹下执行 `go build`，编译到缓存再去`import`才不会报错。

## run
编译执行 ，`go run main.go`

## tidy

## env
显示当前环境 `go env`

一般需要设置这两个配置：(使用powershell)

`$env:GO111MODULE = "on"`

`$env:GOPROXY = "http://goproxy.cn"`

## get
获取包。 `go get packet_url`

