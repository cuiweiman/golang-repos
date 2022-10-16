package basic

import (
	"fmt"
	"strings"
)

func Main_interface() {
	fmt.Println("-------------------类型定义和类型别名-------------------")
	typeDemo()
	fmt.Println("-------------------结构体-------------------")
	structDemo()
	fmt.Println("-------------------结构体方法-------------------")
	structFunc()
	fmt.Println("-------------------接口测试-------------------")
	interfaceDemo()
	fmt.Println("-------------------指针接口测试-------------------")
	interfacePointer()
	fmt.Println("-------------------构造方法模拟-------------------")
	syy, err := newPerson(-27, 27, "SYY", "syy@go.com")
	if err == nil {
		fmt.Println(*syy)
	} else {
		fmt.Println(err)
	}

}

func newPerson(id int, age int, name string, email string) (*Person, error) {
	if id < 0 {
		return nil, fmt.Errorf("id 错误")
	}
	if age < 0 {
		return nil, fmt.Errorf("age 错误")
	}
	if name == "" {
		return nil, fmt.Errorf("name 不能为空")
	}
	if email == "" {
		return nil, fmt.Errorf("email 不能为空")
	}
	return &Person{id: id, name: name, email: email}, nil
}

/* 指针接口测试 */
func interfacePointer() {
	dog := Dog{name: "黑毛"}
	dog.eat()
	fmt.Println(dog.name)

	dog_p := &Dog{name: "金毛"}
	dog_p.eat2()
	fmt.Println(dog_p.name)
}

type Pet interface {
	eat()
}
type Dog struct {
	name string
}

func (dog Dog) eat() {
	fmt.Println(dog.name + " eat")
	dog.name = "修改__" + dog.name
}
func (dog *Dog) eat2() {
	fmt.Println(dog.name + " eat")
	dog.name = "修改__" + dog.name
}

/* 接口测试示例 */
func interfaceDemo() {
	c := Computer{
		name: "MacBook Pro",
	}
	c.read()
	m := Mobile{
		name: "Mi8",
	}
	m.write()
}

/* USB 接口 */
type USB interface {
	read()
	write()
}

/* 电脑 实现 USB 接口方法 */
type Computer struct {
	name string
}

func (c Computer) read() {
	fmt.Println(c.name + " read")
}

func (c Computer) write() {
	fmt.Println(c.name + " write")
}

/* 手机 实现 USB 接口方法 */
type Mobile struct {
	name string
}

func (m Mobile) read() {
	fmt.Println(m.name + " read")
}

func (m Mobile) write() {
	fmt.Println(m.name + " write")
}

/* 猫的结构体 和 相关方法 */
type Cat struct {
	name string
}

func (cat Cat) eat() {
	fmt.Printf("%s is eating.\n", cat.name)
}

func (cat Cat) sleep() {
	fmt.Printf("%s is sleeping.\n", cat.name)
}

/* 结构体方法 示例 */
func structFunc() {
	grayCat := Cat{name: "小灰"}
	grayCat.eat()
	grayCat.sleep()
}

/* 结构体 */
type Person struct {
	id, age int
	name    string
	email   string
}

/* 结构体示例 */
func structDemo() {
	structBasicDemo()
	structPointerDemo()

	paramDemo := Person{
		id:   1,
		name: "Jerry",
		age:  12,
	}
	structParamDemo(paramDemo)
	fmt.Println(paramDemo.age)
	structParamPointer(&paramDemo)
	fmt.Println(paramDemo.age)
}

/* 结构体入参: 指针传递 */
func structParamPointer(person *Person) {
	fmt.Println(person)
	person.age = 0
}

/* 结构体入参: 值传递 */
func structParamDemo(person Person) {
	fmt.Println(person)
	person.age = 0
}

/* 指针结构体 */
func structPointerDemo() {
	jerry := Person{
		id:    1,
		name:  "jerry",
		age:   12,
		email: "jerry@sina.com",
	}
	fmt.Println(jerry)
	// var p_person *Person
	var p_person = new(Person)
	p_person = &jerry
	fmt.Printf("%T, %p, %v, %d\n", p_person, p_person, *p_person, p_person.age)

}

/* 普通结构体 */
func structBasicDemo() {
	var tom Person
	fmt.Printf("tom.name=%s\n", tom.name)

	/* 匿名结构体 */
	var car struct {
		id    int
		brand string
	}
	car.brand = "宝马"
	car.id = 123
	fmt.Println(car)

	/* 结构体初始化 */
	jerry := Person{
		id:   1,
		name: "jerry",
		age:  12,
	}
	fmt.Println(jerry)
}

/* 类型定义和类型别名 */
func typeDemo() {
	typeDefinition()
	typeAlias()
}

/* 类型定义*/
func typeDefinition() {
	type MyInt string
	var sum MyInt = "100"
	fmt.Printf("类型别名: %T, %s, %d\n", sum, sum, len(sum))
}

/* 类型别名 */
func typeAlias() {
	type String = string
	var name String = "SYY"
	fmt.Printf("类型别名: %T, %s, %d, %s\n", name, name, len(name), strings.ToUpper(name))
}
