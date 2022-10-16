package osstandard

import (
	"fmt"
	"os"
)

func EnvironmentApi() {
	// 获得所有 环境变量
	envInfo := os.Environ()
	fmt.Printf("envInfo: %v\n", envInfo)

	// 获得 指定 的环境变量
	goPath := os.Getenv("GOPATH")
	fmt.Printf("goPath: %v\n", goPath)

	// 查找环境变量
	javaHome, canFind := os.LookupEnv("JAVA_HOME")
	fmt.Printf("canFind: %v,  javaHome: %v\n", canFind, javaHome)

}
