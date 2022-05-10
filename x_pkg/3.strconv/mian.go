package main

// Parse: string -> Other
// Format: Any ->string

// Quote: 将不可打印的数据，转义
// IsPrint: 判断是否可打印
// Append: 添加数据

import (
	"fmt"
	"strconv"
)

func strToInt(str string) {

	n, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v type:%T \n", n, n)
}

func intTostr(i int) {
	str := strconv.Itoa(i)
	fmt.Printf("%#v type is %T\n", str, str)

}

func stringToBool() {
	fmt.Println("string to Bool")
	fmt.Println(strconv.ParseBool("0"))
	fmt.Println(strconv.ParseBool("f"))
	fmt.Println(strconv.ParseBool("F"))
	fmt.Println(strconv.ParseBool("FALSE"))
	fmt.Println(strconv.ParseBool("false"))

	fmt.Println(strconv.ParseBool("1"))
	fmt.Println(strconv.ParseBool("t"))
	fmt.Println(strconv.ParseBool("T"))
	fmt.Println(strconv.ParseBool("TRUE"))
	fmt.Println(strconv.ParseBool("true"))

}

func stringToFloat(strF string) {
	str, _ := strconv.ParseFloat(strF, 64)
	fmt.Printf("%#v %T\n", str, str)

}

func floatToStr(f float64) {
	str := strconv.FormatFloat(f, 'f', 4, 64)
	fmt.Printf("%#v %T\n", str, str)

}

func main() {
	// str ->int
	str := "1000"
	strToInt(str)

	// int -> string
	num := 1024
	intTostr(num)

	// string -> bool
	// 当str为：1，t，T，TRUE，true，True中的一种时为真值
	// 当str为：0，f，F，FALSE，false，False中的一种时为假值
	stringToBool()

	// string -> float
	strF := "3.1415"
	stringToFloat(strF)

	// float -> string
	floatStr := 3.1415
	floatToStr(floatStr)

}
