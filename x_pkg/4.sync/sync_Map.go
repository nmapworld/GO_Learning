package main

import (
	"fmt"
	"strconv"
	"sync"
)

func syncMap() {
	var m1 = sync.Map{}       //声明支持goroutine的map，无需make初始化
	for i := 0; i < 21; i++ { //循环21次
		wg.Add(1)        //goroutine计数+1
		go func(n int) { //匿名函数
			key := strconv.Itoa(n)   //强制转换数据类型
			m1.Store(key, n)         //存储数据至map中
			value, _ := m1.Load(key) //读取map中value
			fmt.Printf("K=%v\t V:=%v\n", key, value)
			wg.Done() //goroutine计数-1
		}(i)
	}
	wg.Wait() //等待goroutine计数器归零

}
