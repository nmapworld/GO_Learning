package main

import "fmt"

// 结构体声明
// 结构体为值类型，是类的一种
type person struct {
	name  string
	age   int
	hobby []string
}

type books struct {
	title   string
	author  string
	subject string
	book_id int
}

// 结构体匿名字段
type nofiled struct {
	string
	int
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

	var book1 books
	book1.subject = "go语言教程"
	book1.author = "gopher"
	book1.title = "GO语言"
	book1.book_id = 1

	book2 := books{
		subject: "Python",
		author:  "Pythor",
		title:   "Python语言",
		book_id: 2,
	}
	fmt.Println(book1.book_id, book1.subject, book1.author, book1.title) //1 go语言教程 gopher GO语言
	fmt.Println(book2.book_id, book2.subject, book2.author, book2.title) //2 Python Pythor Python语言

	func(book1, book2 *books) {
		fmt.Println(book1.book_id, book1.subject, book1.author, book1.title) //1 go语言教程 gopher GO语言
		fmt.Println(book2.book_id, book2.subject, book2.author, book2.title) //2 Python Pythor Python语言
	}(&book1, &book2)

	filed1 := nofiled{
		"Hello",
		18,
	}
	fmt.Printf("%v %T\n", filed1, filed1)
}
