package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"sqlx_demo/dao"
	"sqlx_demo/models"
	"strings"
)

// 使用sqlx.In 进行批量插入和查询

// 传统批量插入
func BatchInsertUser(users []*models.User) {
	//占位符切片,存放（?,?）
	valueString := make([]string, 0, len(users))
	//实际参数切片
	valueArgs := make([]interface{}, 0, len(users)*2)

	//变量users切片
	for _, u := range users {
		valueString = append(valueString, "(?,?)")
		valueArgs = append(valueArgs, u.Name)
		valueArgs = append(valueArgs, u.Age)
	}
	//使用join方法
	sqlStr := fmt.Sprintf("INSERT INTO user(name,age) VALUES %s", strings.Join(valueString, ","))
	_, err := dao.Db.Exec(sqlStr, valueArgs...)
	if err != nil {
		fmt.Printf("insert failed err:%v\n", err)
		return
	}

	fmt.Println("insert success")
}

// sqlx.In 做字符串拼接工作
func BatchInsertUser2(users []interface{}) error {
	query, args, _ := sqlx.In("INSERT INTO user (name,age) VALUES (?),(?),(?)", users...)
	fmt.Println(query) //查看生成的querystring
	fmt.Println(args)  //查看生成的args
	result, err := dao.Db.Exec(query, args...)
	if err != nil {
		return err
	}
	n, err := result.RowsAffected()
	if n != 3 {
		return err
	}
	return nil
}

// 使用NameExec批量插入
func BatchInsertUser3(users []*models.User) error {
	_, err := dao.Db.NamedExec("INSERT INTO user (name,age) VALUES (:name,:age)", users)
	return err
}

// 使用sqlx.In 进行查询
func queryBySqlx(ids []int) (users []models.User, err error) {
	query, args, err := sqlx.In("SELECT * FROM user WHERE id IN (?)", ids) //查询语句拼接
	if err != nil {
		return nil, err
	}
	fmt.Println(query)
	fmt.Println(args)
	query = dao.Db.Rebind(query) //重新绑定参数

	err = dao.Db.Select(&users, query, args...)
	return
}

// 按照指定顺序去查询
func queryAndOrderBySqlx(ids []int) (users []models.User, err error) {
	//动态填充id
	strIds := make([]string, 0, len(ids))
	for _, id := range ids {
		strIds = append(strIds, fmt.Sprintf("%d", id))
	}

	query, args, err := sqlx.In("SELECT * FROM user WHERE id IN (?) ORDER BY FIND_IN_SET(id,?)",
		ids, strings.Join(strIds, ",")) //查询语句拼接
	if err != nil {
		return nil, err
	}
	fmt.Println(query)
	fmt.Println(args)
	query = dao.Db.Rebind(query) //重新绑定参数

	err = dao.Db.Select(&users, query, args...)
	return
}
func main() {
	err := dao.InitMysql()
	if err != nil {
		panic(err)
	}

	/*users := []*models.User{
		&models.User{
			Name: "吉米",
			Age:  34,
		},
		&models.User{
			Name: "阿乐",
			Age:  46,
		},
		&models.User{
			Name: "大B",
			Age:  40,
		},
	}
	BatchInsertUser(users)*/
	/*users := []interface{}{
		models.User{
			Name: "华为",
			Age:  68,
		},
		models.User{
			Name: "小米",
			Age:  9,
		},
		models.User{
			Name: "Vivo",
			Age:  54,
		},
	}
	err = BatchInsertUser2(users)
	if err != nil {
		fmt.Printf("insert failed err:%v\n", err)
	}*/
	/*users := []*models.User{
		&models.User{
			Name: "心灵杀手2",
			Age:  98,
		},
		&models.User{
			Name: "生化危机8",
			Age:  58,
		},
		&models.User{
			Name: "漫威蜘蛛侠2",
			Age:  13,
		},
	}
	err = BatchInsertUser3(users)
	if err != nil {
		fmt.Printf("insert failed err:%v\n", err)
	}
	fmt.Println("insert success!")*/

	users, err := queryAndOrderBySqlx([]int{11, 6, 7, 9, 4, 14})
	if err != nil {
		fmt.Printf("select failed,err:%v\n", err)
	}
	for _, u := range users {
		fmt.Printf("user:%v\n", u)
	}
}
