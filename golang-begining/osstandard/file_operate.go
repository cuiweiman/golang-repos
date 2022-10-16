package osstandard

import (
	"fmt"
	"io"
	"os"
)

func FileReadWrite() {
	var directory string = "/Users/cuiweiman/Desktop"
	var filePath string = "/Users/cuiweiman/Desktop/rwTest.txt"
	// fmt.Println("----------------- 文件读操作---------------------")
	// openCloseFile(filePath)
	fmt.Println("----------------- 目录存在判断---------------------")
	fileExists(directory)
	// fmt.Println("----------------- 创建文件 ---------------------")
	// createDemo(filePath)
	fmt.Println("----------------- 文件内容读取 ---------------------")
	readDemo(filePath)

	fmt.Println("----------------- 文件写 ---------------------")
	writeDemo(filePath)
}

/* 打开&关闭文件 */
func openCloseFile(filePath string) {
	// 获取一个 只读 的文本对象,文件不存在时会报错
	/* f, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("read rile name is : %v\n", f.Name())
	f.Close() */

	// 获取一个 可读&可写&可创建 的文件对象,文件不存在时会创建, D-文件,W-写,R-读;权限 755(wxr): 111 011 011
	fWrite, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Println(fWrite.Name())
	fWrite.Close()
}

/* 判断路径下 是否存在 某目录 */
func fileExists(directory string) {
	f, _ := os.Open(directory)
	de, _ := f.ReadDir(-1)
	for index, v := range de {
		fmt.Printf("index %d v.Name(): %v  , v.IsDir(): %v\n", index, v.Name(), v.IsDir())
	}
}

/* 创建文件 */
func createDemo(filePath string) {
	// 等价于 os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	f, _ := os.Create(filePath)
	fmt.Printf("文件创建成功 f.Name(): %v\n", f.Name())
	// 创建临时文件，目录为空表示默认，temp 表示文件名前缀
	f2, _ := os.CreateTemp("", "temp")
	fmt.Printf("f2.Name(): %v\n", f2.Name())
}

func readDemo(filePath string) {
	f, _ := os.Open(filePath)
	defer f.Close()
	for {
		// 一次读取 16 个字节, buf 对象需要每次重新创建,避免脏数据覆盖
		buf := make([]byte, 16)
		n, err := f.Read(buf)
		if err == io.EOF {
			// 读取结束
			break
		}
		fmt.Printf("读取到的字节长度n: %v\n", n)
		fmt.Printf("buf: %v\n", buf)
		fmt.Printf("string(buf): %v\n", string(buf))
	}
	// 从第26个字节开始读，最多读取 16 个(一个缓冲区的长度)
	buf := make([]byte, 16)
	n, _ := f.ReadAt(buf, 26)
	fmt.Printf("字节长度n= %d,  string(buf): %v\n", n, string(buf))
}

/* 文件写 */
func writeDemo(filePath string) {
	// 原文件内容 被 从头开始覆盖写入
	// f, _ := os.OpenFile(filePath, os.O_RDWR, 0755)

	// 原文件内容 被清空，重新写入
	f, _ := os.OpenFile(filePath, os.O_RDWR|os.O_TRUNC, 0755)

	// 追加写入
	// f, _ := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0755)

	f.Write([]byte("It's writing content"))
	f.Close()
}
