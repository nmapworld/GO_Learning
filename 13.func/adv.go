package main

import "fmt"

// 函数的返回类型:
// 无返回值的函数
func advF1() {
	fmt.Println("Hello world")
}

// 有返回值的函数
func advF2() int {
	return 10
}

// 函数作为参数的函数
func advF3(x func() int) {
	ret := x()
	fmt.Println(ret)
}

func ff(a, b int) int {
	a = 1
	b = 2
	c := a + b
	return c

}

// 函数作为返回值
// 参数为返回int类型函数，返回值为（int类型参数的函数，且返回值为int）
func advF4(x func() int) func(int, int) int {
	return ff
}
