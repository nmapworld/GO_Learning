package main

import "fmt"

// 同一个结构体中多个接口

// 接口嵌套
type animail interface {
	mover
	eater
}

// 定义接口
type mover interface {
	move()
}

type eater interface {
	eat(string)
}

// 定义cat结构体
type cat3 struct {
	name string
	feet uint8
}

// cat的方法
func (c *cat) move() {
	fmt.Println("走猫步")

}

func (c *cat) eat(food string) {
	fmt.Printf("猫猫吃%s\n", food)

}

func mover1(x mover) {
	x.move()

}

func interfaceDuo() {
	var c1 animail
	c1.eat("虫子")
	c1.move()
}
