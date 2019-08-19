package main

import _ "github.com/go-sql-driver/mysql"
import "database/sql"
import "fmt"

/*
go get github.com/go-sql-driver/mysql  /安装 mysql 驱动
*/
func main() {
	// 加载数据库驱动
	db, err := sql.Open("mysql", "root:wangtianA12#@tcp(localhost:3306)/test?charset=utf8")
	// 驱动加载失败之后直接报错
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	rows, err := db.Query("select id, age from mysql_test where age <= 15")
	defer rows.Close()
	for rows.Next() {
		var id int
		var age string
		err = rows.Scan(&id, &age)
		fmt.Printf("rows id = %d, age = %s\n", id, age)
	}
	err = rows.Err()
	if err != nil {
		panic(err.Error())
	}
}
