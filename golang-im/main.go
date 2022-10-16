package main

import (
	"fmt"
	"time"
)

// 编译打包  go build -o server main.go server.go user.go
// 运行服务端 ./server
func main_s() {
	// test()

	// 启动 服务器
	server := NewServer("127.0.0.1", 8890)
	server.Start()

}

var isLive = make(chan bool)

func test() {
	go do_test()

	for {
		select {
		case <-isLive:
			fmt.Println("lalala")
		case <-time.After(time.Second * 2):
			fmt.Println("时间过去了")
			return
		}
	}
}

func do_test() {
	ticker := time.NewTicker(1 * time.Second)
	counter := 0
	for _ = range ticker.C {
		isLive <- true
		counter++
		if counter >= 5 {
			ticker.Stop()
			break
		}
	}
}
