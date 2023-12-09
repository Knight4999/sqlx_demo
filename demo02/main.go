package main

import (
	"fmt"
	"sqlx_demo/dao"
	"sqlx_demo/models"
)

//sqlx CRUD

// 查询多条数据
func queryDemo() {
	var u models.User

	sqlStr := "SELECT id,name,age FROM user WHERE id = ?"
	err := dao.Db.Get(&u, sqlStr, 1)
	if err != nil {
		fmt.Printf("get failed,err:%v\n", err)
		return
	}
	fmt.Println(u)
}

// 查询多条数据
func selectDemo() {
	var users []models.User
	sqlStr := "SELECT id,name,age FROM user WHERE id > ?"
	err := dao.Db.Select(&users, sqlStr, 0)
	if err != nil {
		fmt.Printf("get failed,err:%v\n", err)
		return
	}
	fmt.Println(users)
}

// 插入、更新、删除数据，语法一致
func insertDemo(name string, age int) {
	sqlStr := "INSERT INTO user(name,age) VALUES (?,?)"
	ret, err := dao.Db.Exec(sqlStr, name, age)
	if err != nil {
		fmt.Printf("insert failed,err:%v\n", err)
		return
	}
	fmt.Println(ret.LastInsertId())
}

// 使用NameExec函数插入数据
func insertDemo2(u map[string]interface{}) {
	sqlStr := "INSERT INTO user(name,age) VALUES (:name,:age)"
	ret, err := dao.Db.NamedExec(sqlStr, u)
	if err != nil {
		fmt.Printf("insert failed,err:%v\n", err)
		return
	}
	fmt.Println(ret.RowsAffected())
}

func namedQuery() {
	sqlStr := "SELECT * FROM user WHERE name=:name"
	//使用map做命名查询
	rows, err := dao.Db.NamedQuery(sqlStr, map[string]interface{}{"name": "liquid"})
	if err != nil {
		fmt.Printf("NameQuery faild,err:%v\n", err)
	}
	defer rows.Close()
	for rows.Next() {
		var u models.User
		rows.StructScan(&u) // structScan 以结构体的方式查询
		fmt.Println(u)
	}

	//使用结构体做命名查询
	u := models.User{
		Name: "Tom",
	}
	rows, err = dao.Db.NamedQuery(sqlStr, u)
	if err != nil {
		fmt.Printf("NameQuery faild,err:%v\n", err)
	}
	defer rows.Close()
	for rows.Next() {
		var u models.User
		rows.StructScan(&u)
		fmt.Println(u)
	}
}
func main() {
	err := dao.InitMysql()
	if err != nil {
		panic(err)
	}
	fmt.Println("success!")

	/*queryDemo()
	selectDemo()
	insertDemo("Spirit", 12)*/
	/*m := map[string]interface{}{
		"name": "liquid",
		"age":  26,
	}

	insertDemo2(m)*/
	namedQuery()
}
