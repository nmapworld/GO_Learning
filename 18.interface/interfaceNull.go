package main

import "fmt"

// interface{}  空接口,任意类型的变量均可传递
// 若空接口中为实现任何方法，则任意类型实现了该接口，切任何类型的数据均可保存在空接口中

// interface :关键字
// interface{}:空接口

func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}

func interfaceNull() {
	// var m1 map[string]interface{}
	m1 := make(map[string]interface{}, 16)
	fmt.Printf("%T\n", m1)
	m1["name"] = "chou"
	m1["age"] = 18
	m1["merried"] = true
	m1["hobby"] = [3]string{"sing", "jump", "rap"}
	fmt.Println(m1)

	show(false)
	show(nil)
	show(m1)

}
