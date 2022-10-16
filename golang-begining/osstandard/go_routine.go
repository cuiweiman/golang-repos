package osstandard

import (
	"fmt"
	"os"
)

func OsGoRoutine() {
	fmt.Println("----------------- OS库 协程相关API---------------------")
	routinueDemo()
	routineAttr()
	fmt.Println("----------------- OS库 环境变量相关API---------------------")

}

func routinueDemo() {
	fmt.Printf("当前线程ID: %v, 父线程ID: %v\n", os.Getpid(), os.Getppid())
}

func routineAttr() {
	// 配置 新进程 的相关属性
	// attr := &os.ProcAttr{
	// 	// files 指定 新进程 继承的活动文件对象, 标准输入、标准输出、标准错误输出
	// 	Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
	// 	// 新进程的环境变量
	// 	Env: os.Environ(),
	// }

	// p, err := os.StartProcess("/System/Applications/Sublime Text.app",
	// 	[]string{"/Users/Applications/Sublime Text.app", "/Users/cuiweiman/Desktop/rwTest.txt"}, attr)
	// if err != nil {
	// 	fmt.Printf("err: %v\n", err)
	// }
	// fmt.Printf("p进程ID: %v, 进城p %v", p.Pid, p)

	// 根据 pid 查找进程
	p2, _ := os.FindProcess(1546)
	fmt.Printf("p2: %v\n", p2)
}
