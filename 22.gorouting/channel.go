package main

import "fmt"

var a1 []int
var b1 chan int //指定通道中元素类型

// 通过channel实现多个gorotine相互通信

// 需要make初始化的变量类型为：slice、map、chan

func channOp() {
	// 发送 ch1 <-  1
	b1 <- 10
	// 接收 <-ch2
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-b1
		fmt.Printf("通道接收成功,x=%v\n", x)
	}()
	wg.Wait()
	// 关闭 close()
}

func chann() {
	fmt.Println(a1, b1)
	b1 = make(chan int) //初始化chan,无缓冲区时，会阻塞等待
	channOp()

	b1 = make(chan int, 16) //带缓冲区，初始化chan，当缓冲区满时，会阻塞等待

}
