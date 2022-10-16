package ioapi

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

/* bufio 缓冲区io 默认缓冲区大小为 4096 */
func BufIoDemo() {
	var filePath string = "/Users/cuiweiman/GoProjects/golang-begining/ioapi/rwTest.txt"
	fmt.Println("----------------- bufio 创建缓冲区并读取---------------------")
	bufioReadString()
	bufioReadSlice()
	fmt.Println("----------------- bufio 创建缓冲区并写入---------------------")
	bufioWriteTo()
	bufioNewWriter(filePath)
	bufioReset()
	fmt.Println("----------------- bufio Scanner 扫描---------------------")
	bufioScan()
}

func bufioReadString() {
	r := strings.NewReader("hello world")
	r2 := bufio.NewReader(r)

	s, _ := r2.ReadString(' ')
	fmt.Printf("s: %v\n", s)

	s2, _ := r2.ReadString('\n')
	fmt.Printf("s2: %v\n", s2)

	r3 := bufio.NewReaderSize(strings.NewReader("test\nhello world\n"), 2048)
	s3_1, _ := r3.ReadString('\n')
	fmt.Printf("s3_1: %v", s3_1)
	s3_2, _ := r3.ReadString('\n')
	fmt.Printf("s3_2: %v", s3_2)
}

func bufioReadSlice() {
	s := strings.NewReader("ABC,DEF,GHI,JKL")
	r := bufio.NewReader(s)

	line, _ := r.ReadSlice(',')
	fmt.Printf("line: %q\n", line)

	line2, _ := r.ReadSlice('.')
	fmt.Printf("line2: %q\n", line2)

	// fmt.Printf("%q", []byte("hello world"))
}

func bufioWriteTo() {
	s := strings.NewReader("ABCDEFGHIJKLMN")
	r := bufio.NewReader(s)

	// 创建一个 bufio 写缓冲区，并将 字节写入
	b := bytes.NewBuffer(make([]byte, 0))
	r.WriteTo(b)

	fmt.Printf("b: %v\n", b)
}

func bufioNewWriter(filePath string) {
	f, _ := os.OpenFile(filePath, os.O_RDWR, 0755)
	defer f.Close()
	w := bufio.NewWriter(f)
	w.WriteString("bufio write Hello World!")
	w.Flush()
}

/* Reset 丢弃 缓冲区中的数据，清除所有错误，从而支持缓冲区重新写入 */
func bufioReset() {
	b := bytes.NewBuffer(make([]byte, 0))
	w := bufio.NewWriter(b)
	w.WriteString("123456789")

	b2 := bytes.NewBuffer(make([]byte, 0))
	// b2 替代 b，清空缓冲区中 b 字节的内容
	w.Reset(b2)
	// 重新向 缓冲区写入数据
	w.WriteString("abcdefg")

	fmt.Printf("缓冲区已使用的字节数: w.Buffered(): %v\n", w.Buffered())
	fmt.Printf("缓冲区中未使用的字节数: w.Available(): %v\n", w.Available())

	w.Flush()
	fmt.Printf("b: %v\n", b)
	fmt.Printf("b2: %v\n", b2)
}

/* Split 设置 Scanner 的分隔函数, 必须在 Scan 方法之前调用 */
func bufioScan() {
	r := strings.NewReader("ABC DEF GHI JKL")
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanWords)

	for s.Scan() {
		fmt.Printf("s.Text(): %v\n", s.Text())
	}

}
