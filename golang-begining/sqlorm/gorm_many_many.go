package sqlorm

import (
	"encoding/json"
	"fmt"
	"log"

	"gorm.io/gorm"
)

// User 拥有并属于多种 language，`user_languages` 是连接表
type User struct {
	gorm.Model
	Name      string
	Languages []Language `gorm:"many2many:user_languages;"`
}

type Language struct {
	gorm.Model
	Name string
}

func initSchema() {
	Gorm_Db.AutoMigrate(&User{}, &Language{})
}

func initTableMany() {
	user_list := []User{
		{Name: "小王",
			Languages: []Language{
				{Name: "English"}, {Name: "Chinese"}, {Name: "Germany"},
			},
		},
		{Name: "Michael",
			Languages: []Language{
				{Name: "English"}, {Name: "Germany"},
			},
		},
	}
	d := Gorm_Db.Create(&user_list)
	fmt.Println("多对多插入 ", d.RowsAffected)
}

func doManyToMany() {
	var users []User
	err := Gorm_Db.Model(&User{}).Preload("Languages").Find(&users).Error
	if err != nil {
		log.Fatal(err)
	}
	for index, data := range users {
		dataStr, _ := json.Marshal(data)
		fmt.Printf("index=%d,  user = %s\n", index, dataStr)
	}
}

func ManyToMany() {
	// 初始化 GORM 客户端 连接
	InitGrom()

	// 初始化 表结构
	// initSchema()

	// 初始化 表数据
	// initTableMany()

	// 多对多 查询
	doManyToMany()
}
