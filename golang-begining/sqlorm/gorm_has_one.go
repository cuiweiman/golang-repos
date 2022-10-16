package sqlorm

import (
	"encoding/json"
	"fmt"
	"log"

	"gorm.io/gorm"
)

/* 一个公司有多个员工，员工结构体中 companyId 作为外键
HasOne ==> 一个公司有一个员工
HasMany ==> 一个公司有多个员工
*/
type Company struct {
	gorm.Model
	Name string
	// Emplo Emplo `gorm:"foreignKey:CompanyId"`
	Emplo []Emplo `gorm:"foreignKey:CompanyId"`
}

type Emplo struct {
	gorm.Model
	Name      string
	CompanyId uint
}

func initTable() {
	// 初始化 生成 数据表
	// Gorm_Db.AutoMigrate(&Company{}, &Emplo{})

	// 初始化 报数据
	com := Company{Name: "可口可乐"}
	create_res := Gorm_Db.Create(&com)
	fmt.Println("插入 公司 影响行数", create_res.RowsAffected)

	emp1 := Emplo{Name: "张三", CompanyId: com.ID}
	emp2 := Emplo{Name: "李四", CompanyId: com.ID}
	var emp_list []Emplo = []Emplo{emp1, emp2}
	emp_res := Gorm_Db.CreateInBatches(&emp_list, 2)
	fmt.Println("插入 员工 影响行数", emp_res.RowsAffected)
}

/* 多对多 时的 插入 */
func initTable2() {
	com := Company{
		Name: "中国李宁",
		Emplo: []Emplo{
			{Name: "王朝"},
			{Name: "马汉"},
		},
	}
	Gorm_Db.Create(&com)
}

func doHasOne() {
	var companies []Company
	err := Gorm_Db.Model(&Company{}).Preload("Emplo").Find(&companies).Error
	if err != nil {
		log.Fatal(err)
	}
	jsonStr, err := json.Marshal(companies)
	if err != nil {
		log.Fatal(err)
	}
	// 默认 %v 输出 json 会是 字节类型, %s 表示以 字符串 输出
	fmt.Printf("companies: %s\n", jsonStr)
}

func HasOne() {
	// 初始化 GORM 客户端 连接
	InitGrom()

	// 初始化 表数据
	// initTable()
	initTable2()

	// 一对多 查询
	doHasOne()
}
