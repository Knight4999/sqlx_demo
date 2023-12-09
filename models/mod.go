package models

import "database/sql/driver"

type User struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

// 要使用sqlx.In 要实现driver.Valuer 接口
func (u User) Value() (driver.Value, error) {
	return []interface{}{u.Name, u.Age}, nil
}
