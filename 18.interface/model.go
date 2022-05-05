package main

import "fmt"

// 接口实例1
// 无论什么品牌的车都会跑

// run 的接口
type car interface {
	run()
}

// 结构体
type falali struct {
	brand string
}

type baoshijie struct {
	brand string
}

// 方法
func (f falali) run() {
	fmt.Printf("%s跑的快\n", f.brand)
}
func (b baoshijie) run() {
	fmt.Printf("%s跑的更快\n", b.brand)
}

// 函数接受一个car 类型的变量
func drive(c car) {
	c.run()
}

func model() {
	// 初始化结构体
	var f1 = falali{
		brand: "法拉利",
	}
	var b1 = baoshijie{
		brand: "保时捷",
	}

	// 调用接口
	drive(f1)
	drive(b1)

}
