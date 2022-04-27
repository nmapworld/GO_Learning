package main

import "fmt"

func defer1() {
	// defer语句在函数中最后被执行，当存在多个defer时，最后一个defer最先执行
	// defer多用于释放资源、文件关闭、解锁、记录时间
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

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func Que() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1

	// 1.a=1 && b=2
	// 2.defer calc("1", a, calc("10", a, b)) -->defer calc("1",1,calc("10",1,2)) -->defer calc("1",1,3)
	// 3.a=0
	// 4.defer calc("2", a, calc("20", a, b)) -->defer calc("2",0,calc("20,0,2)) -->defer calc("2",0,2)
}
