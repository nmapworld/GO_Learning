package main

import (
	"fmt"
	"reflect"
)

// 定义学生结构体
type student struct {
	name  string `json:"name" zhoulin:"嘿嘿嘿"`
	score int    `json:"score" zhoulin:"嘿嘿嘿"`
}

func reflectStruct() {

	// 实例化结构体：student
	newStu := student{
		name:  "小王子",
		score: 99,
	}
	t := reflect.TypeOf(newStu)
	// 打印newStu类型：类型名，种类
	fmt.Println(t.Name(), t.Kind()) //student struct

	// for循环遍历结构体中的所有字段
	// NumField()方法返回结构中字段的数量，Field(i int)方法返回字段I的reflect.Value。
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("name:%v index:%v type:%v json tag %v\n", field.Name, field.Index, field.Type, field.Tag.Get("zhoulin"))
	}

	// 通过字段名获取指定结构体字段信息
	fmt.Println("通过字段名获取指定结构体字段信息")
	if scoreField, ok := t.FieldByName("score"); ok {
		fmt.Printf("name:%v index:%v type:%v json tag %v\n", scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag.Get("zhoulin"))
	}
}
