package basic

import "fmt"

func BuiltinDemo() {
	fmt.Println("----------------- builtin basic ---------------------")
	builtinAppend()
	fmt.Println("----------------- new 创建变量 ---------------------")
	newDemo()
	fmt.Println("----------------- make 内建函数: 创建 slice/map/channel ---------------------")
	makeDemo()
}

func builtinAppend() {
	s := []int{1, 2, 3}
	i := append(s, 100)
	fmt.Printf("slice 追加: i: %v\n", i)
	s1 := []int{9, 8}
	i2 := append(i, s1...)
	fmt.Printf("i2: %v\n", i2)

	fmt.Printf("len(s1): %v\n", len(s1))
	print("tom\t", 10, "\tTom@sina.com\n")
}

func newDemo() {
	b := new(bool)
	fmt.Printf("b.address: %v, b.val: %v\n", b, *b)
	i := new(int)
	fmt.Printf("i.address: %v, i.val: %v\n", i, *i)
	s := new(string)
	fmt.Printf("s.address: %v, s.val: %v\n", s, *s)
}

func makeDemo() {
	v := make([]int, 5)
	fmt.Printf("v: %T %v\n", v, v)
}
