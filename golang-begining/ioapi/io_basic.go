package ioapi

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var commonFilePath string = "/Users/cuiweiman/Desktop/rwTest.txt"

/* 官方文档: https://pkg.go.dev/os@go1.17.3 */
func IoBasicApi() {
	fmt.Println("----------------- IO库 基本接口---------------------")
	readerDemo()
	pipeDemo()
	fmt.Println("----------------- IOUtil 工具类---------------------")
	ioUtilReadAll()
	ioUtilReadDir()
	ioUtilReadFile()
	ioUtilWriteFile()
	ioUtilTempFile()
}

func readerDemo() {
	newReader := strings.NewReader("hello world")
	buf := make([]byte, 20)
	newReader.Read(buf)
	fmt.Printf("string(buf): %v\n", string(buf))

	newReader2 := strings.NewReader("hello world")
	// 将 newReader2 的内容 复制到 系统标准输出(控制台)
	/* _, err := io.Copy(os.Stdout, newReader2)
	if err != nil {
		log.Fatal(err)
	} */
	if _, err := io.Copy(os.Stdout, newReader2); err != nil {
		log.Fatal(err)
	}
	fmt.Println()
}

/* 管道 读写 */
func pipeDemo() {
	pr, pw := io.Pipe()
	go func() {
		fmt.Fprint(pw, "some io.Reader stream to be read.\n")
		pw.Close()
	}()
	if _, err := io.Copy(os.Stdout, pr); err != nil {
		log.Fatal(err)
	}
}

/* IOUtil 工具类 Demo */
func ioUtilReadAll() {
	// newReader := strings.NewReader("hello world")

	// File 实现了 Reader 接口
	newReader, _ := os.Open(commonFilePath)
	defer newReader.Close()

	if content, err := ioutil.ReadAll(newReader); err != nil {
		log.Fatal("错误 ", "pid=", os.Getpid(), err)
	} else {
		fmt.Printf("content: %T %v\n", content, string(content))
	}
}

func ioUtilReadDir() {
	fi, _ := ioutil.ReadDir(".")
	for _, v := range fi {
		fmt.Printf("v.Name(): %v\n", v.Name())
	}
}

func ioUtilReadFile() {
	content, _ := ioutil.ReadFile(commonFilePath)
	fmt.Println(string(content))
}

func ioUtilWriteFile() {
	// 文件路径, 写入内容, 文件权限::类似方法 os.OpenFile(filePath, os.O_RDWR|os.O_TRUNC, 0755)
	err := ioutil.WriteFile(commonFilePath, []byte("ioutil#WriteFile"), 0755)
	if err != nil {
		log.Fatal(err)
	}
}

func ioUtilTempFile() {
	tmpFile, _ := ioutil.TempFile("", "example")
	defer os.Remove(tmpFile.Name())

	content := []byte("temporary file's content")
	tmpFile.Write(content)
	tmpFile.Close()

	content2, _ := ioutil.ReadFile(tmpFile.Name())
	fmt.Println(string(content2))
}
