package main

import (
	"fmt"
	"time"
)

func unixStand(t int64) {
	// unix时间戳转化
	t1 := time.Unix(t, 0)
	fmt.Println(t1)
	fmt.Println(t1.Year())
	fmt.Println(t1.Month())
	fmt.Println(t1.Day())

}

func timeOp() {
	// 获取本地时间
	t := time.Now()
	fmt.Println(t)
	// now +24 hour
	t = t.Add(24 * time.Hour)
	fmt.Println(t)

}

func timetick() {
	// 定时器
	timer := time.Tick(time.Second)
	for t := range timer {
		fmt.Println(t)
	}
}

func timeFormat() {
	// 2019-08-03
	now := time.Now()
	fmt.Println(now)
	fmt.Println("Golang 时间格式: 2006-01-02 15:04:05 Mon Jan")
	fmt.Println(now.Format("2006-01-02 15:04:05 Mon Jan"))     //24小时制
	fmt.Println(now.Format("2006-01-02 03:04:05(PM) Mon Jan")) //12小时制
	time1 := now.Format("2006-01-02 15:04:05 Mon Jan")

	// Format-unix
	time, err := time.Parse("2006-01-02 15:04:05 Mon Jan", time1)
	if err != nil {
		fmt.Printf("err: %s\n", err)
		return
	}
	fmt.Println(time.Unix())

}

func subTime(newYear string) {
	t := time.Now()
	fmt.Println(t)
	// 定义时区myzone
	myzone := time.FixedZone("UTC+8", 8*3600)
	fmt.Println(t.In(myzone).Format("2006-01-02 15:04:05 Mon Jan"))
	// 对变量时间添加时区myzone
	timeObj, err := time.ParseInLocation("2006-01-02 15:04:05", newYear, myzone)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}
	fmt.Println(timeObj)
	// 求时间差tSub
	tSub := timeObj.Sub(time.Now())
	fmt.Println(tSub)

}
func main() {
	t := time.Now()
	fmt.Println(t)
	fmt.Println(t.Year())
	fmt.Println(t.Month())
	fmt.Println(t.Day())
	fmt.Println(t.Hour())
	fmt.Println(t.Minute())
	fmt.Println(t.Second())

	// uninx时间戳
	fmt.Println(t.Unix())
	// fmt.Printf("%T\n", t.Unix())
	fmt.Println(t.UnixNano())
	unixStand(t.Unix())

	// 时间操作
	timeOp()

	// 定时器
	// timetick()

	// 时间格式化
	timeFormat()

	// 时间差
	newYear := "2022-05-05 00:00:00"
	subTime(newYear)

}
