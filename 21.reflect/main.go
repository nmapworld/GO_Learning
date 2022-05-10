package main

import (
	"fmt"
	"reflect"
)

// 反射的使用

type cat struct {
	name string
}

func reflectType(x interface{}) {

	// reflect.TypeOf() 获取对象的类型
	// reflect.ValueOf() 获取对象的值
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)
	fmt.Printf("valus:%v\ttype:%v\n", v, t)
	fmt.Printf("typename:%v\t typekind:%v\n", t.Name(), t.Kind()) //t.Name():类型的名字 t.Kind()：类型的在go语言中归属的种类
	fmt.Printf("Valuekind:%v\n", v.Kind())                        //v.Kind():值的种类

}

func reflectName(n interface{}) {

	v := reflect.ValueOf(n)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		fmt.Printf("type is int64 ,value is %d\n", int64(v.Int()))
	case reflect.Float32:
		fmt.Printf("type is Float32 ,value is %f\n", float32(v.Float()))
	case reflect.Float64:
		fmt.Printf("type is Float64 ,value is %f\n", float64(v.Float()))
	}

}

func reflectSetValue(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Elem().Kind() == reflect.Int {
		v.Elem().SetInt(200)
	} else {
		fmt.Println("err")
	}
}

func main() {
	n1 := 100
	reflectType(n1) //int 100
	n2 := 3.1415
	reflectType(n2) //float64 3.1415
	c := cat{
		name: "tom",
	}
	// 反射类型
	// 反射的类型：类型(type)+种类(kind)
	reflectType(c)

	// 反射的值
	a1 := 3.1415
	reflectName(a1)

	// 修改反射的值
	// 使用值引用，则修改的副本；使用地址引用，则修改原件
	reflectSetValue(&n1)
	fmt.Println(n1)

	// 结构体反射
	reflectStruct()
}
