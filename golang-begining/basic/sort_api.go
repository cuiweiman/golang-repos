package basic

import (
	"fmt"
	"sort"
	"strings"
)

func SortDemo() {
	sortInt()
}

func sortInt() {
	arrInt := []int{1, 5, 3, 7, 3}
	sort.Ints(arrInt)
	fmt.Printf("arrInt: %v\n", arrInt)
	fmt.Printf("arrInt 是否排序好了: %v\n", sort.IntsAreSorted(arrInt))

	arrStr := []string{"abandon", "wonder", "ask", "basic", "you"}
	sort.Strings(arrStr)
	fmt.Printf("字符串排序&区分大小写:  %v\n", arrStr)

	sort.Sort(StringList(arrStr))
	fmt.Printf("自定义字符串排序接口: %v\n", arrStr)

	hanzi := sort.StringSlice{"王", "安", "石"}
	fmt.Printf("hanzi切片: %v\n", hanzi)
	sort.Strings(hanzi)
	fmt.Printf("hanzi排序后: %v\n", hanzi)

	mapTest := mapSlice{
		{"a": 12, "val": 21},
		{"a": 9, "val": 90},
		{"a": 35, "val": 53},
	}
	sort.Sort(mapTest)
	fmt.Printf("mapTest: %v\n", mapTest)

}

/* 自定义排序接口
1. 自定义排序的数据类型
2. 实现 Len 接口，用于获取 切片 的 元素个数
3. 实现 Less 接口，自定义判断大小的逻辑方法；由小到大排序: x<y时输出true; 由大到小排序: x>y时输出true
*/
type StringList []string

func (str StringList) Len() int {
	return len(str)
}

/* 接口 Less */
func (str StringList) Less(x, y int) bool {
	// 0 两个字符串相等;1 表示 x>y; -1 表示 x<y;
	i := strings.Compare(str[x], str[y])
	// 由小到大排序, 当 x<y 时输出 true
	return i <= 0
}

func (str StringList) Swap(x, y int) {
	str[x], str[y] = str[y], str[x]
}

/* 自定义排序: 复杂结构体 map 排序. map[type1]type2: type1-key类型;type2-value类型. */
type mapSlice []map[string]int64

func (l mapSlice) Len() int      { return len(l) }
func (l mapSlice) Swap(i, j int) { l[i], l[j] = l[j], l[i] }
func (l mapSlice) Less(i, j int) bool {
	// 按照 key 为 "a" 的 value 进行排序, 有点类似于 二维切片
	return l[i]["a"] < l[j]["a"]
}
