## go build
> go build main.go
> go build -o main day01
## go run
> go run main.go

## go help
> go help build 
> go version
> go env

## 交叉编译
* linux或者mac交叉编译
```CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go```
```CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go```
```CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go```
* window上交叉编译
```SET CGO_ENABLED=0 SET GOOS=darwin SET GOARCH=amd64 go build main.go```
```SET CGO_ENABLED=0 SET GOOS=linux SET GOARCH=amd64 go build main.go```
* 交叉编译参数含义
  * CGO_ENABLED 是否使用cgo编译，0为不使用，1为使用，使用cgo进行交叉编译时需
要编译机器安装对应的cgo程序
  * GOOS 目标操作系统标识，windows对应Windows操作系统exe可执行文件，darwin
对应Mac可执行文件，linux对应Linux可执行文件，freebsd对应UNIX系统
  * GOARCH 目标可执行程序操作系统构架，包括 386，amd64，arm

## 帮助文档
> go doc strconv
> go doc strconv.Itoa

## 测试文档
> go test 文件名_test.go
> 函数内容
```go
package main
import (
    "testing"
    "time"
)
func TestHelloWorld(t *testing.T) {
    timestamp := time.Now().Unix()
    t.Log(timestamp)
}
```