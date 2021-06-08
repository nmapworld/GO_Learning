package main

import "fmt"

func numList() {
	// i := 1
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i = i + 1
	// }
	// 等同于
	fmt.Println("1-10")
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d\t", i)
	}
	fmt.Println("")
}

func jishu() {
	fmt.Println("奇数列表")
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			fmt.Printf("%d\t", i)
		}
	}
	fmt.Println("")
}

func oushu() {
	fmt.Println("偶数列表")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Printf("%d\t", i)
		}
	}
	fmt.Println("")
}
func gustNum() {
	fmt.Println("猜数字游戏开始")
	var num int
	for {
		fmt.Print("请输入数字:\t")
		fmt.Scanln(&num)
		if num == 66 {
			fmt.Println("You are right")
			break
		} else if num > 66 {
			fmt.Println("大了呀")
		} else {
			fmt.Println("猜小了呀")
		}
	}
}

func sumNum() {
	fmt.Print("1-100累加和为：")
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func daoxu() {
	fmt.Print("\n10-1倒序\n")
	for i := 10; i > 0; i-- {
		fmt.Printf("%d\t", i)
	}
}
