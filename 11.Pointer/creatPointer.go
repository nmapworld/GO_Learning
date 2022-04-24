package main

import "fmt"

func new_pointer() {
	// 创建指针
	// 空指针，nil
	// 1.new函数创建内存地址(对象为int; string...),返回为对应类型的指针
	var a *int
	fmt.Println(a)
	// 合规方式,先申请内存
	var a1 = new(int)
	fmt.Println(*a1)
	*a1 = 100
	fmt.Println(a1)

	// 2.make函数创建内存地址(对象为：slice\map\chan)，返回值为类型本身

}
