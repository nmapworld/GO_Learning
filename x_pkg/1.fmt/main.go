package main

import (
	"fmt"
	"strconv"
)

func outPut(a int) {
	// 输出
	note := `
	// 格式化占位符：
	// %d:10进制	%b:2进制	%o:8进制	%x:16进制(a-f)	%X:16进制(A-F)
	// %c:字符	%s:字符串	%f:浮点型	%p:指针
	// %T:值的类型	%t:布尔
	// %v:原值	%#v:
	// %q:对应的单引号括起的字符面值
	`
	fmt.Println(note)
	fmt.Println("-------output------")
	fmt.Print("fmt.Print()无换行输出;\t")
	fmt.Print("fmt.Print()无换行输出;\t")
	fmt.Println()
	fmt.Printf("fmt.Printf()格式化输出:%d;\t", a)
	fmt.Printf("fmt.Printf()格式化输出:%d;\t", a)
	fmt.Println()
	fmt.Println("fmt.Println()换行输出;")
	fmt.Println("fmt.Println()换行输出;")

}

func inPut() int {
	fmt.Print("请输入数字：")
	var x int
	fmt.Scan(&x)
	// fmt.Printf("输入的值为：%d\n", x)
	return x
}

func inPut2() {
	// 输入姓名、年龄、身高
	var (
		name string
		age  int
		tall float64
	)
	fmt.Print("请输入【姓名 年龄 身高】:")
	fmt.Scanf("%s %d %f", &name, &age, &tall)
	// 浮点数保存小数点后一位
	tailF := strconv.FormatFloat(tall, 'f', 1, 64)
	fmt.Printf("%s的年龄时%d岁，身高达到了%s米\n", name, age, tailF)

}

// fscan

func main() {
	a := 100
	outPut(a)
	fmt.Printf("输入的值为：%d\n", inPut())
	// Scan、Scanf和Scanln从标准输入os.Stdin读取文本；
	// Fscan、Fscanf、Fscanln从指定的io.Reader接口读取文本；
	// Sscan、Sscanf、Sscanln从一个参数字符串读取文本。
	inPut2()
}
