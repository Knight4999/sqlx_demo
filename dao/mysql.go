package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
)

var Db *sqlx.DB

func InitMysql() (err error) {
	dsn := "root:Wzh@1998@tcp(127.0.0.1:3306)/dbtest2?charset=utf8mb4&parseTime=True&loc=Local"
	//MustConnect 连接不成功，直接panic
	Db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect database failed,err:%v\n", err)
		return
	}
	Db.SetConnMaxLifetime(time.Second * 10)
	Db.SetMaxOpenConns(20)
	Db.SetMaxIdleConns(10)
	return
}
