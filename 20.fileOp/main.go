package main

import "github.com/nmapworld/GO_Learning/20.fileOp/Op"

func main() {
	readFile := "Op/read.go"
	writeFile := "test.txt"

	// 读取文件
	// read.Osopen(readFile)
	// read.BufioRead(readFile)
	Op.IoReader(readFile)

	// 写入文件
	newLine := "Hello World"
	// Op.OsWrite(writeFile, newLine)
	// Op.BufioWrite(writeFile, newLine)
	Op.IoWrite(writeFile, newLine)

}
