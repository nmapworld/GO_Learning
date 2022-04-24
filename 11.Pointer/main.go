package main

import "fmt"

func main() {
	// 指针
	// go语言中不存在指针操作
	// &：取地址
	// *：根据地址取值

	str1 := "Hello world"
	// 查询内存地址(内存地址为十六进制)
	fmt.Println(&str1)
	// *string:string类型指针 *int:int类型指针
	fmt.Printf("内存地址10进制表示为: %d,数据类型为：%T\n", &str1, &str1)
	// 根据内存地址查询对应的值
	fmt.Println(*(&str1))

	fmt.Println()
	str2 := "北京欢迎你"
	str3 := str2
	str4 := &str2
	fmt.Printf("str2:内存地址为：%x;\tstr3:内存地址为：%x;\tstr4:值为：%s\n", &str2, &str3, *str4)

	runModify()

	new_pointer()
}

func modify1(x int) {
	// 定义值为100的int变量x
	x = 100
}

func modify2(x *int) {
	// 定义值为100的int（变量x：指针int）
	*x = 100
}

func runModify() {
	a := 10
	modify1(a)
	// a=10> >x=10
	fmt.Println(a)
	modify2(&a)
	// ???
	fmt.Println(a)

}
