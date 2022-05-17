package main

import (
	"fmt"
	"math"
)

// 方法
// 方法是作用特定类型的函数
// format
// func (接受者变量 接受者类型) 方法名(参数列表) (返回参数){
// 	函数体
// }

// 标识符：变量名、函数名、类型名、方法名
// GO语言中如果标识符首字母大写，则对外部包可见(公有的)

// dog 类型的结构体
type dog struct {
	name string
}

// 构造函数，解决结构体重复调用
func newdog(name string) dog {
	return dog{
		name: name,
	}
}

//接收者表示的是调用该方法的具体类型，多用结构体首字符小写表示
func (d dog) voice() {
	fmt.Printf("%s:汪汪汪", d.name)
}

// 矩形结构体
type rectangle struct {
	width, height float64
}

// 圆形结构体
type circle struct {
	radius float64
}

// 矩形面积计算方法
func (r rectangle) area() float64 {
	return r.width * r.height
}

// 圆形面积计算方法
func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

type person struct {
	height int
	age    int
}

func (p person) addList() map[string]int {
	m1 := map[string]int{
		"height": 2,
		"age":    18,
	}
	return m1
}

// A nil *IntList represents the empty list.
type intlist struct {
	value int
	tail  *intlist
}

func (list *intlist) sum() int {
	if list == nil {
		return 0
	}
	return list.value + list.tail.sum()
}

func main() {
	d1 := newdog("周某某")
	d1.voice()

	// 面积计算
	r1 := rectangle{12, 2}
	r2 := rectangle{9, 4}
	c1 := circle{10}
	c2 := circle{25}
	fmt.Printf("%v *%v area :%v\n", r1.width, r1.height, r1.area())
	fmt.Printf("%v *%v area :%v\n", r2.width, r2.height, r2.area())
	fmt.Printf("%v^2 * Pi area :%v\n", c1.radius, c1.area())
	fmt.Printf("%v^2 * Pi area :%v\n", c2.radius, c2.area())

	p1 := person{
		age:    18,
		height: 2,
	}
	fmt.Println(p1.addList())
	var num intlist
	fmt.Println(num)

}
