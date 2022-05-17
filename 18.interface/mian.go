package main

import "fmt"

// 引出接口实例
// 接口是一种类型
// 忽略传入数据类型，只调用
// 方法接受者可以分为值、指针类型

// speak的接口
type speaker interface {
	speak() //只要实现了speak方法的变量都是speaker类型,方法签名
}

type cat struct {
	anmail
}

type dog struct {
	anmail
}

type person struct{}

// 类型嵌套
type anmail struct {
	name string
	age  uint8
}

// 方法
func (c cat) speak() {
	fmt.Println("喵喵喵")

}

func (d dog) speak() {
	fmt.Println("汪汪汪")

}

func (p person) speak() {
	fmt.Println("啊啊啊")
}

func pick(x speaker) {
	// 接受参数，传进来什么打什么
	x.speak() //挨打了，会叫

}

func main() {
	p := anmail{
		name: "叨叨",
		age:  8,
	}
	fmt.Println(p.name, p.age)
	var c1 cat
	var d1 dog
	var p1 person

	pick(c1)
	pick(d1)
	pick(p1)

	model()

	detail()

	interfaceNull()

	type_duanyan()
	// interfaceDuo()

}
