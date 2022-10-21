# 常用命令

## mod init
初始化项目 `go mod init project_name`

## build
生成一个可执行文件，如：`go build main.go` 会生成一个 `go.exe`

## run
编译执行 ，`go run main.go`

## env
显示当前环境 `go env`

一般需要设置这两个配置：(使用powershell)

`$env:GO111MODULE = "on"`

`$env:GOPROXY = "http://goproxy.cn"`

## get
获取包。 `go get packet_url`

