package main

import (
	"errors"
	"fmt"
	"log"
	"sqlx_demo/dao"
)

// sqlx 事务处理
// sqlx 事务处理语法与database/sql基本一致，写法有些区别

func transactionDemo() (err error) {
	//开启一个事务，模拟两个sql同时执行
	tx, err := dao.Db.Beginx()
	if err != nil {
		fmt.Printf("begin trans failed, err:%v\n", err)
		return err
	}

	//启动一个延迟函数，进行最终收尾判断
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // 先回滚事务后，再进行panic
		} else if err != nil {
			fmt.Println("rollback!")
			tx.Rollback()
		} else {
			err = tx.Commit()
			fmt.Println("commit!")
		}
	}()
	sqlStr1 := "UPDATE user SET age=24 WHERE id = ?"
	ret1, err := tx.Exec(sqlStr1, 6)
	if err != nil {
		return err
	}
	n, err := ret1.RowsAffected()
	if err != nil {
		return err
	}
	if n != 1 {
		return errors.New("exec sqlStr1 failed")
	}
	sqlStr2 := "UPDATE user SET age=45 WHERE id = ?"
	ret2, err := tx.Exec(sqlStr2, 3)
	if err != nil {
		return err
	}
	n2, err := ret2.RowsAffected()
	if err != nil {
		return err
	}
	if n2 != 1 {
		return errors.New("exec sqlStr2 failed")
	}
	return err
}

func main() {
	err := dao.InitMysql()
	if err != nil {
		panic(err)
	}

	err = transactionDemo()
	if err != nil {
		log.Fatal(err)
	}
}
