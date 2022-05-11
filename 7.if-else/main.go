package main

import (
	"fmt"
	"math/rand"
	"time"
)

// if-else
// for
// switch-goto

// 单个判断
func main() {
	// 单个判断
	num := 66
	fmt.Print("猜猜我的年龄：")
	fmt.Scanf("%d", &num)
	if num == 66 {
		fmt.Println("猜对了")
	} else {
		fmt.Println("猜错了")
	}

	// 多个判断
	var age int
	fmt.Print("请输入您的的年龄：")
	fmt.Scanf("%d", &age)

	// if-else
	if age >= 35 {
		fmt.Println("人到中年")
	} else if age >= 18 {
		fmt.Println("恭喜你，成年了")
	} else if age >= 6 {
		fmt.Println("好好学习")
	} else {
		fmt.Println("好好把握童年时光吧")
	}

	// switch-case
	rand.Seed(time.Now().UnixNano()) //添加随机数种子，否则程序build后，出现的值相同
	n := rand.Int63n(4)              //生成一个随机整数，0<=n<4
	switch n {
	case 1:
		fmt.Println("echo 1")
	case 2:
		fmt.Println("echo 2")
	case 3:
		fmt.Println("echo 3")
	case 4:
		fmt.Println("echo 4")
	default:
		fmt.Println(n)
	}

}
