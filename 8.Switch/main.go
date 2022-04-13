package main

import (
	"fmt"
)

// siwtch简化if大量的判断()
// fallthrough为兼容C语言case设计，执行满足条件case和他的下一个case

func main() {
	for1()
	switch1()
	switch2()
	switch3()
}

func for1() {
	var n int
	fmt.Println("请输入1-5的数字")
	fmt.Scanf("%d", &n)
	if n == 1 {
		fmt.Println("大拇指")
	} else if n == 2 {
		fmt.Println("食指")
	} else if n == 3 {
		fmt.Println("中指")
	} else if n == 4 {
		fmt.Println("无名指")
	} else if n == 5 {
		println("小拇指")

	} else {
		fmt.Println("请输入1-5的数字")
	}

}
func switch1() {
	var n int
	fmt.Println("请输入1-5的数字")
	fmt.Scanf("%d", &n)
	// switch n:=1;n
	switch n {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("输入错误，请重新输入")

	}
}

func switch2() {
	switch n := 1; n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	default:
		fmt.Println("n")

	}

}

func switch3() {
	var age int
	fmt.Print("请输入您的年龄：")
	fmt.Scanf("%d", &age)
	switch {
	case age <= 22:
		fmt.Println("好好学习吧")
	case age <= 40:
		fmt.Println("好好工作吧")
	case age <= 60:
		fmt.Println("活着真好")
	default:
		fmt.Println("好好享受")
	}
}
