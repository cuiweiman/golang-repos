package main

import (
	"net"
	"strings"
)

/* 用户名称、ip地址、通讯channel、与服务端的连接对象 */
type User struct {
	Name    string
	Addr    string
	Channel chan string
	// 服务器 的 Socket 连接对象
	conn net.Conn
	// 当前客户端 连接的服务器
	server *Server
}

/* 创建新用户: 根据[服务端]接收到的连接信息 创建 */
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()

	// 指针形式返回
	user := &User{
		Name:    userAddr,
		Addr:    userAddr,
		Channel: make(chan string),
		// 与服务端的 Socket 连接
		conn:   conn,
		server: server,
	}

	// 启动 用户 Channel 消息监听 的 goroutine
	go user.ListenMsg()

	return user
}

/* 消费 客户端 Channel 消息: 监听 当前 用户的 Channel, 一旦有消息写入 Channel, 就直接客户端输出 */
func (user *User) ListenMsg() {
	for {
		msg := <-user.Channel
		user.conn.Write([]byte(msg + "\n"))
	}
}

// 用户上线 业务
func (user *User) Online() {
	// 用户上线,信息缓存入 服务端的 map 中
	user.server.mapLock.Lock()
	user.server.OnlineMap[user.Name] = user
	user.server.mapLock.Unlock()

	// 处理客户端连接信息: 广播 当前用户上线信息
	onlineMsg := "客户端 " + user.Name + " 已上线"
	user.server.BroadCast(user, onlineMsg)
}

// 用户下线 业务
func (user *User) Offline() {
	user.server.mapLock.Lock()
	delete(user.server.OnlineMap, user.Name)
	user.server.mapLock.Unlock()

	onlineMsg := "客户端 " + user.Name + " 下线"
	user.server.BroadCast(user, onlineMsg)
}

// 用户消息 处理业务: 广播业务 和 处理 指令业务
func (user *User) DoMessage(msg string) {
	// fmt.Println("User.DoMessage")
	if msg == "get-online-user" {
		// 查询当前 在线用户信息
		user.server.mapLock.Lock()
		for _, onLineUser := range user.server.OnlineMap {
			msg := "[" + onLineUser.Name + "] 在线\n"
			user.SendMsg(msg)
		}
		user.server.mapLock.Unlock()
	} else if len(msg) > 7 && msg[:7] == "rename|" {
		// 客户端 重命名 rename|张三
		newName := strings.Split(msg, "|")[1]
		// 名称 唯一性 判断
		_, ok := user.server.OnlineMap[newName]
		if ok {
			user.SendMsg("当前用户名已被使用\n")
		} else {
			user.server.mapLock.Lock()
			delete(user.server.OnlineMap, user.Name)
			user.Name = newName
			user.server.OnlineMap[newName] = user
			user.server.mapLock.Unlock()

			user.SendMsg("用户名 " + newName + " 更新成功\n")
		}
	} else if len(msg) > 4 && msg[:3] == "to|" {
		// 私聊功能,消息格式: to|张三|消息内容
		remoteName := strings.Split(msg, "|")[1]
		if remoteName == "" {
			user.SendMsg("消息格式不正确，请使用\"to|张三|消息内容.\"\n")
			return
		}
		// 判断 私聊 用户是否存在
		remoteuser, ok := user.server.OnlineMap[remoteName]
		if !ok {
			user.SendMsg("私聊用户不存在\n")
			return
		}
		content := strings.Split(msg, "|")[2]
		if content == "" {
			user.SendMsg("无消息内容\n")
			return
		}
		remoteuser.SendMsg(user.Name + ": " + content + "\n")
	} else {
		user.server.BroadCast(user, msg)
	}
}

// 向 服务端 的 Socket 连接写入消息，消息被网络传输到 客户端
func (user *User) SendMsg(msg string) {
	// fmt.Println("User.SendMsg")
	user.conn.Write([]byte(msg))
}
