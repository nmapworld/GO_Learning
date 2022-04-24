package main

import "fmt"

func main() {
	// map类型
	var m1 map[string]int
	fmt.Println(m1 == nil)        //还未初始化（没有内存地址）
	m1 = make(map[string]int, 10) //申请内存初始化
	m1["年龄"] = 18
	m1["身高"] = 2
	fmt.Println(m1)
	fmt.Println(m1["身高"])
	// 查询map中的key: value
	value, ok := (m1["年龄"])
	if !ok {
		fmt.Println("查无此key")
	} else {
		fmt.Println(value)
	}

	// map遍历键值对
	for k, v := range m1 {
		fmt.Printf("%s:%d\t", k, v)
	}
	fmt.Println()

	// map遍历键
	fmt.Println("map遍历键")
	for k := range m1 {
		fmt.Println(k)
	}
	// map遍历值
	fmt.Println("map遍历值")
	for _, v := range m1 {
		fmt.Println(v)
	}
	// map删除键值对
	fmt.Println("map删除键值对")
	delete(m1, "身高")
	fmt.Println(m1)
}
