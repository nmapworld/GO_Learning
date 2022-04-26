package main

import "fmt"

// Format：
// func 函数名（参数 参数类型）（返回值 返回值类型）{
// 函数体
// }

// 函数为一段代码的封装，方便重复调用，查看时更简洁清晰
// 函数是一种类型，可以作为参数，也可以作为返回值

// 定义加法函数
func sum(x int, y int) (result int) {
	return x + y
}

// 有参数,没有返回值的函数
// 参数类型相同时，可以简写为
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

//多个返回值
func f7(x, y int) (z, w int) {
	fmt.Printf("x=%d,y=%d\n", x, y)
	z = x + y
	w = x - y
	return
}

// GO语言中没有默认参数概念

func main() {
	r := sum(1, 2)
	fmt.Println(r)

	_, n := f5()
	fmt.Println(n)

	f6("下雨了", 5, 5, 5, 55)
	// fmt.Println()
	fmt.Println(f7(9, 5))

	defer1()
	fmt.Println(defer2(5))
	fmt.Println(defer3())
	fmt.Println(defer4())
	fmt.Println(defer5())

	fmt.Printf("%T\n", advF1)
	fmt.Printf("%T\n", advF2)

	advF3(advF2)

	advF4(advF2)

}
