package main

import (
	"fmt"
)

// channel练习
//
// 1.启动一个goroutine，生成100个数发送到ch1
// 2.启动一个goroutine，从ch1中取值，计算其平方后发送到ch2
// 3.在main中，从ch2中取值打印
// var wg sync.WaitGroup

// 函数生成100个数，并将其传送到通道,仅接存放值进通道
func chan1(x chan<- int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		x <- i //写入通道
	}
	close(x) //关闭通道
}

// 函数从通道遍历值，并将值的平方传送到新的通道
func chan2(x <-chan int, y chan<- int) {
	defer wg.Done()
	for n := range x {
		y <- n * n //写入通道
	}
	close(y) //关闭通道

}
func channel2() {
	// 声明并初始化通道
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)
	wg.Add(2)
	// 执行函数
	go chan1(ch1)
	go chan2(ch1, ch2)
	// 等待并发运行结束
	wg.Wait()
	// 遍历通道中保存的数据
	for value := range ch2 {
		fmt.Println(value)
	}

}
