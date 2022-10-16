package basic

import (
	"errors"
	"fmt"
	"os"
	"time"
)

func ErrorDemo() {
	errorBasic()

	_, err := check("")
	fmt.Printf("err: %v\n", err)

	// 自定义错误
	if err := oops(); err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

func errorBasic() {
	_, err := os.OpenFile("no_file.txt", os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

func check(s string) (string, error) {
	if s == "" {
		err := errors.New("入参不能为空")
		return s, err
	} else {
		return "pass", nil
	}
}

/* 自定义错误格式: 错误信息结构体 和 错误信息 打印格式 */
type MyError struct {
	When time.Time
	Err  string
}

/* 实现 error 接口 */
func (e MyError) Error() string {
	return fmt.Sprintf("%v: %v", e.When, e.Err)
}

func oops() error {
	return MyError{
		time.Date(1001, 2, 25, 9, 30, 25, 123456789, time.UTC),
		"the file system has gone away",
	}
}
