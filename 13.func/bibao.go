package main

import (
	"fmt"
	"strings"
)

// 闭包=函数+引用环境

// 原理：
// 1.函数可以作为返回值
// 2.函数内部查找变量，内部未找到时，外层寻找

// 定义函数adder，参数int，返回值为func类型
func adder(x int) func(int) int {
	// 定义匿名函数，参数为int，返回值int
	return func(y int) int {
		x += y
		return x

	}
}

func main2() {
	ret := adder(100) //100为adder的参数

	ret2 := ret(200) //200为函数体中匿名函数的参数
	fmt.Println(ret2)

	// 判断文件后缀
	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")
	fmt.Println(jpgFunc("test"))
	fmt.Println(jpgFunc("呵呵.jpg"))
	fmt.Println(txtFunc("test"))
	fmt.Println(txtFunc("哈哈.txt"))

}

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}

}
