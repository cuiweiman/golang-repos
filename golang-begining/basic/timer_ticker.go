package basic

import (
	"fmt"
	"time"
)

func Timerticker() {
	// fmt.Println("-------------------Timer 延时器: 延时后只执行一次-------------------")
	// timerDemo()
	// timerDemo2()
	// timerAfter()
	// timerStop()
	fmt.Println("-------------------Ticker 计时器: 周期循环执行-------------------")
	// tickerDemo()
	tickerDemo2()
}

/* timer 内部是通过 channel 实现的 */
func timerDemo() {
	// 创建一个定时器，2 秒后执行一次
	timer1 := time.NewTimer(time.Second * 2)
	fmt.Println("开始时间 : ", time.Now())

	// timer1.C 会造成阻塞，直到设定好的时间
	t2 := <-timer1.C
	fmt.Println("t2: ", t2)
	fmt.Println("结束时间 : ", time.Now())
}

func timerDemo2() {
	fmt.Println("only wait 开始 : ", time.Now())
	timer := time.NewTimer(time.Second * 2)
	<-timer.C
	fmt.Println("only wait 结束 : ", time.Now())
}
func timerAfter() {
	fmt.Println("2 秒的阻塞开始 : ", time.Now())
	<-time.After(time.Second * 2)
	fmt.Println("2 秒的阻塞结束 : ", time.Now())
}

/* time.Stop 关闭定时器 的等待, timer.Reset 可以重新设置时间 */
func timerStop() {
	timer3 := time.NewTimer(time.Second)
	go func() {
		<-timer3.C
		fmt.Println("Timer3 expired")
	}()
	stop := timer3.Stop()
	// 阻塞 timer 事件发生，当该函数执行后，timer计时器停止，且相应的事件不再执行
	if stop {
		fmt.Println("Timer 3 stopped")
	}
}

/* ticker 周期循环执行 */
func tickerDemo() {
	ticker := time.NewTicker(time.Second)
	counter := 0
	for _ = range ticker.C {
		counter++
		fmt.Println("ticker execute ", counter)
		if counter >= 5 {
			// 停止计时器，并结束循环
			ticker.Stop()
			break
		}
	}
}

/* ticker 周期 向 channel进行读写操作 */
func tickerDemo2() {
	chanInt := make(chan int)
	chanTicker := time.NewTicker(time.Second * 1)

	// 写入 channel
	go func() {
		for _ = range chanTicker.C {
			select {
			case chanInt <- 1:
			case chanInt <- 3:
			case chanInt <- 5:
			}
		}
	}()

	sum := 0
	for v := range chanInt {
		fmt.Println("收到数值：", v)
		sum += v
		if sum >= 10 {
			break
		}
	}
	fmt.Println("sum = ", sum)

}
