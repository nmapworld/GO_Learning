#  GO command
## 一. command

### 1. run

编译并运行test.go，但本地不会生成可执行文件

```go
go run test.go
```

![](0.png)

### 2. build

编译生成可执行文件

```go
go build test.go
```

![](1.png)

### 3. clean

清除当前目录下的代码产生的对象文件，仅保留源代码

```go
go clean
```

![](3.png)

### 4. env

查看当前go环境的配置

```go
go env
```

![](4.png)

### 5. fmt

运行gofmt对代码进行格式化处理

```go
go fmt test.go
```

![](5.png)

### 6. get

下载并安装包和依赖

```go
go get https://github.com/lxn/walk
```

### 6. list

列出的包

```go
go list
```

![](6.png)

## 二、set

### 1. China source

```bash
go env -w GOPROXY=https://goproxy.cn,direct
```

### 2. Enable modules

```bash
go env -w GO111MODULE=on
```

### 3. Golang Env

```bash
go env -w set GOROOT=D:\Program Files\Go
go env -w set GOPATH=C:\Users\Administrator\go
```

