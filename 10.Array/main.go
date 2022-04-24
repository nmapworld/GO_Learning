package main

import "fmt"

func main() {
	// 数组
	// 存放元素的容器，需指定容器类型和长度
	var a1 [3]bool
	var a2 [4]bool
	// 声明数组a1,长度为5,类型为bool
	// 声明数组a2,长度为4,类型为bool
	// 数组的长度是类型的一部分，长度不同不能进行运算
	fmt.Printf("a1:%T,a2:%T\n", a1, a2)
	fmt.Println(a1)
	fmt.Println(a2)
	// 数组初始化

	// 初始化.1
	a1 = [3]bool{true, true, true}
	fmt.Println(a1)

	// 初始化.2
	// 根据内容长度，自动设置数组长度
	a100 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(a100)

	// 初始化.3
	// a3 := [5]int{1, 2}
	// 根据索引位置进行初始化
	a3 := [5]int{0: 1, 1: 2}
	fmt.Println(a3)

	arrayList()
}

func arrayList() {

	a := [...]string{"北京", "上海", "深圳"}
	fmt.Printf("a:%T\n", a)
	fmt.Println(a)
	// 遍历数组1
	for i := 0; i < len(a); i++ {
		fmt.Printf("%s\t", a[i])
	}
	fmt.Println()
	// 遍历数组2 range
	for i, value := range a {
		fmt.Printf("%d:%s\t", i, value)

	}
	fmt.Println()
	// 多维数组
	// [[1 2][2 3][3 4]]
	// var a11 [3][2]int
	a11 := [3][2]int{
		{1, 2},
		{2, 3},
		{3, 4},
	}
	fmt.Println(a11)
	// 遍历数组3 多维
	fmt.Println("多维for遍历")
	for i := 0; i < len(a11); i++ {
		// fmt.Printf("%d\t", a11[i])
		for y := 0; y < len(a11[i]); y++ {
			fmt.Printf("%d\t", a11[i][y])
		}
	}
	fmt.Println()
	// 遍历数组4  多维range
	fmt.Println("多维range遍历")
	for _, vaule := range a11 {
		// fmt.Println(vaule)
		for _, vaule2 := range vaule {
			fmt.Printf("%d\t", vaule2)
		}
	}
	fmt.Println()
	// 值类型数组
	b1 := [5]int{1, 2, 3, 4, 5}
	b2 := b1
	b2[0] = 100
	fmt.Println(b1, b2)
	Question()
	make1()
	sliceMy()
	slice_append()
	slice_copy()
	slice_del()
	
}
