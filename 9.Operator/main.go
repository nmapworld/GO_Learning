package main

import (
	"fmt"
)

var (
	a = 5
	b = 2
)

func main() {

	// 算术运算符
	// + - * / % 加减乘除余

	fmt.Println("例: 加减乘除余")
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	// 特殊语句
	a++
	// a=a+1
	b--
	// b = b - 1
	fmt.Println(a)
	fmt.Println(b)

	relationOp()
	logicOp()
	bitOp()
	assOp()

}

func relationOp() {

	// 关系运算符
	// == != > >= < <=
	// 返回值为true & false
	fmt.Println("关系运算符")
	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a < b)
	fmt.Println(a <= b)
	fmt.Println(a > b)
	fmt.Println(a >= b)

}

func logicOp() {
	// 逻辑运算符
	// && || ! 与或非
	// 返回值为true & false
	fmt.Println("逻辑运算符")
	fmt.Println(a == 1 && b == 1)
	fmt.Println(a == 1 || b == 1)
	fmt.Println(a != b)

}

func bitOp() {
	// 位运算：针对二进制数
	// & ｜ ^ << >> 与；或；非 ；左移n位，相当于2的n次方；右移n位，除以2的n次方
	fmt.Println("位运算符")
	fmt.Println(2 & 5) //对应位都为1时，对应位为1  010:101     000>0
	fmt.Println(2 | 5) //对应位有一个为1时，对应位为1 010:101   111>7
	fmt.Println(2 ^ 5) //对应位相异时，对应位为1 010:101   111>7

	fmt.Println(1 << 10)    //等于1*(2的10次方) 1<<10=10000000000(二进制)=1024(十进制)
	fmt.Println(1024 >> 10) //等于1024/(2的10次方) 1024(十进制)=10000000000(二进制)>>10=1(二进制)

}

func assOp() {
	fmt.Println("赋值运算符")
	// 赋值运算符
	// =  += -= *= /= %= <<= >>= &= \= ^=
	//直接赋值、相加、减、乘、除、余。。。。后赋值
	x := 10
	x += 2 //x=x+2
	fmt.Println(x)
	x -= 2 //x=x-2
	fmt.Println(x)
	x *= 2
	fmt.Println(x)
	x /= 2
	fmt.Println(x)
	x %= 2
	fmt.Println(x)
	x = 1
	x <<= 10
	fmt.Println(x)
	x >>= 10
	fmt.Println(x)

}
