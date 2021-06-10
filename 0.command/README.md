#  GO command
## 1. run

编译并运行test.go，但本地不会生成可执行文件

```go
go run test.go
```

![image-20210610161210700](https://github.com/nmapworld/GO_Learning/blob/main/0.command/0.png?raw=true)

## 2. build

编译生成可执行文件

```go
go build test.go
```

![image-20210610161252187](https://github.com/nmapworld/GO_Learning/blob/main/0.command/1.png?raw=true)

## 3. clean

清除当前目录下的代码产生的对象文件，仅保留源代码

```go
go clean
```

![image-20210610161334583](https://github.com/nmapworld/GO_Learning/blob/main/0.command/3.png?raw=true)

## 4. env

查看当前go环境的配置

```go
go env
```

![image-20210610161745052](https://github.com/nmapworld/GO_Learning/blob/main/0.command/4.png?raw=true)

## 5. fmt

运行gofmt对代码进行格式化处理

```go
go fmt test.go
```

![image-20210610161554944](https://github.com/nmapworld/GO_Learning/blob/main/0.command/5.png?raw=true)

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

![image-20210610161919061](https://github.com/nmapworld/GO_Learning/blob/main/0.command/6.png?raw=true)

