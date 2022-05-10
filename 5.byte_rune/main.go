package main

import "fmt"

// Byte和rune类型
// go语言中为了非处理ASCII码类型的字符，定义了新的rune类型

func main() {

	a := '2'
	b := '新'
	fmt.Printf("%T,%T\n", a, b)
	stringMy := "Hello안녕하세요.こんにちは"
	stringMy2 := "Helloword"
	n := len(stringMy)
	fmt.Printf("%s\t", stringMy)
	fmt.Println(n)

	// 遍历字符串
	// byte遍历
	for i := 0; i < len(stringMy2); i++ {
		fmt.Printf("(%c)字符的十进制ASCII码为：%v\n", stringMy2[i], stringMy2[i])
	}
	fmt.Println("")
	// rune遍历
	for _, c := range stringMy {
		fmt.Printf("%T\t", c)
		fmt.Printf("%c\t", c)
		fmt.Println()
	}
}
