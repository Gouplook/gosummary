package common

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

//创建mysql 连接
func NewMysqlConn() (db *sql.DB, err error) {
	dataSourceName := "root:123456@tcp(127.0.0.1:3306)/imooc?charset=utf8"
	db, err = sql.Open("mysql", dataSourceName)
	return
}

//获取返回值，获取一条
func GetResultRow(rows *sql.Rows) map[string]string {
	// todo
	return nil
}

//获取所有
func GetResultRows(rows *sql.Rows) map[int]map[string]string {
	// todo
	return nil
}
