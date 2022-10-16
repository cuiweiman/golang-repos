package concurrency

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"
)

/* Go并发编程——协程 */
func Coroutine() {
	// fmt.Println("-------------------协程基础运行方式-------------------")
	// showMsg()
	// fmt.Println("-------------------协程基础运行: 实例Demo-------------------")
	// go responseSize("http://www.baidu.com")
	// go responseSize("http://www.163.com")
	// time.Sleep(time.Second * 2)
	// fmt.Println("-------------------协程通道——Channel-------------------")
	// channelDemo()
	// fmt.Println("-------------------协程通道的遍历——Channel-------------------")
	// channelTraverse()
	// channelTraverse2()
	// fmt.Println("遍历时若通道关闭,读多写少,没有值会打印默认值；若未关闭则会死锁")
	// channelTraverse3()
	// fmt.Println("协程通信——协程同步: WaitGroup")
	// waitGroupDemo()
	fmt.Println("-------------------select 遍历多个通道时，获取合适的内容-------------------")
	selectChannelDemo()
}

/* 子线程执行时，会随着主线程结束而终结 */
func showMsg() {
	go doShowMsg("GoRoutines")
	go doShowMsg("子线程")
	fmt.Println("主线程执行")
	time.Sleep(time.Second * 1)
	fmt.Println("主线程退出")
}

func doShowMsg(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("%d %s \n", i, msg)
		time.Sleep(time.Millisecond * 100)
	}
}

/* 协程使用示例, 获取 输入 URL 的页面内容  */
func responseSize(url string) {
	fmt.Println("step1: ", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("step2: 这里表示最终关闭 response 流")
	defer response.Body.Close()

	fmt.Println("step3: ", response)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("step4: body length ", len(body))
}

/* 通道Demo */
var values = make(chan int)

func channelDemo() {
	defer close(values)
	go send()
	fmt.Println("wait...")
	// 接收 channel 的数据
	value := <-values
	fmt.Printf("receive: %v\nend \n", value)
}

func send() {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(10)
	fmt.Printf("send :%v \n", value)
	// 随机数值 放入 channel
	values <- value
}

/* 协程同步：协程等待 示例 */
var wp sync.WaitGroup

func waitGroupDemo() {
	for i := 0; i < 5; i++ {
		wp.Add(1)
		go waitGroupDemoShowMsg("展示信息 " + strconv.Itoa(i))
	}

	wp.Wait()
	fmt.Println("WaitGroupDemo 协程执行结束 ")
}

func waitGroupDemoShowMsg(msg string) {
	fmt.Println(msg)
	wp.Done()
}

/* Channel 的遍历方法  For + if  */
func channelTraverse() {
	channel := make(chan int)
	// 使用 协程 执行匿名函数 进行 channel 的内容初始化
	go func() {
		for i := 0; i < 10; i++ {
			channel <- i
		}
		close(channel)
	}()
	for {
		if data, ok := <-channel; ok {
			fmt.Println("data: ", data)
		} else {
			// 若不使用 break，会 永远等待 channel 内容的写入，造成死锁
			break
		}
	}
}

func channelTraverse2() {
	channel := make(chan int)
	// 使用 协程 执行匿名函数 进行 channel 的内容初始化
	go func() {
		for i := 0; i < 10; i++ {
			channel <- i
		}
		close(channel)
	}()
	for data := range channel {
		fmt.Println("forRange 方式实现: ", data)
	}
}

func channelTraverse3() {
	channel := make(chan int)
	// 使用 协程 执行匿名函数 进行 channel 的内容初始化
	go func() {
		for i := 0; i < 10; i++ {
			channel <- i
		}
	}()
	for data := range channel {
		fmt.Println("channel未关闭导致的死锁案例: ", data)
	}
}

/* select channel：
当有一个符合的case时则直接执行；
当有多个 符合的 case 时，则随机选择一个case 执行一次；
没有符合的case ，则执行default；
没有default则直接阻塞(死锁) */
func selectChannelDemo() {
	var chanInt = make(chan int, 0)
	var chanStr = make(chan string)
	go func() {
		// 使用协程 初始化 channel
		chanInt <- 100
		chanStr <- "hello"
		// 后 defer 的 先关闭:推测defer是栈结构
		defer close(chanInt)
		defer close(chanStr)
	}()
	for {
		select {
		case r := <-chanInt:
			fmt.Println("chanInt : ", r)
		case r := <-chanStr:
			fmt.Println("chanStr : ", r)
		default:
			fmt.Println("default")
		}
		time.Sleep(time.Second)
	}
}
