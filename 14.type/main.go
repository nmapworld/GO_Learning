package main

// golang 类型定义

import "fmt"

type myInt int    //自定义类型
type youInt = int //类型别用

// 区别：
// 自定义类型为：包名.类型名
// 别名类型为原本的类型

func main() {
	var n myInt
	n = 100
	fmt.Printf("%d的类型为%T\n", n, n) // 100的类型为main.myInt
	var m youInt
	m = 1000
	fmt.Printf("%d的类型为%T\n", m, m) // 1000的类型为int
	var c rune
	c = '名'
	fmt.Printf("%c的类型为%T\n", c, c) // 名的类型为int32
}
