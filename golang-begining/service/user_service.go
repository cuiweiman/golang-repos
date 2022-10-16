package service

import (
	"fmt"
	"golang-begining/dao"
)

func GetUserByIds(id int) {
	fmt.Printf("service__调用 Dao: %d\n", id)
	dao.GetUserById(id)
}
