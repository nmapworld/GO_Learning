package main

import "fmt"

// 函数格式
// func 函数名（参数）（返回值）{
// 函数体
// }
// 一段代码的封装，方便重复调用，查看时更简洁清晰

// 定义加法函数
func sum(x int, y int) (result int) {
	return x + y
}

// 有参数,没有返回值的函数
// 参数相同时，可以简写为
// func  f1(x,y int)
// 同理f1(x,y int,m n string,i,o bool)
func f1(x int, y int) {
	fmt.Println(x + y)

}

// 没有参数、有返回值的
func f2() int {
	return 3
}

// 没有参数没有返回值的
func f3() {
	fmt.Println("Hello word")
}

// 参数可以命名也可以不命名
// 命名参数的返回值相当于声明了个变量
func f4(x int, y int) (ret int) {
	ret = x + y
	return
	// 相当于return ret

}

// 多个返回值
func f5() (int, string) {
	return 1, "Hello"
}

// 可变长度的参数，只能放在最后参数

func f6(x string, y ...int) {
	fmt.Println(x)
	fmt.Println(y) //返回的类型的slice []int

}

// GO语言中没有默认参数概念

func main() {
	r := sum(1, 2)
	fmt.Println(r)

	_, n := f5()
	fmt.Println(n)

	f6("下雨了", 5, 5, 5, 55)
	// fmt.Println()

}
