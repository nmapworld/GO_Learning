#  GO command
## 1. run

编译并运行test.go，但本地不会生成可执行文件

```go
go run test.go
```

![image-20210610161210700](D:\GO\Learning\GO_Learning\command\0.png)

## 2. build

编译生成可执行文件

```go
go build test.go
```

![image-20210610161252187](D:\GO\Learning\GO_Learning\command\1.png)

## 3. clean

清除当前目录下的代码产生的对象文件，仅保留源代码

```go
go clean
```

![image-20210610161334583](D:\GO\Learning\GO_Learning\command\3.png)

## 4. env

查看当前go环境的配置

```go
go env
```

![image-20210610161745052](D:\GO\Learning\GO_Learning\command\4.png)

## 5. fmt

运行gofmt对代码进行格式化处理

```go
go fmt test.go
```

![image-20210610161554944](D:\GO\Learning\GO_Learning\command\5.png)

## 6. get

下载并安装包和依赖

```go
go get https://github.com/lxn/walk
```

## 6. list

列出的包

```go
go list
```

![image-20210610161919061](D:\GO\Learning\GO_Learning\command\6.png)

