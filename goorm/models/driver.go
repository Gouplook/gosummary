package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

func init() {
	//设置驱动数据库连接参数

	maxIdle := 10
	maxConn := 1000
	maxTime := 10086

	// test 为数据库名称
	dataSource := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8"
	dbtype := "mysql"
	//设置注册数据库
	err := orm.RegisterDataBase("default", dbtype, dataSource, maxConn, maxTime, maxIdle)
	if err != nil {
		fmt.Println("mysql register err")
	}

}
