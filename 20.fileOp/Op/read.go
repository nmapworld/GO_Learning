package Op

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func Osopen(file string) {
	// 利用Read()读取文件

	fileObj, err := os.Open(file)
	if err != nil {
		fmt.Printf("read file failed,err:%v\n", err)
		return
	}

	// 关闭文件
	defer fileObj.Close()

	// 读取文件
	// var tmp = make([]byte, 128)
	var tmp [128]byte
	for {
		n, err := fileObj.Read(tmp[:])

		// 当读取到文件末尾时
		if err == io.EOF {
			fmt.Println("read finish!")
			return
		}
		if err != nil {
			fmt.Printf("err:%v\n", err)
			return
		}
		fmt.Printf("read %d bytes\n", n)
		fmt.Println(string(tmp[:n]))
		if n < 128 {
			return
		}
	}

}

func BufioRead(file string) {

	// 利用bufio读取文件
	fileObj, err := os.Open(file)
	if err != nil {
		fmt.Printf("read file failed,err:%v\n", err)
		return
	}

	// 关闭文件
	defer fileObj.Close()

	// 创建一个从文件读取的对象
	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("read finish!")
			return
		}
		if err != nil {
			fmt.Printf("read line failed,err:%v\n", err)
			return
		}
		fmt.Print(line)
	}

}

func IoReader(file string) {
	ret, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("read line failed,err:%v\n", err)
		return
	}
	fmt.Println(string(ret))

}
