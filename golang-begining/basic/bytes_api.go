package basic

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func BytesDemo() {
	fmt.Println("----------------- bytes 强制类型转换 ---------------------")
	byteTrans()
	fmt.Println("----------------- bytes 基本操作API ---------------------")
	bytesOper()
	fmt.Println("----------------- bytes Reader类型 ---------------------")
	bytesReader()
	fmt.Println("----------------- bytes Buffer 类型 ---------------------")
	bytesBuffer()
}

/* 强制类型转换 */
func byteTrans() {
	var i int = 100

	b := byte(i)
	fmt.Printf("b: %v\n", b)

	i2 := int(b)
	fmt.Printf("i2: %v\n", i2)

	var s string
	bs := []byte{88, 89, 90}

	// 字节 slice 转化成 字符串 s
	s = string(bs)
	fmt.Printf("s: %v\n", s)

	// 字符串 s 转化成 字节 slice
	bs1 := []byte(s)
	fmt.Printf("bs1: %v\n", bs1)

}

/* 字节 常用API */
func bytesOper() {
	b := []byte("baidu.com")
	b1 := []byte("baidu")
	b2 := []byte("Baidu")

	// 是否切片包含 bytes.Contains
	fmt.Printf("bytes.Contains(b, b1): %v\n", bytes.Contains(b, b1))
	fmt.Printf("bytes.Contains(b, b2): %v\n", bytes.Contains(b, b2))

	fmt.Printf("strings.Contains(\"Hello World\", \"World\"): %v\n", strings.Contains("Hello World", "World"))

	// 切片计数 bytes.Count
	fmt.Printf("bytes.Count(b, []byte(\"b\")): %v\n", bytes.Count(b, []byte("b")))

	// 切片重复 bytes.Repeat
	repests := []byte("重复两次;")
	fmt.Printf("string(bytes.Repeat(repests, 2)): %v\n", string(bytes.Repeat(repests, 2)))

	// 切片替换 bytes.Replace(slice,old,new,times) times 表示替换次数,-1表示替换全部
	b3 := bytes.Replace(b, []byte("d"), []byte("D"), -1)
	fmt.Printf("b3: %p, %v, %s\n", b3, b3, b3)

	// Runes切片，便于计算 byte 中汉字的长度,一个汉字3个字节: bytes.Runes
	fmt.Printf("len(repests): %v\n", len(repests))
	r := bytes.Runes(repests)
	fmt.Printf("len(r): %v\n", len(r))

	// Join Slice 拼接 二维 字节 切片
	s2 := [][]byte{[]byte("Hello"), []byte("World")}

	fmt.Printf("bytes.Join(s2, []byte(\"-\")): %s\n", bytes.Join(s2, []byte("-")))
}

/* Reader 实现了 io.Reader、io.ReaderAt、io.WriterTo、io.Seeker、io.ByteScanner、io.RuneScanner 接口，且只读 */
func bytesReader() {
	// 通过 []byte 创建 Reader 对象 re
	data := "123456789"
	re := bytes.NewReader([]byte(data))

	// 获取 Reader 中的数据可读长度
	fmt.Printf("Reader尚且可读数据为: %v\n", re.Len())
	fmt.Printf("Reader数据总长度为: %v\n", re.Size())

	// 每次读取两个字节，循环读取
	buf := make([]byte, 2)
	for {
		// 将 re 中的数据读取到 buf 缓冲区中
		len, err := re.Read(buf)
		if err == io.EOF {
			break
		}
		fmt.Printf("缓冲区中读取出所有的数据: %s\n", string(buf[:len]))
		fmt.Printf("此时Reader尚且可读数据为: %v\t; 数据总长度为: %v\n", re.Len(), re.Size())
	}

	fmt.Println("..................................................")

	// 修改 re 的偏移量，使其数据读取的偏移量回归到 零点
	re.Seek(0, 0)
	fmt.Printf("偏移量为零后Reader可读数据为: %v\t; 数据总长度为: %v\n", re.Len(), re.Size())
	for {
		readed_byte, err := re.ReadByte()
		if err == io.EOF {
			break
		}
		fmt.Printf("%s \t", string(readed_byte))
		fmt.Println("")
	}

	fmt.Println("..................................................")
	// 偏移量为 int64(5) 即从第5个字节开始读，读取 数据到 buf 中，buf是2个字节的缓冲区
	re.Seek(0, 0)
	readed_at, _ := re.ReadAt(buf, int64(5))
	fmt.Printf("readed_at: %v\t %v\n", readed_at, string(buf[:readed_at]))

}

/*
buffer 缓冲区 具有读取和写入的方法、可变大小的字节缓冲区，零值标识准备使用的空缓冲区:
	var b bytes.Buffer
	b:=new(bytes.Buffer)
	b:=bytes.NewBuffer(s []byte): 从 []byte 切片 构造buffer
	b:=bytes.NewBufferString(s string): 从 string 变量构造 buffer
*/
func bytesBuffer() {
	var b bytes.Buffer
	fmt.Printf("b: %v\n", b)
	b1 := bytes.NewBufferString("HelloWorld")
	fmt.Printf("b1: %v\n", b1)
	b2 := bytes.NewBuffer([]byte("Test"))
	fmt.Printf("b2: %v\n", b2)

	b1.WriteString("你好世界")
	fmt.Printf("b1: %v\n", b1)
}
