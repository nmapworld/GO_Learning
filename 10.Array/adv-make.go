package main

import "fmt"

func make1() {
	// 切片的本质：框选一块连续的内存，属于引用类型，真正的数据保存在底层数组中
	// make创建指定的的切片长度和容量
	s1 := make([]int, 5, 10)
	// 定义长度为5的切片，容量为10：实际数组为5个0
	fmt.Printf("s1=%v len:%d cap:%d\n", s1, len(s1), cap(s1))
	// 定义长度为0，容量为10的切片：实际数组为空
	s2 := make([]int, 0, 10)
	fmt.Printf("s2=%v len:%d cap:%d\n", s2, len(s2), cap(s2))

	// 判断切片==nil时，其仅声明切边变量，未申请长度和容量
	// var s1 []int      //len=0;cap=0;s1==nil

	// 判断切片!=nil时，可能内存或长度值为0
	// s2:=[]int{}       //len=0;cap=0;s2!=nil
	// s3:=make([]int,0) //len=0;cap=0;s3!=nil

	// 判断切片是否为空时：len(a)==0

}

func slice_append() {

	// 切片追加元素,当本身切片固定容量存储不下时:
	//     if old cap<1024;new cap=2*old cap
	//     if old cap>=1024;new cap=old cap+1/4(old cap)*n
	//     如果new cap溢出，则new cap=new len
	// 使用append时，必须使用变量接受返回值

	city := []string{"北京", "上海", "深圳"}
	fmt.Printf("city=%s,len(city)=%d,cap(city)=%d,path(city)=%p\n", city, len(city), cap(city), &city)

	city = append(city, "西安")
	fmt.Printf("city=%s,len(city)=%d,cap(city)=%d,path(city)=%p\n", city, len(city), cap(city), &city)

	city = append(city, "成都", "苏州", "杭州")
	fmt.Printf("city=%s,len(city)=%d,cap(city)=%d,path(city)=%p\n", city, len(city), cap(city), &city)

	city2 := []string{"龙湖", "保定"}
	city2 = append(city, city2...) //表示拆开切片
	fmt.Printf("city2=%s,len(city2)=%d,cap(city2)=%d,path(city2)=%p\n", city2, len(city2), cap(city2), &city2)

}

func slice_copy() {
	// 切片copy后，内容另外保存一份
	a1 := []int{1, 3, 5}
	a2 := a1
	a3 := make([]int, len(a1))
	copy(a3, a1)
	fmt.Printf("a1=%d;len(a1)=%d;cap(a1)=%d\n", a1, len(a1), cap(a1))
	fmt.Printf("a2=%d;len(a2)=%d;cap(a2)=%d\n", a2, len(a2), cap(a2))
	fmt.Printf("a3=%d;len(a3)=%d;cap(a3)=%d\n", a3, len(a3), cap(a3))
	fmt.Println("a1[0] = 100")
	a1[0] = 100
	fmt.Printf("a1=%d;len(a1)=%d;cap(a1)=%d\n", a1, len(a1), cap(a1))
	fmt.Printf("a2=%d;len(a2)=%d;cap(a2)=%d\n", a2, len(a2), cap(a2))
	fmt.Printf("a3=%d;len(a3)=%d;cap(a3)=%d\n", a3, len(a3), cap(a3))

}

func slice_del() {
	a1 := []int{11, 22, 33, 44, 55, 66, 77}
	// 删除切片中值为33的元素
	fmt.Println(a1)
	for index, value := range a1 {
		if value == 33 {
			a1 = append(a1[:index], a1[index+1:]...)
			fmt.Println(a1)
		}
	}
	fmt.Println(a1)
}
