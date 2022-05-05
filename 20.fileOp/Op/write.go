package Op

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func OsWrite(file, newLine string) {
	// fileObj, err := os.OpenFile(file, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644) //写入文件，当文件不存在时创建,存在时进行追加
	fileObj, err := os.OpenFile(file, os.O_WRONLY|os.O_APPEND|os.O_TRUNC, 0644) //写入文件，当文件不存在时创建,存在时进行清空

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	defer fileObj.Close()
	// write
	// fileObj.Write([]byte(newLine))

	// WriteString
	fileObj.WriteString(newLine)

}

func BufioWrite(file, newLine string) {
	fileObj, err := os.OpenFile(file, os.O_WRONLY|os.O_APPEND|os.O_TRUNC, 0644) //写入文件，当文件不存在时创建,存在时进行清空

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	defer fileObj.Close()

	// 创建写的对象
	w := bufio.NewWriter(fileObj)
	w.WriteString(newLine) //写入缓存中
	w.Flush()              //缓存写入硬盘

}

func IoWrite(writeFile, newLine string) {
	err := ioutil.WriteFile(writeFile, []byte(newLine), 0666) //写入文件，当文件不存在时创建,存在时进行清空
	if err != nil {
		fmt.Printf("read line failed,err:%v\n", err)
		return
	}

}
