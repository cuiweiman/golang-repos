package basic

import (
	"fmt"
	"log"
	"os"
)

func LogDemo() {
	fmt.Println("----------------- 日志打印示例 ---------------------")
	logBasic()
	// logPanic()
	// logFatal()
	fmt.Println("----------------- 日志格式配置 ---------------------")
	logSet()
	fmt.Println("----------------- logger 日志格式配置 ---------------------")
	loggerDemo()
}

func logBasic() {
	log.Print("格式:年月日 时分秒 日志信息")
	log.Printf("这里可以传参 %d", 100)
	log.Println("换行输出 ", 200, " 逗号拼接字符串")
}
func logPanic() {
	defer fmt.Println("[panic] 结束后再执行……")
	log.Panic("[panic]发生异常，并停止函数执行")
}

func logFatal() {
	defer fmt.Println("[fatal] 结束后也不会被执行: os.Exit(1) 系统级别退出进程")
	log.Fatal("[fatal] 发生异常，并停止『进程』的运行")
}

/* log 默认只会打印时间，我们还可以配置文件名、行号等信息 */
func logSet() {
	// 查看 log 当前的 打印配置枚举
	logFlag := log.Flags()
	fmt.Printf("logFlag: %v\n", logFlag)
	// 配置 log 的日志格式
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.Print("my log format")
	// 配置 日志 的打印 前缀
	log.SetPrefix("[前缀]: ")
	log.Print("配置了前缀")

	// 日志 输出到 文件中
	f, err := os.OpenFile("a.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0664)
	if err != nil {
		log.Fatal("日志文件错误")
	}
	defer f.Close()
	log.SetOutput(f)

	log.Print("输出到日志文件中")
}

/* 使用 logger 输出统一配置的日志格式 */
func loggerDemo() {
	var logger *log.Logger
	logFile, err := os.OpenFile("a.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Panic("打开日志文件错误")
	}
	logger = log.New(logFile, "[success-prefix] ", log.Ldate|log.Ltime|log.Lshortfile)
	loggerDemoTest(logger)
}
func loggerDemoTest(logger *log.Logger) {
	logger.Print("[logger] 自定义logger, 可以在 init() 初始化方法中提前定义好，并全局使用")
}
