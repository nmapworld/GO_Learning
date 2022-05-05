package main

import "fmt"

// 方法
// 方法是作用特定类型的函数
// format
// func (接受者变量 接受者类型) 方法名(参数列表) (返回参数){
// 	函数体
// }

// 标识符：变量名、函数名、类型名、方法名
// GO语言中如果标识符首字母大写，则对外部包可见(公有的)

// dog 类型的结构体
type dog struct {
	name string
}

// 构造函数，解决结构体重复调用
func newdog(name string) dog {
	return dog{
		name: name,
	}
}

//接收者表示的是调用该方法的具体类型，多用结构体首字符小写表示
func (d dog) voice() {
	fmt.Printf("%s:汪汪汪", d.name)
}

func main() {
	d1 := newdog("周某某")
	d1.voice()
}
