package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup //线程同步，等待一些列任务执行完成后向下继续

func f1(i int) {
	defer wg.Done() //减少 等待的goroutine计数
	// time.Sleep(time.Second)
	fmt.Println(i)

}

func waitGroup() {
	for i := 0; i < 100; i++ {
		wg.Add(1) //添加 等待的goroutine计数
		go f1(i)

	}
	wg.Wait() //等待的goroutine计数为0
}