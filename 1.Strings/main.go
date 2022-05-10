package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
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
	stringDo()
	changeString()
	learnQuer()
	chineseLen()

	boolOp()

	logic()
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
		fmt.Println("Login Success")
	} else {
		fmt.Println("Username Error!!")
	}

}

func stringDo() {

	name := "Sherlock"
	goods := "dog"
	// 字符串拼接
	words := fmt.Sprintf("%s like %s", name, goods)
	println(words)
	// 字符串长度（包含空格）
	println(len(words))
	// 字符串切割
	ret := strings.Split(words, "l")
	fmt.Println(ret)
	// 判断是否包含
	ret1 := strings.Contains(words, "like")
	fmt.Printf("%T\t", ret1)
	fmt.Println(ret1)
	// 前缀是否包含
	fmt.Println(strings.HasPrefix(words, "S"))
	// 后缀是否包含
	fmt.Println(strings.HasSuffix(words, "S"))
	// 判断字符第一次出现的位置
	fmt.Println(strings.Index(words, "l"))
	// 判断字符最后一次出现的位置
	fmt.Println(strings.LastIndex(words, "l"))
	// 列表拼接
	fmt.Println(strings.Join(ret, "+"))
}

func changeString() {
	// 字符串修改

	s1 := "Hello world"
	// (int8)英文字符串转为byte类型切片，才能修改
	s2 := []byte(s1)
	s2[3] = 'H'
	fmt.Printf("%v\n", s1)
	fmt.Printf("%c\n", s2)
	//Byte类型字符串长度读取
	fmt.Println(len(s1))
	for i := 0; i < len(s2); i++ {
		fmt.Printf("%c", s2[i])
	}
	fmt.Println(string(s2))
	fmt.Println("\n------------")
	n1 := "我的家在东北，松花江上SOHO"
	// (utf-8)非ASCII字符串转为rune类型切片，才能修改
	n2 := []rune(n1)
	n2[0] = '你'
	fmt.Println(n1)
	//rune类型字符串长度读取
	fmt.Println(len(n1))
	fmt.Println(utf8.RuneCountInString(n1))
	fmt.Printf("%c\n", n2)
	for _, c := range n2 {
		fmt.Printf("%c", c)
	}
	fmt.Println("\n------------")
	fmt.Println(string(n2))

}

func learnQuer() {
	// 打印不同变量的值和对应的类型
	var (
		n1   = 100
		Pi   = 3.14151926
		way  = false
		str1 = "今天是个好日子"
	)
	fmt.Printf("%v的数据类型为:\t%T\n", n1, n1)
	fmt.Printf("%v的数据类型为:\t%T\n", Pi, Pi)
	fmt.Printf("%v的数据类型为:\t%T\n", way, way)
	fmt.Printf("%v的数据类型为:\t%T\n", str1, str1)
}

func chineseLen() {

	strin1 := "hello中国人"
	fmt.Println(len(strin1))
	fmt.Println(utf8.RuneCountInString(strin1))
	num := len(strin1) - utf8.RuneCountInString(strin1)
	fmt.Println(num / 2)

}

func boolOp() {
	a, b := 3, 4
	r1, r2, r3, r4 := a == b, a != b, a > b, a < b
	fmt.Printf("%v:%T\t%v:%T\t%v:%T\t%v:%T\t\n", r1, r1, r2, r2, r3, r3, r4, r4)
}

func logic() {
	yes, no := true, false
	r1, r2, r3 := yes && yes, yes || no, !no
	fmt.Printf("%v:%T\t%v:%T\t%v:%T\t\n", r1, r1, r2, r2, r3, r3)
	// true:bool       true:bool       true:bool
}
