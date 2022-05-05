package api

import "fmt"

// 特殊的函数，当导入所在包时会自动触发
func init() {
	fmt.Println("import api packet")
}

func Add(x, y int) int {
	return x + y
}
