package main

import "fmt"

func numList() {
	// i := 1
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i = i + 1
	// }
	// 等同于
	// i := 0
	// for i < 10 {
	// 	i++
	// 	fmt.Printf("%d\t", i)
	// }
	// fmt.Println("")

	fmt.Println("1-10")
	// for 初始语句；条件语句；结束语句
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
	fmt.Print("1-100累加和为:")
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
	fmt.Println("")
}

func forRange() {
	words1 := "HelloMM你好"
	for i, v := range words1 {
		fmt.Printf("%d\t%c", i, v)
	}
	fmt.Println("")
	for _, v := range words1 {
		fmt.Printf("%c", v)
	}

}

func chengfabiao() {
	fmt.Println("")
	for i := 1; i < 10; i++ {
		fmt.Println("")
		for y := 1; y < 10; y++ {
			if y <= i {
				fmt.Printf("%d*%d=%d\t", y, i, i*y)

			}
			// fmt.Println("")
		}
	}

}
