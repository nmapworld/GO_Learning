//嵌套结构体
package main

import "fmt"

type same struct {
	address string
	age     int
	name    string
}

type people struct {
	sex  string
	some same
}

type company struct {
	staff int
	some  same
}

func structDuo() {
	t1 := company{
		staff: 10,
		some: same{
			name:    "某某公司",
			address: "北京东路",
			age:     10,
		},
	}
	fmt.Println(t1)
	fmt.Println(t1.some.name, t1.some.address)
	p2 := people{
		sex: "man",
		some: same{
			name:    "小明",
			age:     18,
			address: "北京东路",
		},
	}
	fmt.Println(p2)

}
