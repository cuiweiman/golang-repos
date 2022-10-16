package basic

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func MathDemo() {
	fmt.Println("-------------------Math 标准库 常量 & 简单API-------------------")
	mathBasic()
	fmt.Println("-------------------Math 随机数-------------------")
	randomTest()
}

func mathBasic() {
	fmt.Printf("math.MaxFloat64: %v\n", math.MaxFloat64)
	fmt.Printf("math.MinInt64: %v\n", math.MinInt8)
	fmt.Printf("math.Pi: %v\n", math.Pi)
	fmt.Printf("math.E: %v\n", math.E)
	fmt.Printf("math.Ln2: %v\n", math.Ln2)
	// 对数计算
	fmt.Printf("以10为底100的对数: %v\n", math.Log10(100))
	// 绝对值
	fmt.Printf("-19的绝对值: %v\n", math.Abs(-19))
	// 幂计算
	fmt.Printf("2的10次幂: %v\n", math.Pow(2, 10))
	// 10的3次幂
	fmt.Printf("10的3次幂: %v\n", math.Pow10(3))
	// 开平方
	fmt.Printf("81的平方根为: %v\n", math.Sqrt(81))
	// 开立方
	fmt.Printf("729的立方根为: %v\n", math.Sqrt(729))
	// 向上 & 向下 取整
	fmt.Printf("向上去整: %v\t 向下取整: %v\n", math.Ceil(3.1), math.Floor(3.9))

	// 四舍五入: golang没有直接的函数，可以接住算法思想：参数 先加 0.5，再向下取整。(或减0.5后向上取整)
	fmt.Printf("3.56四舍五入: %v\n", math.Floor(3.51+0.5))

	// 取余计算
	fmt.Printf("取余计算: %v\n", math.Mod(10, 3))
	mod := 10 % 7
	fmt.Printf("mod: %v\n", mod)

	// 获取 整数部分 & 小数部分
	int2, frac := math.Modf(math.Pi)
	fmt.Printf("%.5f 的整数部分为 %.f, 小数部分(取两位)为 %.2f", math.Pi, int2, frac)

	fmt.Println()
}

/* 随机数 生成
[伪随机—seed相同时] 发现: 循环执行多次，随机数的生成都是一样的
*/
func randomTest() {

	for i := 0; i < 5; i++ {
		a := rand.Int()
		fmt.Println(a)
	}

	fmt.Println("==============设置的种子是 变化的，可以使每次生成的随机数不通==============")
	rand.Seed(time.Now().UnixMicro())

	// 0~100 之间的随机数
	fmt.Printf("%d \n", rand.Intn(100))
	// 生成 8位数 随机数
	fmt.Printf("%08v\n", rand.Int31n(100000000))

	var min, max int
	min = 1
	max = 3
	var randBetween int = rand.Intn(max-min) + min
	fmt.Printf("随机数区间为 [1,5): %v\n", randBetween)
}
