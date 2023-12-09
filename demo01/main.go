package main

// 连接mysql数据库
import (
	"fmt"
	"sqlx_demo/dao"
)

func main() {
	err := dao.InitMysql()
	if err != nil {
		panic(err)
	}
	fmt.Println("connect database success!")
}
