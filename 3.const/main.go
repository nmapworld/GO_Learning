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
	b1 = iota //0
	b2        //1
	_         //丢弃赋值
	b3        //3
	b4        //4
)

// const中每新增一行，则计数一次
const (
	c1 = iota //0
	c2 = 100  //破坏累加队形
	c3 = iota //2
	c4        //3
)

// const多列声明
const (
	d1, d2 = iota + 1, iota + 2 //1,2
	d3, d4 = iota + 1, iota + 2 //2,3
)

// 定义存储单位
const (
	_  = iota             //0
	KB = 1 << (10 * iota) //1*2^(10*1)=1024
	MB                    //1*2^(10*2)=1048576
	GB                    //1*2^(10*3)=1073741824
	TB                    //1*2^(10*4)=1099511627776
	PB                    //1*2^(10*5)=1125899906842624
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
	// fmt.Printf("a1", &a2) //数字常量不会分配存储空间，因此无法获取内存地址
}
