package main

import (
	"fmt"
	"sort"
)

func Question() {
	//1.求数组中元素和
	// [1,3,5,7,8]
	var sum int
	array1 := [5]int{1, 3, 5, 7, 8}
	for _, value := range array1 {
		sum += value
	}
	fmt.Println(sum)

	// 2.找出数组中和为8的下标
	// [1,3,5,7,8]
	fmt.Println(array1)
	for index, value2 := range array1 {
		// fmt.Println(index, value2)
		// fmt.Println()
		for i := index + 1; i < len(array1); i++ {
			// fmt.Println(array1[i])
			if array1[i]+value2 == 8 {
				fmt.Printf("array1(%d)+array1(%d)=8\n", index, i)
			}
		}
	}
	//3. 创建切片，并增加元素
	var a = make([]int, 5, 10)
	fmt.Println(a)
	for i := 0; i < 10; i++ {
		a = append(a, i)
	}
	fmt.Printf("a=%d;len(a)=%d;cap(a)=%d\n", a, len(a), cap(a))

	// 4.对数组sort排序,需先转为切片
	var a1 = [...]int{3, 7, 8, 9, 1}
	fmt.Println(a1)
	// 方法一
	// sort.Sort(sort.IntSlice(a1[:]))
	// 方法二
	sort.Ints(a1[:])
	fmt.Println(a1)
}
