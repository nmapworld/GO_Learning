package main

import "fmt"

// 结构体声明
// 结构体为值类型，是类的一种
type person struct {
	name  string
	age   int
	hobby []string
}

// 构造函数1
type person4 struct {
	name string
	age  int
}

// 构造函数2，建议函数名为：new+结构体名
// 返回的是结构体或结构体指针
// 当结构体本身比较大时，应尽量使用结构体指针，减少程序开销
func newPerson(name string, age int) *person4 {
	return &person4{
		name: name,
		age:  age,
	}
}

func main() {
	// 实例化结构体
	var p person
	p.name = "周杰伦"
	p.age = 18
	p.hobby = []string{"唱", "跳", "rap"}
	fmt.Println(p) //类型为main.person
	fmt.Printf("%T\n", p)
	fmt.Println(p.name)

	// 匿名结构体：多用于临时场景，存在于函数体中
	var s struct {
		name string
		age  int
	}
	s.name = "小明"
	s.age = 18
	fmt.Printf("%v的类型为 :%T\n", s, s)

	f()

	// 构造函数使用
	p3 := newPerson("元帅", 18)
	p4 := newPerson("周林", 20)
	fmt.Println(*p3, *p4)

	// 嵌套结构体
	structDuo()

	// 继承

	jicheng()

	// json
	jsonM()
}
