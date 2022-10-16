package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	Ip   string
	Port int
	// 在线用户存储: 客户端用户连接信息存储, 作为 全局变量需要 使用锁 保持安全性
	OnlineMap map[string]*User
	mapLock   sync.RWMutex

	// 服务端 消息 广播的 Channel
	ServerChan chan string
}

/* 创建一个 server 的接口 */
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:   ip,
		Port: port,
		// 初始化 客户端 连接 map
		OnlineMap: make(map[string]*User),
		// 初始化 服务端 Channel
		ServerChan: make(chan string),
	}
	return server
}

/* 启动 服务器 的接口 */
func (server *Server) Start() {
	// 开启 Socket 端口 监听
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", server.Ip, server.Port))
	if err != nil {
		fmt.Println("net.Listen err: ", err)
		return
	}
	// 最终 关闭 Socket 监听
	defer listener.Close()

	// 监听 服务端 Channel 并进行广播
	go server.ServerChanHandler()

	for {
		// 接收 客户端连接
		conn, err2 := listener.Accept()
		if err2 != nil {
			fmt.Println("listener.Accept err: ", err2)
			continue
		}

		// 客户端连接成功 处理连接业务: 协程处理, 避免阻塞
		go server.Handler(conn)
	}
}

/* 监听、处理 服务端 channel 消息的 handler */
func (server *Server) ServerChanHandler() {
	// 一旦 服务端的 channel 有消息，就发送给 当前 所有 在线 User 的 channel
	for {
		msg := <-server.ServerChan
		server.mapLock.Lock()
		for _, sendUser := range server.OnlineMap {
			sendUser.Channel <- msg
		}
		server.mapLock.Unlock()
		time.Sleep(time.Duration(time.Duration.Milliseconds(10)))
	}
}

/* 处理 当前客户端的 连接、消息广播、断开连接 的业务 */
func (server *Server) Handler(conn net.Conn) {
	// times := time.Now()
	// fmt.Println(times, "\t客户端 [", conn.RemoteAddr().String(), "] 建立连接成功")

	// 获取用户信息
	user := NewUser(conn, server)
	user.Online()

	// 用于服务端进行定时任务，强制踢掉 长时间空闲 的客户端
	isLive := make(chan bool)
	// 处理客户端上报信息: 匿名 gorountine 广播 客户端 上报的 消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				user.Offline()
				return
			}
			if err != nil && err != io.EOF {
				// 有报错 或 数据未读取结束
				fmt.Println("Conn Read err: ", err)
				return
			}

			// 提取 客户端上报 的 消息
			msg := string(buf[:n-1])
			user.DoMessage(msg)

			// 接收到客户端消息，表示用户活跃
			isLive <- true
		}
	}()

	// 用户活跃度 判断，空闲超时则直接踢掉
	go func() {
		for {
			select {
			case <-isLive:
				// fmt.Println("倒计时重置了")
				// 为 true 就表示 当前用户活跃，继续向下执行，表示更新定时器时间
			case <-time.After(time.Minute * 10):
				// 10秒钟 客户端 不活跃，则强制踢除
				user.SendMsg("由于不活跃，你被踢了")

				// 关闭 客户端 连接
				// delete(server.OnlineMap, user.Name)
				close(user.Channel)
				conn.Close()

				// 退出当前 Handler
				return
			}
			time.Sleep(time.Millisecond * 10)
		}
	}()
}

/* 消息广播: 将需要广播的 msg 放入 服务端 channel */
func (server *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Name + "]: " + msg
	server.ServerChan <- sendMsg
}
