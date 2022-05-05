package main

import "fmt"

// 结构体中模拟继承关系

// 动物结构体
type animal struct {
	name string
}

// animal 的方法
func (a animal) move() {
	fmt.Printf("%s move move move\n", a.name)
}

type dog struct {
	feet uint8
	animal
}

// dog voice方法
func (d dog) wang() {
	fmt.Printf("%s 汪汪汪\n", d.name)

}

func jicheng() {
	d1 := dog{
		animal: animal{"周林"},
		feet:   4,
	}
	fmt.Println(d1)
	d1.move()
	d1.wang()

}
