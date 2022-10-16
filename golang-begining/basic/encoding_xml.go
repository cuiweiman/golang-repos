package basic

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

/*
encoding/xml 实现 xml 与 struct 的相互转换
序列化 func Marshal(v interface{})([]byte,error)
反序列化 func Unmarshal(data []byte, v interface{})error
读取并解析 xml 为 结构体: type  Decoder struct
写 xml 到输出流 结构体 : type Encoder struct
*/
func XmlDemo() {
	fmt.Println("-------------------XML 与 struct 的相互转换-------------------")
	xmlMarshalTest()
	xmlUnMarshalTest()
	fmt.Println("-------------------XML 写入文件& 文件解析-------------------")
	xmlFileTest()
}

func xmlMarshalTest() {
	xmlP := XmlStruct{
		Name:  "Tom",
		Age:   17,
		Email: "Tom@sina.com",
	}
	// 结构体,前缀,对齐
	b, _ := xml.MarshalIndent(xmlP, " ", "  ")
	fmt.Printf("b: %s\n", b)
}

func xmlUnMarshalTest() {
	var xmlStr string = `<XmlStructOutput>
	<name>Tom</name>
	<Age>17</Age>
	<email>Tom@sina.com</email>
   </XmlStructOutput>`
	var unmarshal_result XmlStruct
	xml.Unmarshal([]byte(xmlStr), &unmarshal_result)
	fmt.Printf("unmarshal_result: %v\n", unmarshal_result)
}

func xmlFileTest() {
	xmlP := XmlStruct{
		Name:  "Tom",
		Age:   17,
		Email: "Tom@sina.com",
	}

	// 将结构体 写入到 xml 文件中
	xmlFile, _ := os.OpenFile("xmlTest.xml", os.O_CREATE|os.O_RDWR, 0744)
	defer xmlFile.Close()
	e := xml.NewEncoder(xmlFile)
	e.Encode(xmlP)

	// 从 xml 文件中读取出 结构体
	ioRead, _ := ioutil.ReadFile("xmlTest.xml")
	var fromFile XmlStruct
	xml.Unmarshal([]byte(ioRead), &fromFile)
	fmt.Printf("fromFile: %v\n", fromFile)
}

type XmlStruct struct {
	XMLName xml.Name `xml:"XmlStructOutput"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age`
	Email   string   `xml:"email"`
}
