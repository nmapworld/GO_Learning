package main

import "fmt"

type person2 struct {
	name, nikename string
	age            int
}

func f2(x person2) {
	x.nikename = "小呆瓜"
}

func f() {
	fmt.Println("struct_test")
	var p2 person2
	p2.name = "周生辰"
	p2.nikename = "小南辰王"
	p2.age = 20
	fmt.Printf("%v的类型为: %T\n", p2, p2)
	// f2(person2)
}
