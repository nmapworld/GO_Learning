package main

import (
	"fmt"
)

// 类型断言，接口数据类型的判断实现

func assign(a interface{}) {
	fmt.Println("if-else")
	fmt.Printf("%T\n", a)
	str, ok := a.(string) //判断是否为字符串类型
	if !ok {
		fmt.Printf("%v非字符串类型数据\n", str)
	} else {
		fmt.Printf("%v:为字符串\n", str)

	}
}

func assign2(a interface{}) {
	fmt.Println("switch-case")
	switch v := a.(type) {
	case int:
		fmt.Printf("%v: It's int\n", v)
	case string:
		fmt.Printf("%v: It's string\n", v)
	case bool:
		fmt.Printf("%v: It's bool\n", v)
	default:
		fmt.Printf("%v: Unknow\n", v)
	}

}

func type_duanyan() {
	assign(100)
	assign2(`c`)

}
