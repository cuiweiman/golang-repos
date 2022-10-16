package basic

import (
	"fmt"
	"strings"
)

var initVars int = initVar()

func init() {
	fmt.Println("初始化函数，先于main()执行")
}
func initVar() int {
	fmt.Println("初始化 变量，先于 init() 和 main() 执行")
	return 10
}

func Main_function() {
	fmt.Println("-------------------简单函数-------------------")
	sumRes := sum(1, 9)
	fmt.Printf("求和: %d\n", sumRes)

	compareRes := compare(98, 70)
	fmt.Printf("比较大小: %d\n", compareRes)

	sumR, ComR := sumAndCompare(18, 12)
	fmt.Printf("求和: %d, 比较大小: %d\n", sumR, ComR)

	fmt.Println("-------------------函数变量-------------------")
	var f funType
	f = sumAndCompare
	sum_r, com_r := f(1, 2)
	fmt.Printf("求和: %d, 比较大小: %d\n", sum_r, com_r)

	fmt.Println("-------------------函数作为入参和返回值-------------------")
	var f1 funType
	f1 = sumAndCompare
	sumH, comH, f := higherOrder(1, 2, f1)
	fmt.Printf("sumH=%d, comH=%d, f.Type=%T \n", sumH, comH, f)
	fmt.Println("-------------------匿名函数-------------------")
	fnonymousTest()

	fmt.Println("-------------------闭包closure-------------------")
	fClosure := closureTest()
	v1 := fClosure(1)
	v2 := fClosure(2)
	v3 := fClosure(3)
	fmt.Printf("%d, %d, %d \n", v1, v2, v3)

	goFunc := makeSuffix(".go")
	fmt.Printf("03function+go: %s\n", goFunc("03function"))

	baseAdd, baseSub := calc(10)
	fmt.Printf("10+1=%d, -9=%d \n", baseAdd(1), baseSub(9))

	fmt.Println("-------------------递归recursion-------------------")
	num := 20
	factorial := recursionFactorial(num)
	fmt.Printf("%d!=%d\t", num, factorial)
	hanoi := recursionHanoi(num)
	fmt.Printf("%d汉诺塔次数为%d\t", num, hanoi)
	fibonacci := recursionFibonacci(num)
	fmt.Printf("斐波那契数列第%d个为%d\t", num, fibonacci)
	fmt.Println("")
	fmt.Println("-------------------defer延迟调用-------------------")
	deferRes := deferTest(num)
	fmt.Printf("逆序执行结果:%d", deferRes)

	fmt.Println("-------------------指针Demo-------------------")
	pointerDemo()
}

/* 求和 */
func sum(a int, b int) (res int) {
	res = a + b
	return res
}

/* 比较大小 */
func compare(a int, b int) (max int) {
	if a >= b {
		return a
	}
	return b
}

/* 返回两个参数 */
func sumAndCompare(a int, b int) (sum int, max int) {
	sum = a + b
	if a >= b {
		max = a
	}
	max = b
	return
}

/* 定义一个 函数变量 */
type funType func(int, int) (int, int)

/* 函数作为入参和返回值 */
func higherOrder(a int, b int, f funType) (sum int, coms int, res funType) {
	sum, coms = f(a, b)
	return sum, coms, f
}

/* 匿名函数示例 */
func fnonymousTest() {
	sub := func(a int, b int) int {
		return b - a
	}
	i := sub(1, 9)
	fmt.Printf("9-1=%d\t", i)

	res := func(c int, d int) int {
		if c >= d {
			return c
		}
		return d
	}(9, 10)
	fmt.Printf("%d与%d比大: %d 赢\n", 9, 10, res)
}

/* 闭包示例 */
func closureTest() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

/* 闭包示例，判断name无后缀时添加后缀 */
func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

/* 闭包示例 */
func calc(base int) (func(int) int, func(int) int) {
	add := func(a int) int {
		base += a
		return base
	}
	sub := func(b int) int {
		base -= b
		return base
	}
	return add, sub
}

/* 递归函数——阶乘 */
func recursionFactorial(num int) int {
	if num == 1 {
		return 1
	}
	return num * recursionFactorial(num-1)
}

/* 递归函数——汉诺塔Hanoi */
func recursionHanoi(num int) int {
	if num == 1 {
		return 1
	}
	return 1 + 2*recursionHanoi(num-1)
}

/* 递归函数——斐波那契(兔子出生数列) */
func recursionFibonacci(num int) int {
	// 小兔子在第二个月具备繁殖能力
	if num == 1 || num == 2 {
		return 1
	}
	// f(n)=f(n-1)+f(n-2), f(1)=f(2)=1
	return recursionFibonacci(num-1) + recursionFibonacci(num-2)
}

/* defer 关键字 使用 */
// defer的执行时机: https://blog.csdn.net/qq_36867807/article/details/116406954
func deferTest(num int) (x int) {
	/* - 返回值开辟了一块新的内存空间0x aaaaa，赋值为x=5
	   - 运行defer x=6 。同时修改了指向0x aaaaa的值
	   - return 指向0x aaaaa区域的值 */

	defer func() {
		num *= 0
		x = num
	}()
	num += 3
	return x
}

/* 指针使用示例 */
func pointerDemo() {
	// 基础使用
	pointerBasic()
	// 元素都是 指针 类型的数组
	pointerArray()
}

func pointerBasic() {
	var ip *int
	fmt.Printf("赋值前: %v, %T \n", ip, ip)
	ip = &initVars
	fmt.Printf("赋值后: %v,%d,%T \n", ip, *ip, ip)
}

/* 元素都是 指针 类型的数组 */
func pointerArray() {
	var ptr [3]*int
	fmt.Println(ptr)
	arr := []int{1, 2, 3}
	for i := range arr {
		ptr[i] = &arr[i]
	}
	for _, v := range ptr {
		fmt.Printf("%v,%d; ", &v, *v)
	}
}
