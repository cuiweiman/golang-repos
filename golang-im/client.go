package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn
	command    int
}

func NewClient(serverIp string, serverPort int) *Client {
	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
		command:    99,
	}
	//  连接 server
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("net.Dial err: ", err)
		return nil
	}
	client.conn = conn

	return client
}

var serverIp string
var serverPort int

func init() {
	// 把 形参 绑定到 flag 中, 可以通过命令行 输入 参数
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置服务器IP地址(默认 127.0.0.1)")
	flag.IntVar(&serverPort, "port", 8890, "设置服务器端口(默认 8890)")
}

// 菜单指令
func (client *Client) menu() bool {
	fmt.Println("-------------------")
	fmt.Println("|请输入指令:      |")
	fmt.Println("| 1. 公聊模式     |")
	fmt.Println("| 2. 私聊模式     |")
	fmt.Println("| 3. 更新用户名   |")
	fmt.Println("| 0. 退出         |")
	fmt.Println("-------------------")

	var command int
	fmt.Scanln(&command)
	if command >= 0 && command <= 3 {
		client.command = command
		return true
	}
	fmt.Println(">>>>>>>> 指令输入不合法")
	return false
}

/* 指令执行 */
func (client *Client) Run() {
	// 当 客户端 输入指令 为 非0 时，则循环执行 客户端指令，直到 指令为0
	for client.command != 0 {
		// 输入指令若不合法，则循环执行
		for client.menu() != true {

		}
		switch client.command {
		case 1:
			fmt.Println("进入公聊频道")
			client.PublicChatting()
			break
		case 2:
			fmt.Println("进入私聊频道")
			client.PrivateChatting()
			break
		case 3:
			// fmt.Println("更新用户名")
			client.UpdateName()
			break
		}
	}
	client.Offline()
}

/* 公聊模式: 消息写入 conn 发送到 Server */
func (client *Client) PublicChatting() {
	fmt.Print(">>> 请输入消息,exit退出: ")
	var sendMsg string
	fmt.Scanln(&sendMsg)
	for sendMsg != "exit" {
		if len(sendMsg) != 0 {
			sendMsg = sendMsg + "\n"
			_, err := client.conn.Write([]byte(sendMsg))
			if err != nil {
				fmt.Println("PublicChatting Send err: ", err)
				break
			}
		}
		sendMsg = ""
		fmt.Print(">>> 请输入消息,exit退出: ")
		fmt.Scanln(&sendMsg)
	}
}

/* 私聊频道 */
func (client *Client) SelectUser() {
	sendMsg := "get-online-user\n"
	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("SelectUser error: ", err)
		return
	}
}
func (client *Client) PrivateChatting() {
	var remoteName string
	var sendMsg string

	client.SelectUser()
	fmt.Println(">>> 请输入聊天对象,exit 退出")
	fmt.Scanln(&remoteName)

	for remoteName != "exit" {
		fmt.Println("请输入消息内容,exit退出")
		fmt.Scanln(&sendMsg)
		for sendMsg != "exit" {

			if len(sendMsg) != 0 {
				sendMsg = "to|" + remoteName + "|" + sendMsg + "\n"
				_, err := client.conn.Write([]byte(sendMsg))
				if err != nil {
					fmt.Println("PrivateChatting Send err: ", err)
					break
				}
			}
			sendMsg = ""
			fmt.Println("请输入消息内容,exit退出")
			fmt.Scanln(&sendMsg)
		}
		client.SelectUser()
		fmt.Println(">>> 请输入聊天对象,exit 退出")
		fmt.Scanln(&remoteName)
	}

}

/* 更新用户名 */
func (client *Client) UpdateName() bool {
	fmt.Println(">>>>>> 请输入 用户名")
	fmt.Scanln(&client.Name)
	// 向 服务端 发送消息
	sendMsg := "rename|" + client.Name + "\n"
	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn.Write err: ", err)
		return false
	}
	return true
}

/* 客户端下线申请 */
func (client *Client) Offline() {
	err := client.conn.Close()
	if err != nil {
		fmt.Println("client.Offline error, err: ", err)
		return
	}
}

/* 处理 Server 响应的内容 */
func (client *Client) RespHandler() {
	// 永久阻塞，一旦 Socket 连接 有消息，就 copy 到 标准输出中
	io.Copy(os.Stdout, client.conn)
	fmt.Println("")
}

// 编译打包  go build -o client client.go
// 运行客户端 ./client
func main_c() {
	// 命令行 解析
	flag.Parse()

	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println(">>>>> 连接服务器失败 >>>>>")
		return
	}
	// 接收 服务端 的响应消息
	go client.RespHandler()

	fmt.Println(">>>>> 连接服务器成功 >>>>>")

	client.Run()
	// 阻塞 客户端业务
	select {}
}
