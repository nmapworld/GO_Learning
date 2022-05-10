package main

import (
	"fmt"
	"runtime"
)

func a() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("A:%d\n", i)
	}
}

func b() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("B:%d\n", i)
	}
}

func goMaxProcs() {

	// runtime.GOMAXPROCS(12)  默认值为CPU线程数
	wg.Add(2)
	go a()
	go b()
	wg.Wait()
	n := runtime.NumCPU()
	fmt.Println(n)

}
