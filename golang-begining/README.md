[TOC]

# golang-begining
[golang study学习文档](https://www.yuque.com/aceld/mo95lb/dsk886)  
[golang study学习视频](https://www.bilibili.com/video/BV1zR4y1t7Wj)
[Golang 圣经](http://shouce.jb51.net/gopl-zh/)
[Golang 技术栈](https://golang-tech-stack.com/)


## 概述
**优势**
- 部署方式极简
    - 可直接编译
    - 不依赖其他库
    - 直接运行即可部署
- 静态类行语言: 编译时检查出隐藏的大多数问题
- 语言层面的并发:
  - 天生支持高并发
  - 可以充分利用多核
- 强大的标准库
  - 内置runtime，支持垃圾回收
  - 高效的GC垃圾回收
  - 丰富的标准库

**不足**
- 包管理，大部分包都在github上
- 无泛化类型
- 所有Excepiton都用Error来处理(比较有争议)
- C的降级处理，并非无缝，没有C降级到asm那么完美(序列化问题)

### 常用命令集&关键字&标识符
- go bug:  start a bug report
- go build:  编译并生成可执行文件
- go clean:  清除编译内容
- go doc:  show documentation for package or symbol
- go env:  打印环境变量
- go fix:  update packages to use new APIs
- go fmt:  代码格式化
- go generate:  generate Go files by processing source
- go get:  下载安装包 [依赖下载地址](https://pkg.go.dev/)
- go install:  编译并安装打包和依赖
- go list:  查看项目依赖包或者项目下的目录
- go mod:  mod管理
- go run:  运行go文件
- go test:  单元测试及基准测试常用
- go tool:  run specified go tool // 性能分析
- go version:  print Go version
- go vet:  report likely mistakes in packages

|关键字|意义|
|:--|:--|
|for|循环语句|
|break|跳出循环|
|continue|跳过本次循环|
|default|用于选择结构的默认选项（switch、select）|
|func|用于函数定义|
|select|Go 中的一个控制结构，类似于用于通信的 switch 语句。每个 case 必须是一个通信操作，要么是发送要么是接收。|
|switch|选择结构|
|case|选择结构标签|
|chan|定义channel|
|interface|定义接口|
|const|定义常量|
|defer|延迟执行函数|
|go|并发执行|
|map|map类型|
|struct|定义结构体|
|if|选择结构|
|else|选择结构|
|goto|跳转语句|
|package|包|
|fallthrough|case加上它，程序会继续执行下一条，不会判断下一条case的值|
|var|定义变量|
|return|返回|
|import|导入包|
|type|定义类型|
|range|遍历slice、map等结构元素|

|append	|bool	|byte	|cap	|close	|complex	|complex64	|complex128	|uint16|
|:--|:--|:--|:--|:--|:--|:--|:--|:--|
|copy	|false	|float32	|float64	|imag	|int	|int8	|int16	|uint32|
|int32	|int64	|iota	|len	|make	|new	|nil	|panic	|uint64|
|print	|println	|real	|recover	|string	|true	|uint	|uint8	|uintptr|


- new 可以分配任意类型的数据; make 只能分配和初始化 类型为 slice、map、chan(即channel) 的数据;
- new 返回的是 指针类型 *T; make 返回的是引用类型 T;
- new 分配的空间被清零, make 分配后会进行初始化。


### Go的命名规范
**命名规范**
- 需要对外暴露的，名字必须大写开头(public)，否则必须小写开头(private)
- 包名称: 小写，包名称需要和目录名称一致，简短且有意义，同时避免标准库冲突
- 文件名称: 小写，简短、有意义、下划线分隔单词
- 结构体名称: 驼峰命名，首字母根据 public或private 决定大小写
- 接口名称: 类同结构体名称规范，通常以 "er" 作为后缀
- 变量命名: 驼峰命名，首字母根据 public或private 决定大小写，
- 常量命名: 大写字母+下划线

**错误处理**
原则：不丢弃任何有返回err的调用
```go
if err != nil{
  // 错误问题处理
  return // 或者继续
}
// 其它正常代码
```


**单元测试**
- 文件命名: 必须 **_test.go，测试用例的函数以 Test 为前缀。


### Go的数据类型
> go可以根据 数据类型 来申请内存大小，充分利用内存。

|类型|描述|
|:--|:--|
|布尔型|常量true或false|
|数字类型|int整型/float32/float64浮点型|
|字符串|UTF-8编码的Unicode文本|
|派生类行|Pointer指针/数组/struct结构体/Channel类型/函数类型/切片类型/接口类型/Map类型|

**数字类型**
- uint8: 无符号8位整型(0~2^8-1)
- uint16: 无符号16位(0~65535)
- uint32: 略
- uint64: 略
- int8: 有符号8位整型(-128~127)
- int16: 有符号16位整型(-2^16~2^16-1)
- int32: 略
- int64: 略

**浮点型**
- float32: IEEE-754 32位浮点数
- float64: IEEE-754 64位浮点数
- complex64: 32位实数和虚数
- complex128: 64位实数和虚数

**其它数字类型**
- byte: 类似 unit8
- rune: 类似int32
- uint: 32或64位
- int: 与uint一样大小
- uintptr: 无符号整型，用于存放一个指针


### 函数特性
- go中有3种函数: 普通函数、匿名函数、方法(定义在struct上的函数)
- go中不允许函数被 重载(overload)，即不允许同名函数
- go中函数『不能嵌套』函数，但可以嵌套『匿名函数』
- 函数是一个值，可以将函数赋值给变量、使得这个变量成为函数
- 函数可以作为参数，传递给另一个函数
- 函数的返回值也可以是一个函数
- 函数调用时，若有参数，则先拷贝参数的副本，再将副本传递给函数
- 函数参数可以没有名称，只有类型

> 1. 传参时注意 形参传参、引用传参 的不同效果
> 2. go的函数可以作为 函数参数传递给另一个函数，也可以作为另一个函数的返回值返回


### 闭包
> 定义在一个函数内部的函数. 本质上, 闭包是连接函数内部与外部的桥梁, 或者说是函数与其引用环境的组合体

闭包 = 函数 + 应用环境


### defer关键字
> defer语句会将其后跟随的语句进行延迟处理。defer归属的函数即将返回时，将延迟处理的语句按defer定义的逆序进行执行。即先被defer的语句最后被执行，后被defer的语句先执行

**特性**
- 关键字 **defer** 用于注册延迟调用
- 这些调用直到 return 前才被执行，可以用于资源清理
- 多个defer 语句 按照 先进后出方式执行
- defer 语句中的变量，在 defer 声明时就决定了

**用途**
1. 关闭文件句柄
2. 锁资源释放
3. 数据库链接释放


### init函数
**特性**
- 先于main函数自动执行，不能被调用；
- init函数没有输入参数、返回值；
- 每个包可以有多个init函数；
- 包的每个源文件也可以有多个init函数；
- 同一个包中init执行顺序没有明确定义，不能依赖这个执行顺序；
- 不同包的init函数按照包导入的依赖关系决定执行顺序；

```go
var initVars int = initVar()

func init() {
	fmt.Println("初始化函数，先于main()执行")
}
func initVar() int {
	fmt.Println("初始化 变量，先于 init() 和 main() 执行")
	return 10
}
```

### 指针
go中函数传参都是值拷贝，当需要修改某个变量时，可以创建指向该变量地址的指针，通过传递指针来修改内存地址存储的值。

**& 取地址符；\* 根据地址取值**
>语法:  var var_name *var-type

## 面相对象(OOP)
### 类型定义和类型别名
**类型定义**
- 创建了一个全新的类型，与之前类型不同，因此无法使用原类型的方法.

**类型别名**
- 仅是使用别名替换之前的类型;
- 只存在于代码中，编译后别名不会存在;
- 可以使用原类型的方法;
```go
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
```

### 结构体
```go
/* 结构体-嵌套结构体 */
type Dog struct {
	id,age   int
	name string
}
type Person struct {
	id,age   int
	name string
	dog Dog
}
```

**结构体方法**
> Go语言中没有对象、类的概念，但是可以通过结构体进行模拟，可以在结构体中声明一些方法。

func (recv myType) my_method(param) return_type {}
func (recv *myType) my_method(param) return_type {}

### 接口
> 接口只能定义方法，结构体实现接口时，必须实现接口的所有方法，否则编译失败。

1. 一个结构体，可以实现多个接口
2. 多个结构体，可以实现同一个结构
3. 接口中可以嵌套接口，形成一个新的接口
	```go
	type Flyer interface{
		fly()
	}
	type Swimmer interface{
		swim()
	}
	type FlyFish interface{
		Flyer
		Swimmer
	}
	```


### Golang对OOP思想的实现
- Golang没有面向对象的概念,但是可以通过 **结构体 struct 和 函数绑定**,实现OOP中属性+方法的特性;
- 通过 **结构体嵌套**,实现OOP的继承思想;
- 通过 **普通函数** ,模拟实现 OOP 中的构造方法;

# Go进阶
## golang 包管理
> go env -w GO111MODULE=on/off

|命令|释义|
|:--|:--|
|go mod init <项目模块名称>|初始化模块|
|go mod tidy|依赖关系处理,根据go.mod文件|
|go mod vendor|将依赖包复制到项目下的vendor目录|
|go list -m all|显示依赖关系|
|go list -m -json all|显示详细的依赖关系|
|go mod download [path@version]|下载依赖|

例如本项目建立包管理方法：
```bash
# 初始化包结构
go mod init golang-begining 
# 构建 dao 包
cd dao
go build
# service 层引用
import "golang-begining/dao"
### dao.FuncName()
# 同理构造 service 层
cd service
go build
```

## Concurrency 并发编程 
### Goroutine — 协程
> golang中的并发是函数相互独立运行的能力。GoRoutines 是并发运行的函数，go提供了协程(GoRoutines)作为并发处理操作的一种方式, 执行过程更类似于 **子线程** 。[协程介绍](https://blog.csdn.net/nicolelili1/article/details/124674213)


### 协程的通道channel
> 用于goroutine之间**共享数据**，充当gorountine之间的管道并提供一种机制保证数据同步交换。

数据交换的行为有两种:
1. **缓冲通道**: 执行异步通信，数据可以持续发送到通道，接收方需要时即时接收即可。容量不足时，发送阻塞；通道空数据时，接收阻塞。
2. **无缓冲通道**: 执行goroutine之间的同步通信，保证在发送和接收发生的瞬间执行两个 gorountine 之间的数据交换，数据交换瞬时发生的。

语法：
```go
// 整型有缓冲的通道，通道大小为10
buffered:=make(chain int , 10) 
 // 整型无缓冲通道
Unbuffered:=make(chan int)

// 将数据发送到 channel
buffered <- 15
// 接收 channel 中的数据
receiver := <-buffered

```

### 协程等待: sync.WaitGroup
> var wp sync.WaitGroup 类似于 Java 中的 线程等待 (CountDownLatch) 

```go
// 加入协程数
WaitGroup.Add(coroutineCount int)
// 协程执行结束，每执行一次此方法，waitGroup 中记录的协程数目减一
WaitGroup.Done()
// 等待协程执行，waitGroup 记录的协程数目为0后等待结束。
WaitGroup.Wait()
```

### 协程互斥锁: sync.Mutex
> var mutexWp sync.WaitGroup 类似于 Java 中的 Lock 接口

### runtime 包的协程工具类
#### Gosched 让出CPU
> 当前线程让出CPU时间片，重新等待任务安排。 类似与 Java 中的 Thread.yield()。

#### Goexit 当前协程直接退出

#### GOMAXPROCS 设置CPU最大核心数
> 若不设置，则默认使用当前资源的最大核心数



## IO 标准库
io包中最重要的两个 接口: Reader 和 Writer，只要实现了这两个接口，那么就具备IO功能。
```go
type Reader interface{
	Read(p []byte) (n int, err error)
}
type Writer interface{
	Write(p []byte) (n int, err error)
}
```
例如一些实现了 Writer 和 Reader 接口的方法:
|方法|实现Reader接口|实现Writer接口|
|:--|:--|:--|
|os.File|✓|✓|
|strings.Reader|✓|x|
|bufio.Reader/Writer|✓|✓|
|bytes.Buffer|✓|✓|
|bytes.Reader|✓|x|
|compress/gzip.Reader/Writer|✓|✓|
|crypto/cipher.StreamReader/StreamWriter|✓|✓|
|crypto/tls.Conn|✓|✓|
|encoding/csv.Reader/Writer|✓|✓|


|方法|作用|
|:--|:--|
|RaedAll|读取数据,返回读到的字节 slice|
|ReadDir|读取一个目录, 返回目录入口数组 []os.FileInfo|
|ReadFile|读一个文件并返回文件内容(字节slice)|
|WriteFile|根据文件路径写入字节 slice|
|TempDir|在一个目录中创建指定前缀名的临时目录,返回新临时目录的路径|
|TempFile|在一个目录中创建指定前缀名的临时文件,返回 os.File|


## Log 使用
> golang 内置了 log 包，实现简单的日志服务，通过调用log 包的函数，可以实现简单的日志打印功能。

|函数系列|作用|
|:--|:--|
|print|单纯打印日志|
|panic|打印日志，抛出panic异常|
|fatal|打印日志，强制结束程序(os.Exit(1)), 结束时 defer 函数都不会执行|

## 操作数据库

### 连接 MySQL

1. 安装MySQL驱动
[驱动官方地址](http://pkg.go.dev/github.com/go-sql-driver/mysql)
```go
go get -u github.com/go-sql-driver/mysql
```
2. 初始化模块
go mod init m

3. 执行 go mod tidy

4. 导入驱动
```go
package main
import(
	"fmt"
	"database/sql"
	_ "github.com/go-sql-dirver/mysql"
)
func main(){}
```

### 连接MongoDB

#### 安装MonGO驱动
[驱动地址](https://pkg.go.dev/go.mongodb.org/mongo-driver)
```go
go get go.mongodb.org/mongo-driver/mongo
```
2. 初始化模块
go mod init m

3. 执行 go mod tidy

4. 导入驱动
```go
import (
    "context"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readpref"
)
```


#### Mongo的BSON
> MongoDB 中的 JSON 文档存储在名为 BSON(二进制编码的JSON) 的二进制表中。BSON 扩展了JSON表示，使其包含额外的类型，如int、long、date、浮点数 和 decimal128，使得应用程序更容易可靠地处理、排序、比较数据。

连接MongoDB的go驱动中有两大类型表示BSON数据 D 和 Raw。
**其中D包含四类:**
- D: BSON文档，顺序重要的情况下使用，比如MongoDB命令；
- M: 一张无序的map，它和D是一样的，但不保持顺序；
- A: 一个BSON数组；
- E: D中的一个元素；

使用 BSON 需导入包:
```go
import "go.mongodb.org/mongo-driver/bson"
```

```json
// 使用D类型构建的过滤器文档的李子，可以用于查找name字段与 张三 或 李四 匹配的文档
bson.D{{
	"name",
	bson.D{{
		"$in",
		bson.A{"张三", "李四"},
	}},
}}
```

**Raw**类型用于验证 字节切片 。也可以使用 **Lookup()** 从原始类型检索单个元素。如果要避免BSON反序列化程另一种类型的开销时，会很有用。



## gorm 持久层框架
[gorm-mysql 官方文档](https://gorm.io/zh_CN/docs/connecting_to_the_database.html)
> orm 对象关系映射，解决面相对象与关系型数据库字段不匹配的技术。orm 通过使用描述对象和数据库之间的映射数据，将程序对象自动匹配关系数据库内容。

### gorm 模型定义
> 模型是 go 标准的 struct，由 go 的基本数据类型、实现了 scanner 和 valuer 接口的自定义类型、指针或别名组成。详细看官方文档。

例如:
```go
type User struct{
	ID uint
	Name string 
	Email *string
	Age uint8
	Birthday *time.Time
	MemberNumber sql.NullString
	ActivatedAt sql.NullTime
	CreatedAt time.Time
	UpdatedAt time.Time
}
```


### mysql的gorm
```go
// 安装
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql

// import 引入
import(
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)
```



















