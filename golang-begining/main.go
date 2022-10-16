package main

import (
	"golang-begining/sqlorm"
)

func main() {

	// 基础变量定义的练习
	// basic.Main_variable()

	// go类型的练习
	// basic.Main_type()

	// go 方法定义的练习
	// basic.Main_function()

	// go 接口定义的练习
	// basic.Main_interface()

	// 包管理练习
	// service.GetUserByIds(1)

	// go-lang 协程练习 和 协程间的简单通讯
	// concurrency.Coroutine()

	// go-lant runtime 包 方法示例
	// concurrency.RuntimePk()

	// time 包中的 timer 定时器 和 ticker
	// concurrency.Timerticker()

	// OS 标准库: 实现了跨操作系统平台的API接口, 文件目录的创建等
	// osstandard.FileOperate()

	// 文件的读写操作
	// osstandard.FileReadWrite()

	// OS 标准库 协程相关的API
	// osstandard.OsGoRoutine()

	// OS 标准库 环境变量相关API
	// osstandard.EnvironmentApi()

	// IO 标准库
	// ioapi.IoBasicApi()

	// 缓冲 IO 流 bufio, 包装了一个 io.Reader或io.Writer接口，提供缓冲和文本io的帮助方法
	// ioapi.BufIoDemo()

	// 日志打印方法
	// ioapi.LogDemo()

	// builtin 标准库: 提供类型声明、变量、常量声明，还有一些便利函数，这个包不需要导入,变量和函数就可以直接使用。
	// basic.BuiltinDemo()

	// bytes 标准库: 字节操作常用方法
	// ioapi.BytesDemo()

	// errors 标准库
	// ioapi.ErrorDemo()

	// sort 排序的API 使用方法
	// basic.SortDemo()

	// Time 标准库
	// basic.TimeDemo()

	// JSON 格式转换
	// basic.JsonDemo()

	// xml 读取与解析
	// basic.XmlDemo()

	// Math 标准库 API
	// basic.MathDemo()

	// MySQL 连接操作
	// sqlapi.MysqlConnect()

	// MySQL CURD 操作
	// sqlapi.CurlDemo()

	// MongoDB 连接操作
	// sqlapi.MongoConnect()

	// MongoDB 数据操作
	// sqlapi.MongoCrudDemo()

	// 持久层框架 gorm 操作
	// sqlorm.GormMysql()

	// gorm 一对一,一对多 同理
	// sqlorm.HasOne()

	// gorm 多对多
	sqlorm.ManyToMany()

}
