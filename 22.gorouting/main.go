package main

import (
	"fmt"
	"time"
)

// goroutine实例
// GMP：goroutine\Machine\Process

func showTime(i int) {
	fmt.Println(time.Now().Local(), i)
}

func baseGoroutine() {
	for i := 0; i <= 1000; i++ {
		go showTime(i) //开启i 个goroutine执行函数
	}
	time.Sleep(time.Second)
}

func main() {
	// 程序启动后会创建一个主goroutine去执行
	// 并发
	baseGoroutine()
	// 使用sync.WaitGroup并发延迟
	waitGroup()

	// 多核并发
	goMaxProcs()

	// 并行通信
	chann()
	fmt.Println("main") //mian函数结束后

	channel2()
}
