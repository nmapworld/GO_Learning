package main

import "fmt"

func sliceMy() {
	// 切片
	// 切片为引用类型，底层指向到数组
	// 修改切片元素的值时，其对应的内存地址不变

	a := [...]int{55, 56, 57, 88, 99, 96, 100}
	// 左包右不包含,（左闭右开）
	b := a[2:6]
	// b := a[:]     a[0:6]  6=len(a)
	// b := a[:4]    a[0:4]
	// b := a[1:]  //a[1:6]
	fmt.Printf("%d %T\n", b, b)
	fmt.Println("打印切片数组长度及容量")
	// 切片的长度为元素个数，切片的长度从切片时的索引到底层元素的最后一个值
	fmt.Printf("len(b):%d cap(b)%d\n", len(b), cap(b))
	c := a[:4]
	fmt.Printf("len(c):%d cap(c)%d\n", len(c), cap(c))
	d := c[:3]
	fmt.Printf("len(d):%d cap(d)%d\n", len(d), cap(d))

	e1 := []int{1, 3, 5}
	e2 := e1
	e2[0] = 5
	fmt.Println(e1, e2) //[5 3 5] [5 3 5]

	var slice1 []int
	if len(slice1) == 0 {
		fmt.Println("slice1 is Null")
		slice1 = make([]int, 10)
		fmt.Printf("%#v %p \n", slice1, &slice1)
		// []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0} 0xc00000c108
	}
	fmt.Printf("%#v %p \n", slice1, &slice1)
	// []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0} 0xc00000c108

	slice2 := make([]bool, 6)
	if len(slice2) == 0 {
		fmt.Println("slice1 is Null")
		slice2 = make([]bool, 10)
		fmt.Printf("%#v %p \n", slice2, &slice2)
	}
	fmt.Printf("%#v %p \n", slice2, &slice2)
	// []bool{false, false, false, false, false, false} 0xc0000a4108
}
