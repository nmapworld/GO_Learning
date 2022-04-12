package main

import "fmt"

const pi = 3.141526

// 批量声明常量
const (
	statusOk = 200
	notFound = 404
)

// 若批量声明常量，仅第一个赋值，则其与皆赋相同的值
const (
	n1 = 100
	n2
	n3
)

// iota枚举使用
const (
	a1 = iota //0
	a2        //1
	a3        //2
	a4        //3
)

const (
	b1 = iota
	b2
	_ //丢弃赋值
	b3
	b4
)

// const中每新增一行，则计数一次
const (
	c1 = iota
	c2 = 100 //破坏累加队形
	c3 = iota
	c4
)

// const多列声明
const (
	d1, d2 = iota + 1, iota + 2
	d3, d4 = iota + 1, iota + 2
)

// 定义存储单位
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
)

func main() {
	fmt.Println(n1, n2, n3)
	fmt.Printf("Web 正常时状态码为:%d\t异常时为:%d\n", statusOk, notFound)
	fmt.Printf("pi的值为:%f\n", pi)
	fmt.Println(a1, a2, a3, a4)
	fmt.Println(b1, b2, b3, b4)
	fmt.Println(c1, c2, c3, c4)
	fmt.Println(d1, d2, d3, d4)
	fmt.Println(KB, MB, GB, TB, PB)
}
