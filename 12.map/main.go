package main

import (
	"fmt"
)

func main() {
	var map1 map[int]string
	if map1 == nil {
		map1 = make(map[int]string, 10)
	}
	map1[0] = "小明"
	map1[2] = "小花"
	fmt.Println(map1)
	// 根据key，查找value，如果不存在返回对应类型的零值
	fmt.Println(map1[0])
	value, ok := map1[1]
	if !ok {
		fmt.Println("未寻找到map对应的key")
	} else {
		fmt.Println(value)
	}

	// delete,若delete的对象的key不存在，则Nothing to do
	delete(map1, 0)
	fmt.Println(map1)

}
