package main

import (
	"fmt"
	"sort"
)

var arr = [5]int{15, 23, 8, 10, 7}

// 内置排序
func incloudSort() {
	arr0 := arr     //复制数组副本，防止修改原数组值
	arr1 := arr0[:] // 转化为切片
	sort.Ints(arr1)
	fmt.Printf("%#v\n", arr1)
	// 7       8       10      15      23
}

// 冒泡排序：依次比较两个相邻的元素，次数为:n-1
func bubbleSort() {
	arr2 := arr
	// fmt.Println(arr2)
	n := 0 //计数
	for i := 0; i < len(arr2); i++ {
		// fmt.Printf("%v:%v\t", i, arr2[i])
		for j := 0; j < len(arr2)-i-1; j++ {
			if arr2[j] > arr2[j+1] {
				arr2[j], arr2[j+1] = arr2[j+1], arr2[j]
			}
		}
		n++
		// fmt.Printf("%v %v\n", i, arr2)
	}
	// fmt.Println()
	fmt.Printf("共运算了%d次\n", n)
	fmt.Printf("%#v\n", arr2) //[5]int{7, 8, 10, 15, 23}
}

// 插入排序
func insertSort() {

}

func sortArray() {
	fmt.Printf("原数据为：%#v\n", arr)
	incloudSort()
	bubbleSort()
	insertSort()
}
