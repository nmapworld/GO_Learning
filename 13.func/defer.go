package main

import "fmt"

func defer1() {
	// defer语句在函数中最后被执行，当存在多个defer时，最后一个defer最先执行
	// defer多用于释放资源
	fmt.Println("defer延迟演示")
	fmt.Println("Weclome")

	defer fmt.Println(3)
	defer fmt.Println(2)
	defer fmt.Println(1)

	fmt.Println("END")

}

func defer2(x int) int {
	// Go语言中的return非原子操作，分两步
	// 1.赋值给返回值
	// 2.defer
	// 3.真正的RET
	// 匿名函数：在函数体中定义的函数，无函数名
	defer func() {
		x++ //修改的为x的值，非返回值
	}()
	return x
}

func defer3() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func defer4() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return 5
}

func defer5() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}
