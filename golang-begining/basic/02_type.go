package basic

import (
	"fmt"
)

func Main_type() {
	fmt.Println("-------------------格式化输出-------------------")
	printFormat()
	fmt.Println("-------------------流程控制-------------------")
	flowController()
	fmt.Println("-------------------goto关键字测试-------------------")
	gotoTest()
	fmt.Println("-------------------数组测试-------------------")
	arrayTest()
	fmt.Println("\n-------------------切片测试-------------------")
	sliceTest()
	fmt.Println("\n-------------------哈希测试-------------------")
	mapTest()

}

/* 流程控制 */
func flowController() {

	// 键盘输入
	/* var num int
	fmt.Print("请输入数字:")
	fmt.Scan(&num)
	if num%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	} */

	// 遍历数组
	x := [...]int{1, 2, 3}
	for index, value := range x {
		fmt.Printf("index=%d, value=%v \n", index, value)
	}

	// 遍历Map
	m := make(map[string]string, 0)
	m["name"] = "SYY"
	m["age"] = "28"
	m["email"] = "test@sina.com"
	for key, val := range m {
		fmt.Printf("key=%s, val=%s \t", key, val)
	}

	text := "b"
	switch text {
	case "a":
		// 默认 break
		fmt.Println("text=" + text)
	case "b":
		fmt.Println("text=" + text)
		// 穿透，进入下一个 case
		fallthrough
	case "c":
		fmt.Println("text=" + text)
	default:
		fmt.Println("text=default")
	}

	// switch 选择中，每个case语句结尾默认包含 break 关键字

}

func gotoTest() {
	i := 0
	j := 0
	for ; i < 5; i++ {
		for ; j < 3; j++ {
			if i == 2 && j == 1 {
				goto END_LABEL
			}
		}
		j = 0
	}
END_LABEL:
	fmt.Printf("退出双重循环,i=%d j=%d \n", i, j)
}

/* 格式化输出 */
func printFormat() {
	fmt.Printf("site: %v, site.Type=%T, site.TypeValue=%#v bool类型占位符输出:%t \n", site, site, site, true)
	fmt.Printf("指针格式化输出内存地址 %p \n", &site)
}

/* 定义结构体 */
type Website struct {
	Name string
}

var site = Website{Name: "Docker360"}

func arrayTest() {
	var a1 [2]int
	a1[1] = 5
	// 指定索引下的value
	var a2 = [4]string{0: "SYY", 1: "CWM", 3: "TEST"}
	var a3 = [...]int{1, 2}
	fmt.Printf("a1Type=%T,defaultVal=%v; a2Type=%T,defaultVal=%v; a3=%v; a3.len=%d\n", a1, a1, a2, a2, a3, len(a3))
	for i, v := range a3 {
		fmt.Printf("a3[%d]=%v, ", i, v)
	}
}

/* slice切片: 可变长度的数组，底层是数组，增加了自动扩容的功能 */
func sliceTest() {
	var slice1 []int
	array := [3]int{1, 2, 3}
	// 取数组所有元素; array[1:2] 表示 数组下标 [1,2)
	slice1 = array[:]
	fmt.Printf("slice1.Type=%T, slice1.len=%d, slice1.Cap=%d\n", slice1, len(slice1), cap(slice1))
	slice2 := append(slice1, 6)
	slice2 = append(slice2, slice1...)
	fmt.Println(slice2)

	for i, j := 0, len(slice2)-1; i < j; i, j = i+1, j-1 {
		slice2[i], slice2[j] = slice2[j], slice2[i]
	}
	fmt.Println(slice2)
	// 切片没有删除方法，只能『截断』，eg 截断slice2中下标为3，值为6的元素
	slice2 = append(slice2[:3], slice2[4:]...)
	fmt.Println(slice2)

	// 复制: 当直接复制时，只是两个变量共同指向同一个内存地址，一个变全都变
	slice2_copy := slice2
	fmt.Printf("原始: original=%v, copied=%v\n", slice2, slice2_copy)
	// append操作会重新创建新对象
	slice2_copy[0] = 100
	fmt.Printf("修改: original=%v, copied=%v\n", slice2, slice2_copy)

	// 复制: 重新分配内存地址
	slice2_copy2 := make([]int, len(slice2))
	copy(slice2_copy2, slice2)
	fmt.Printf("原始: original=%v, copied=%v\n", slice2, slice2_copy2)
	slice2_copy2[3] = 100
	fmt.Printf("修改: original=%v, copied=%v\n", slice2, slice2_copy2)
}

/* map数据类型 */
func mapTest() {
	var map1 = make(map[string]string)
	map1["name"] = "Jack"
	map1["age="] = "28"
	map1["address"] = "Jack@sina.com"
	fmt.Printf("map1: %v\n", map1)

	map2 := map[string]string{
		"name":    "Tom",
		"age":     "28",
		"address": "Tom@sina.com",
	}
	fmt.Printf("map2[name]=%s\n", map2["name"])
	for k, v := range map2 {
		fmt.Printf("k=%v,v=%v; ", k, v)
	}

	value, isExists := map2["height"]
	if isExists {
		fmt.Printf("map2[height]=%s\n", value)
	} else {
		fmt.Printf("map2存在height么:%v\n", isExists)
	}

}
