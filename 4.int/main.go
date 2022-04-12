package main

import "fmt"

func main() {
	const numberN1 = 1024
	fmt.Printf("%T\n", numberN1)
	// 10进制
	fmt.Printf("%d\n", numberN1)
	// 二进制
	fmt.Printf("%b\n", numberN1)
	// 八进制
	fmt.Printf("%o\n", numberN1)
	// 十六 进制
	fmt.Printf("%x\n", numberN1)

	// 十六进制转十进制
	var rdpPort = 0xd3d
	fmt.Printf("%d\n", rdpPort)

	// 声明不同进制常量
	const (
		n1 = int(10)
		n2 = int8(077)
		n3 = int16(0xd3d)
	)
	fmt.Printf("n1:%T\tn1:%T\tn3:%T\n", n1, n2, n3)
	fmt.Printf("n1:%v\tn1:%v\tn3:%v\n", n1, n2, n3)

}
