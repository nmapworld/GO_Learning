package main

import "fmt"

// close             关闭channel
// len               求长度：string | array |slice |map |channel
// new               分配内存，值类型：   int ｜struct 返回值是指针
// make              分配内存，引用类型： channel |map | slice
// append            追加元素： array | slice
// panic && recover  错误处理

func funcA() {
	fmt.Println("A")
}
func funcB() {
	// 定义链接数据库失败时，程序终止
	// recover()和painc联合使用，且需在painc前定义
	// recover()函数使用时，需配合defer函数
	defer func() {
		err := recover()
		fmt.Println(err)
		fmt.Println("数据库连接失败")
	}()
	panic("出现严重错误，程序退出") //执行到此终止运行
	fmt.Println("B")

}
func funcC() {
	fmt.Println("C")
}

func main3() {
	funcA()
	funcB()
	funcC()
}
