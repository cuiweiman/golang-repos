package basic

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
	"strings"
	"unsafe"
)

func Main_variable() {
	// 变量 常量
	fmt.Println("------------------变量&常量------------------")
	Variable()
	name, pi := getNameAndAge()
	fmt.Printf("#getNameAndAge() name: %s, age: %g\n", name, pi)
	Test_iota()

	fmt.Println("------------------数据类型------------------")
	// 数据类型
	normalDataType()

	fmt.Println("------------------if&for循环------------------")
	addEvenNumbers()

	fmt.Println("------------------数字类型------------------")
	numberType()
	fmt.Println("------------------进制转换------------------")
	numberSystem()

	fmt.Println("------------------字符串------------------")
	stringTest()

}

/* 定义变量 */
func Variable() {
	var name string = "SYY"
	age := 28
	hasMarry := false
	fmt.Printf("name: %s, age: %d, hasMarry: %t \n", name, age, hasMarry)

	// 批量声明
	var (
		heigh float32 = 3.1415926
		width string  = "width"
	)
	fmt.Printf("heigh: %g, width: %s\n", heigh, width)

}

/* 返回一个变量+一个常量 */
func getNameAndAge() (string, float64) {
	// 定义常量
	fmt.Printf("Go内置的PI常量 %v, 省略两位 %.2f\n", math.Pi, math.Pi)
	const PI float64 = 3.1415926
	return "SYY", PI
}

/* iota可以创建一个被修改的常量，默认起始值为1，每调用一次自增1，遇到const关键字后会被重置为0 */
func Test_iota() {
	const (
		a = iota
		b = iota
		c = iota
	)
	fmt.Printf("a=%d, b=%d, c=%d \n", a, b, c)
	// 下划线 可以跳过一次赋值，并且 iota 仍然自增
	const (
		d = iota
		_
		f = iota
	)
	fmt.Printf("d=%d, 下划线:自增但不赋值, f=%d \n", d, f)
	// 中间插一个值，类似于下划线 跳过
	const (
		x = iota
		y = 100
		z = iota
	)
	fmt.Printf("x=%d, y=%d, z=%d \n", x, y, z)
}

/* 数据类型 */
func normalDataType() {
	a := true
	// 指针类型
	b := &a
	// 数组类型
	c := [4]int{1, 2, 3, 4}
	fmt.Printf("类型打印: a=%T, b=%T, c=%T\n", a, b, c)
	fmt.Printf("b=%v, c=%v\n", b, c)

	// 切片类型: 可以动态添加元素的数组
	f := []int{1, 2, 3}
	f = append(f, 8, 9)
	fmt.Printf("f.class=%T, f=%v \n", f, f)

	// 函数类型
	fmt.Printf("%T\n", function)
}

func function() {

}

/* 判断&简单for循环 */
func addEvenNumbers() {
	sum := 0
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			sum += i
		}
	}
	fmt.Printf("sum=%d \n", sum)
}

/* 数字类型 */
func numberType() {
	var i8 int8
	var ui64 uint64
	fmt.Printf("%T, 内存大小: %dB, %d~%d \n", i8, unsafe.Sizeof(i8), math.MinInt8, math.MaxInt8)
	fmt.Printf("%T, 内存大小: %dB, %d~%d \n", ui64, unsafe.Sizeof(ui64), 0, uint64(math.MaxUint64))

	var imax = int(math.MaxInt64)
	fmt.Printf("int64最大值: imax=%d \n", imax)
}

/* 进制转换 */
func numberSystem() {
	// 十进制转换
	a := 10
	fmt.Printf("十进制 %d, 二进制 %b, 八进制 %o, 十六进制 %x \n", a, a, a, a)
	// 八进制转换
	b := 077
	fmt.Printf("十进制 %d, 二进制 %b, 八进制 %o, 十六进制 %x \n", b, b, b, b)
	// 十六进制转换
	c := 0xff
	fmt.Printf("十进制 %d, 二进制 %b, 八进制 %o, 十六进制 %x \n", c, c, c, c)
	// Go 语言无法直接定义 二进制

	// 10进制转 其它
	a2 := strconv.FormatInt(int64(a), 2)
	a3 := strconv.FormatInt(int64(a), 8)
	a4 := strconv.FormatInt(int64(a), 16)
	fmt.Printf("a2=%v, a3=%v,a4=%v\n", a2, a3, a4)

	// 2进制转10进制
	b2 := "1010"
	p2, _ := strconv.ParseInt(b2, 2, 64)
	fmt.Printf("二进制转10进制: %v\n", p2)
	// 8进制转10进制
	b3 := "345"
	p3, _ := strconv.ParseInt(b3, 8, 64)
	fmt.Printf("8进制转10进制: %v\n", p3)
	// 16进制转10进制
	b4 := "a1"
	p4, _ := strconv.ParseInt(b4, 16, 64)
	fmt.Printf("16进制转10进制: %v\n", p4)
}

func stringTest() {
	a := "SYY"
	b := `
	<head>
		<meta charset="utf-8">
		<meta name="viewport" content="width=device-width, initial-scale=1">
	</head>`
	fmt.Printf("a= %s", a)
	fmt.Println(b)

	// 字符串连接，内存类似Java字符串不可变，无用内存片段多，性能不佳
	x := "CWM"
	y := x + " & " + a
	z := fmt.Sprintf("%s Love %s", x, a)
	fmt.Printf("y=%s, z=%s\n", y, z)

	// strings.Join 内部逻辑复杂，且使用到了interface，性能不佳
	z1 := strings.Join([]string{x, a}, ",")
	fmt.Printf("字符数组以逗号连接: z1=%s\n", z1)

	// 推荐,能作为可变字符串使用，对内存增长有优化。若预知字符串长度，可以buffer.Grow()设置容量
	var buffer bytes.Buffer
	buffer.WriteString("Cui")
	buffer.WriteString(" & ")
	buffer.WriteString("Sun")
	fmt.Printf("可变字符串 buffer.String(): %v\n", buffer.String())

	// 字符串截取
	str := "Everything is gonna be ok"
	fmt.Printf("str[1]:%v, str[1]:%c, str[2:5]: %v, len(str)=%d, toUpper=%s \n", str[1], str[1], str[2:5], len(str), strings.ToUpper(str))
	fmt.Printf("strings.Split(str): %v\n", strings.Split(str, " "))

}
