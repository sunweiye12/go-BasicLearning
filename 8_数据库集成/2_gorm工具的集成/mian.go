package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {

	db, err := gorm.Open("mysql", "root:wangtianA12#@tcp(localhost:3306)/test?charset=utf8")

	if err != nil {
		fmt.Println(err)
		return
	} else {
		// 说明链接成功
		fmt.Println("connection succedssed")
	}

	// 创建一个与表对应的结构体
	type mysql_test struct {
		ID   int    `gorm:"primary_key"`
		Name string `gorm:"not_null"`
		Age  int    `gorm:"not_null"`
	}

	func() {
		mysql_test := &mysql_test{ID: 4, Name: "zhangsan", Age: 19}
		db.Create(mysql_test)
	}()

	defer db.Close()
}
