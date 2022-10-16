package basic

import (
	"encoding/json"
	"fmt"
	"os"
)

/* encoding/json 实现JSON字符串 与 struct 的 互相转换
序列化 func Marshal(v interface{})([]byte,error)
反序列化 func Unmarshal(data []byte, v interface{})error
读取并解析 JSON 为 结构体: type  Decoder struct
写 JSON 到输出流 结构体 : type Encoder struct
*/
func JsonDemo() {
	fmt.Println("-------------------JSON 序列化 & 反序列化-------------------")
	marshalTest()
	unMarshalTest()
}

func marshalTest() {
	p := PersonTest{
		Name:  "Jerry",
		Age:   18,
		Email: "Jerry@sina.com",
	}
	jsonStr, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("jsonStr: %T,%s\n%v\n", jsonStr, jsonStr, jsonStr)

	// 使用 JSON 格式 直接将 struct 内容写入到 文件中
	file, _ := os.OpenFile("rwTest.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)
	defer file.Close()
	e := json.NewEncoder(file)
	e.Encode(p)
}

func unMarshalTest() {
	var jsonStr string = `{"Name":"Jerry","Age":18,"Email":"Jerry@sina.com"}`
	var pDecode PersonTest
	// 反序列化时 需要传入 结构体 的『内存地址指针』
	err := json.Unmarshal([]byte(jsonStr), &pDecode)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("pDecode: %v\n", pDecode)

	// 直接将 JSON 反序列化 为 map 类型
	var mapDecode map[string]interface{}
	json.Unmarshal([]byte(jsonStr), &mapDecode)
	fmt.Printf("mapDecode: %v\n", mapDecode)
}

type PersonTest struct {
	Name  string
	Age   int
	Email string
}
