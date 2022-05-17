package main

import (
	"encoding/json"
	"fmt"
)

// 结构体->json
type person6 struct {
	Name string `json: "name"`
	Age  int    `json: "age"`
}

func jsonM() {
	p1 := person6{
		Name: "周林",
		Age:  9000,
	}

	// 序列化
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v\n", string(b)) //"{\"Name\":\"周林\",\"Age\":9000}"

	// 反序列化
	str := `{"name":"小明","age":18}`
	var p6 person6
	json.Unmarshal([]byte(str), &p6)
	fmt.Printf("%v\n", p6) //{小明 18}

}
