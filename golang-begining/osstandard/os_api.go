package osstandard

import (
	"fmt"
	"os"
)

/* 官方文档: https://pkg.go.dev/std */
func FileOperate() {
	var filePath string = "/Users/cuiweiman/Desktop/test.txt"
	var renameFilePath string = "/Users/cuiweiman/Desktop/test2.txt"
	var singleDirectory string = "/Users/Cuiweiman/Desktop/directoryA"
	var multiDirectory string = "/Users/Cuiweiman/Desktop/directoryA/B/C"
	fmt.Println("-------------------创建文件和目录&重命名-------------------")
	createFile(filePath)
	createDirectory(singleDirectory, multiDirectory)
	renameFile(filePath, renameFilePath)
	fmt.Println("-------------------写入文件&读取文件-------------------")
	writeFile(renameFilePath)
	readFile(renameFilePath)
	fmt.Println("-------------------删除文件和目录-------------------")
	removeDir(renameFilePath, singleDirectory)
	fmt.Println("-------------------工作目录操作-------------------")
	getWorkDir()
	getTempDir()
}

/* 创建文件 */
func createFile(filePath string) {
	f, err := os.Create(filePath)
	if err != nil {
		fmt.Println("err: ", err)
	}
	fmt.Println("f: ", f)
}

/* 文件重命名 */
func renameFile(filePath string, renameFilePath string) {
	err := os.Rename(filePath, renameFilePath)
	if err != nil {
		fmt.Println("err: ", err)
	}
}

/* 文件写入 */
func writeFile(renameFilePath string) {
	msg := "hello world"
	os.WriteFile(renameFilePath, []byte(msg), os.ModePerm)
}

/* 读取文件 */
func readFile(renameFilePath string) {
	b, err := os.ReadFile(renameFilePath)
	if err != nil {
		fmt.Println("err: ", err)
	}
	fmt.Println(string(b[:]))
}

/* 创建目录 */
func createDirectory(singleDirectory string, multiDirectory string) {
	// 单个目录
	err := os.Mkdir(singleDirectory, os.ModePerm)
	if err != nil {
		fmt.Println("err: ", err)
	}
	// 多个目录
	errAll := os.MkdirAll(multiDirectory, os.ModePerm)
	if errAll != nil {
		fmt.Println("errAll: ", errAll)
	}
}

func removeDir(renameFilePath string, singleDirectory string) {
	// 删除文件
	err := os.Remove(renameFilePath)
	if err != nil {
		fmt.Println("err: ", err)
	}
	// 删除目录
	errAll := os.RemoveAll(singleDirectory)
	if errAll != nil {
		fmt.Println("errAll: ", errAll)
	}
}

/* 获取当前工作目录 */
func getWorkDir() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("err: ", err)
	}
	fmt.Println("dir = ", dir)
}

/* 获得临时目录 */
func getTempDir() {
	s := os.TempDir()
	fmt.Println("s=", s)
}
