package main

import (
	"fmt"

	"github.com/nmapworld/GO_Learning/19.packet/api"
	//匿名导入包  _ "路径"
)

//包的路径从GOPATH/src开始
// 想被其他包调用，标识符首字母大写
// 单行、多行导入
// 导入包时可指定别名
// 导入包但不使用内部标识，可使用匿名导入
// 每个包导入的时候会自动执行init()函数【没有参数、返回值】,不能手动调用
// 后import用的包，先执行init()

func main() {
	a := api.Add(10, 70)
	fmt.Println(a)
}
