package basic

import (
	"fmt"
	"time"
)

/* time 标准库 */
func TimeDemo() {
	fmt.Println("-------------------time 当前时间-------------------")
	timeGet()
	fmt.Println("-------------------time 时间戳-------------------")
	timestampGet()
	// fmt.Println("-------------------time Ticker定时任务:ticker 本质是 通道channel-------------------")
	// timeTickerDemo()
	fmt.Println("-------------------time 日期格式化-------------------")
	timeFormat()
	timeParse()
}

func timeGet() {
	now := time.Now()
	fmt.Printf("now: %T, %v\n", now, now)
	year := now.Year()
	month := now.Month()
	day := now.Day()
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, now.Hour(), now.Minute(), now.Second())
	fmt.Printf("类型: %T-%T-%T %T:%T:%T\n", year, month, day, now.Hour(), now.Minute(), now.Second())

	// 日期计算
	fmt.Printf("now.Add(time.Hour * 12): %v\n", now.Add(time.Hour*12))
	// 日期前后比较
	fmt.Printf("now.Before(now.Add(time.Hour * 12)): %v\n", now.Before(now.Add(time.Hour*12)))
	fmt.Printf("now.After(now.Add(time.Hour * 12)): %v\n", now.After(now.Add(time.Hour*12)))
}

func timestampGet() {
	now := time.Now()
	fmt.Printf("秒级时间戳 now.Unix(): %T,  %v\n", now.Unix(), now.Unix())
	fmt.Printf("毫秒级 now.UnixMilli():%T,  %v\n", now.UnixMilli(), now.UnixMilli())
	fmt.Printf("now.UnixMicro(): %v\n", now.UnixMicro())
	fmt.Printf("now.UnixNano(): %v\n", now.UnixNano())
}

/* 定时循环任务 */
func timeTickerDemo() {
	ticker := time.Tick(time.Second)
	for i := range ticker {
		fmt.Printf("每秒执行一次 i: %v\n", i)
	}
}

/* 日期格式化 */
func timeFormat() {
	now := time.Now()
	// time——字符串 24小时制: golang 诞生时间是 2006-1-2 15:4:5(2006 1、2、3、4)
	fmt.Printf("24小时制: %v\n", now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	// time——字符串 12小时制: 指定 PM
	fmt.Printf("12小时制: %v\n", now.Format("2006-01-02 03:04:05.000 PM"))

	// time——字符串 其它格式
	fmt.Printf("日期 : %v\n", now.Format("2006/01/02 Mon"))
	fmt.Printf("年月日: %v\n", now.Format("2006年01月02日 15时04分05秒"))
}

/* 日期解析: 字符串——日期 */
func timeParse() {
	now := time.Now()
	// 设置 时区 为  东八区
	loc, _ := time.LoadLocation("Asia/Shanghai")

	t, err := time.ParseInLocation("2006/01/02 15:04:05", "2022/09/26 15:57:30", loc)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Printf("t: %v\n", t)
	// 时间 t 减去 当前时间now
	fmt.Printf("t.Sub(now): %v\n", t.Sub(now))
}
