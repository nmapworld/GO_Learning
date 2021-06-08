package main

import "fmt"

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
			continue
		}
		fmt.Printf("%d\t", i)
	}
}
