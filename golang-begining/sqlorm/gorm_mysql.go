package sqlorm

import (
	"database/sql"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GormMysql() {
	InitGrom()
	// gromBasic()
	// gorm 中的 使用 原生SQL
	gormRaw()
	toSql()
}

var Gorm_Db *gorm.DB

/* 初始化 GORM 与 MySQL 的客户端 */
func InitGrom() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name
	dsn := "root:****@tcp(127.0.0.1:3306)/db_golang?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	Gorm_Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}
}

func gormBasic() {
	// 创建表: 根据 结构体类型 创建数据schema, 即表(products)
	// Gorm_Db.AutoMigrate(&Product{})

	// 插入数据
	/* p := Product{
		Code:  "D1002",
		Price: 135,
	}
	Gorm_Db.Create(&p) */

	var product Product
	// 查询数据: 根据 ID 查找 一个 满足 id=1 的记录
	Gorm_Db.First(&product, 2)
	fmt.Printf("根据ID查找: %v\n", product)
	// 查询数据: 根据 其它字段 查询 一个 满足 条件 的记录
	var product2 Product
	Gorm_Db.First(&product2, "code=?", "D1002")
	fmt.Printf("根据code查找: %v\n", product2)

	// 更新数据: 更新 product 数据字段内容
	Gorm_Db.Model(&product).Updates(Product{Price: 201, Code: "A1001"})
	Gorm_Db.Last(&product, 2)
	fmt.Printf("更新后查找: %v\n", product)

	// 删除数据: 根据 ID 删除，逻辑删除，deleted_at 赋值了删除时间-代表数据被删除。
	Gorm_Db.Delete(&product, 1)

	// 查找多条记录
	var products []Product
	products_list := Gorm_Db.Find(&products)
	fmt.Printf("查询到记录数 RowsAffected: %v\n", products_list.RowsAffected)

}

type Product struct {
	// 继承 gorm 包中的 Model 结构体
	gorm.Model
	Code  string
	Price uint
}

func gormRaw() {
	// Exec DDL 语句
	// Raw DML 语句
	Gorm_Db.Raw("UPDATE products SET price = ? WHERE code = ?", 25, "D1002")

	// Raw 查询语句
	var raw_list []Product
	Gorm_Db.Raw("select id,code,price from products ").Scan(&raw_list)

	for index, data := range raw_list {
		fmt.Printf("index=%d,  product = %v\n", index, data)
	}

	var product_count int
	Gorm_Db.Raw("select count(1) from products ").Scan(&product_count)
	fmt.Printf("数据记录数: %v\n", product_count)

	var pro Product
	Gorm_Db.Where("code = @code_alias", sql.Named("code_alias", "D1002")).Find(&pro)
	fmt.Printf("pro: %v\n", pro)

}

/* 生成 SQL 字符串 */
func toSql() {
	var raw_list []Product
	var sql string = Gorm_Db.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Model(&Product{}).Where("code=?", "D1002").Limit(2).Order("price desc").Find(&raw_list)
	})
	fmt.Printf("sql: %v\n", sql)
}
