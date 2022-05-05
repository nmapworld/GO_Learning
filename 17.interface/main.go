package main

import "fmt"

// 接口是一种特殊的类型
// 当忽略传入参数的类型时，只管调用对应的方法。使用接口来实现
// 可以作为变量 、参数、返回值等设置类型

// fmt
// type 接口名 interface{
// 	方法名1 (参数1,参数2...)(返回值1,返回值2...)
// 	方法名2 (参数1,参数2...)(返回值1,返回值2...)
// 	...
// }

type voiceTest interface {
	// 掉用不同方法下的
	voice()
}

// 定义结构体
type cat struct {
}

type dog struct {
}

// 定义方法
func (c cat) voice() {
	fmt.Println("喵喵喵")
}

func (d dog) voice() {
	fmt.Println("汪汪汪")
}

// 不同方法，调用接口
func pick(x voiceTest) {
	// 接受一个参数，传入动物
	x.voice()
}

func main() {
	var c1 cat
	var d1 dog

	pick(c1)
	pick(d1)

	// model()

}
