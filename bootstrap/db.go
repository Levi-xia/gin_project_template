package bootstrap

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"project/global"
	"strconv"
)

/*
InitializeDB 初始化数据库
*/
func InitializeDB() {
	mysqlConfig := global.App.Config.Mysql

	if mysqlConfig.Database == "" {
		return
	}
	username := mysqlConfig.Username
	password := mysqlConfig.Password
	host := mysqlConfig.Host
	port := mysqlConfig.Port
	database := mysqlConfig.Database
	charset := mysqlConfig.Charset
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=TRUE&loc=Local", username, password, host, port, database, charset)
	db, err := sqlx.Connect("mysql", dsn)

	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	maxOpenConns, _ := strconv.Atoi(mysqlConfig.MaxOpenConns)
	maxIdleConns, _ := strconv.Atoi(mysqlConfig.MaxIdleConns)
	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)

	// 赋值全局变量
	global.App.DB = db
}
