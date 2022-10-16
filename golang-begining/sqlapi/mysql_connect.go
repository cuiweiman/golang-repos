package sqlapi

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var driver_name string = "mysql"

/* 定义 DB 为全局变量 */
var Global_DB *sql.DB

// 初始化 MySQL 连接 的 全局变量(使用时删除_suffix)
func init_suffix() {
	err := initDB()
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("Global_DB: %v\n", Global_DB)
		fmt.Println("连接成功")
	}
}

/* MySQL 连接操作 */
func MysqlConnect() {
	openMysql()
}

/* sql.Open 函数只能验证 入参格式是否正确，实际上并不创建与数据库的连接。
如果要检查数据源的名称是否真实有效，应该调用Ping方法 */
func openMysql() {
	// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	link := "root:****@tcp(127.0.0.1:3306)/db_golang"
	open_db, err := sql.Open(driver_name, link)
	if err != nil {
		panic(err)
	}
	open_db.SetConnMaxLifetime(time.Minute * 3)
	open_db.SetMaxOpenConns(10)
	open_db.SetMaxIdleConns(10)
	fmt.Printf("open_db: %v\n", open_db)
}

/* 初始化 MySQL 连接，创建的 DB 对象可以安全地被多个 goroutine 并发使用，
并且维护其自己的空闲连接池；因此 Open 函数应该仅被调用一次，很少需要关闭 DB 对象 */
func initDB() (err error) {
	dns := "root:****@tcp(127.0.0.1:3306)/db_golang"
	// 变量赋值，非定义变量，其中 db 为全局变量
	Global_DB, err = sql.Open(driver_name, dns)
	if err != nil {
		return err
	}
	err = Global_DB.Ping()
	if err != nil {
		return err
	}
	return nil
}
