package main

import "fmt"

// for 循环跳出方式
// continue & break & goto & return & panic
// continue &break &goto均后可跟标签

func paichu() {
	fmt.Println("")
	fmt.Println("1-10，排除7")
	fmt.Println("方法一")
	for i := 1; i < 11; i++ {
		if i != 7 {
			fmt.Printf("%d\t", i)
		}
	}
	fmt.Println("")
	fmt.Println("方法二")
	for i := 1; i < 11; i++ {
		if i == 7 {
			continue //继续下一次循环
		}
		fmt.Printf("%d\t", i)
	}
}

// 跳出多层for循环

func duoFor() {
	var flag bool
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
		for y := 'A'; y <= 'Z'; y++ {
			if y == 'E' {
				flag = true
				break //跳出内层for循环
			}
			fmt.Printf("%d%c\t", i, y)
			fmt.Println("")
		}
		if flag {
			break //跳出外层for循环
		}

	}
	fmt.Println("Over")

}
