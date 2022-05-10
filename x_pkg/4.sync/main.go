package main

import (
	"fmt"
	"sync"
)

var x = 0

// 使用变量保存groutine的运行结果,容易造成主程序关闭时，goroutine运行并为完成
// 1.用通信方式共享内存,推荐使用通道 channel
// 2.使用全局变量 互斥锁配合 sync.Mutex
// 3.使用全局变量 原子操作配合 atomic.Addint64()

var wg sync.WaitGroup

// 对goroutine加锁、解锁
var lock sync.Mutex

func add() {
	for i := 0; i < 50000; i++ {
		lock.Lock() //加锁
		x = x + 1
		lock.Unlock() //解锁
	}
	wg.Done()
}

func main() {
	wg.Add(2)

	go add()
	go add()

	wg.Wait()
	fmt.Println(x)

	syncMap()

}
