package sqlapi

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func CurlDemo() {
	fmt.Println("----------------- Delete ---------------------")
	deleteDemo()
	fmt.Println("----------------- Insert ---------------------")
	insertDemo()
	fmt.Println("----------------- Update ---------------------")
	updateDemo()
	fmt.Println("----------------- QueryRow ---------------------")
	queryById()
	fmt.Println("----------------- Query ---------------------")
	queryList()
}

func deleteDemo() {
	deleteSql := "delete from t_user where id = ?"
	ret, err := Global_DB.Exec(deleteSql, 4)
	if err != nil {
		fmt.Printf("delete failed, err: %v\n", err)
		return
	}
	rows, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("delete failed, err: %v\n", err)
		return
	}
	fmt.Printf("删除成功: %v\n", rows)
}

func insertDemo() {
	insertSql := "insert into t_user (id,username,password) values (?,?,?)"
	ret, err := Global_DB.Exec(insertSql, "4", "赵虎", "zhaohu")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	theId, err := ret.LastInsertId()
	rows, err2 := ret.RowsAffected()
	if err != nil && err2 != nil {
		fmt.Printf("err: %v\nerr2: %v \n", err, err2)
		return
	}
	fmt.Printf("theId: %v\t 插入成功: %v\n", theId, rows)

}

type user struct {
	id       int64
	username string
	password string
}

func queryById() {
	querySql := "select id,username,password from t_user where id=?"
	var u user
	// 保证 QueryRow 之后即刻调用 Scan 方法，否则持有的 数据库 连接不会被释放
	err := Global_DB.QueryRow(querySql, 3).Scan(&u.id, &u.username, &u.password)
	if err != nil {
		fmt.Printf("scan failed, err: %v\n", err)
	}
	fmt.Printf("id: %d, username: %s, password: %s \n", u.id, u.username, u.password)
}

func queryList() {
	/* querySql := "select id,username,password from t_user"
	rows, err := Global_DB.Query(querySql) */
	querySql := "select id,username,password from t_user where id >= ? and id <= ?"
	rows, err := Global_DB.Query(querySql, 2, 4)
	if err != nil {
		fmt.Printf("query failed, err: %v\n", err)
		return
	}
	// 非常重要，rows 必须关闭 才能释放 数据库连接
	defer rows.Close()

	// 循环读取数据集
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.username, &u.password)
		if err != nil {
			fmt.Printf("scan failed, err: %v\n", err)
			return
		}
		// fmt.Printf("id: %d, username: %s, password: %s \n", u.id, u.username, u.password)
		fmt.Printf("u: %v\n", u)
	}
}

func updateDemo() {
	updateSql := "update t_user set password = ? where id = ?"
	newPass := time.Now().Format("2006-01-02 15:04:05")
	ret, err := Global_DB.Exec(updateSql, newPass, 2)
	if err != nil {
		fmt.Printf("更新失败, err: %v\n", err)
		return
	}
	rows, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("更新失败, err: %v\n", err)
		return
	}
	fmt.Printf("更新成功, 更新的行数: %d\n", rows)
}
