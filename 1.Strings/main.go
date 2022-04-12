package main

import (
	"fmt"
)

/*
%s	文本占位符
%d	数字占位符
%f	浮点数占位符
%.2f	保留两位小数点浮点占位符,四舍五入
百分比
*/

// 推荐使用变量命名：studentName

func main() {
	// 不推荐使用，后续版本可能取消
	print("print\\n\n")
	println("println")
	// 推荐使用
	fmt.Print("fmt.print\\n\n")
	fmt.Println("fmt.println")
	fmt.Printf("%s开着%s去到了%s,买了%d个%s，共消费了%.2f元\n", "小明", "火车", "拉萨", 4, "飞饼", 5.613)
	// 输出制表符
	fmt.Println("小\t明")
	// 输出换行
	fmt.Println("")
	// 输出一个\
	fmt.Println("\\")
	fmt.Println("小\r明")
	// MyInOut()
	MyData()
	CreatData()
}

func MyInOut() {
	// 交互式输入输出
	var age int     //default 0
	var name string //default ""
	// var IsOK bool   //default false
	// fmt.Print("请输出您的姓名：")
	// fmt.Scanf("%s", &name)
	// fmt.Print("请输出您的年龄：")
	// fmt.Scanf("%d", &age)
	// if age >= 18 {
	// 	fmt.Printf("%s,你已经成年！", name)
	// } else {
	// 	fmt.Printf("%s 同学，你还未成年！", name)

	// }
	fmt.Print("请输入您的\"姓名 年龄\":")
	_, err := fmt.Scanln(&name, &age)
	if err == nil {
		fmt.Printf("%v的年龄是%v", name, age)
		if age >= 18 {
			fmt.Println("，已经成年")
		} else {
			fmt.Println("，还未成年")
		}
	} else {
		fmt.Println("输入错误\n,格式应为：小明 18\n", err)
	}

}

func MyData() {
	a := 2
	b := 6
	// 整型
	// 基本运算
	// 显示变量数据类型
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Println(a + b)
	fmt.Println(b - a)
	fmt.Println(b / a)
	fmt.Println(b % a)
	// 字符串拼接，数据类型不同不可拼接
	fmt.Println("love" + "sexy")
	fmt.Println("love" + "69")
	// 布尔
	c := a > b
	fmt.Println(c)
	fmt.Printf("%T\n", c)
}

func CreatData() {

	var name string = "Homless"
	var age int = 46
	var sexies bool = true
	fmt.Printf("her's name %s ,now %d years old,her's man:%t\n", name, age, sexies)
	var name1 string

	fmt.Print("Input your name:")
	fmt.Scanf("%s", &name1)
	if name1 == "Sherlock" {
		fmt.Printf("Login Success")
	} else {
		fmt.Println("Username Error!!")
	}

}
