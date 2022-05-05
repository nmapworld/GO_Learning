package main

import "fmt"

// 接口的实现,包含两种方法
type animal interface {
	move()
	eat(string)
}

// 定义类型
type cat1 struct {
	name string
	feet uint8
}

type chicken struct {
	feet uint8
}

// 定义接口中对应的方法
// 值接受者
func (c cat1) move() {
	fmt.Println("猫猫动动")
}
func (c cat1) eat(food string) {
	fmt.Printf("猫猫吃%s\n", food)
}

func (c chicken) move() {
	fmt.Println("鸡鸡动")
}

func (c chicken) eat(food string) {
	fmt.Printf("鸡鸡吃%s\n", food)
}

func detail() {
	var a1 animal //定义一个接口类型的变量

	bc := cat1{ //定义一个cat类型的变量bc
		name: "淘气",
		feet: 4,
	}
	kfc := chicken{
		feet: 2,
	}
	a1 = bc
	a1.eat("小黄鱼")
	fmt.Println(a1)
	a2 := kfc
	a2.eat("虫虫")
	fmt.Println(a2)

}
