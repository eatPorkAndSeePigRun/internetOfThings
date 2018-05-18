package dao

import (
	"internetOfThings/config"
	"internetOfThings/model"
	_ "github.com/Go-SQL-Driver/MySQL"
	"log"
	"fmt"
)

type UserMapper struct {
}

func (um UserMapper) SelectAll() []model.User {
	rows, err := config.Db.Query("SELECT * FROM user ")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	users := make([]model.User, 10, 30)
	for i := 0; rows.Next(); i++ {
		rows.Scan(&users[i].Id, &users[i].TrueName, &users[i].Password, &users[i].Power)
	}
	return users
}

func (um UserMapper) CheckUsername(id string) int64 {
	sql :="select count(1) from user where id = ?"
	var count int64
	err := config.Db.QueryRow(sql,id).Scan(&count)
	if err != nil {
		fmt.Println(err)
	}
	return count
}

// TODO 如果查询不成功则返回为赋值的model.User类
func (um UserMapper) SelectLogin(id string, password string) model.User {
	var user model.User
	sql := "select id, true_name, password, power, create_time, update_time from user where id = ? and password = ?"
	err := config.Db.QueryRow(sql,id,password).
		Scan(&user.Id, &user.TrueName, &user.Password, &user.Power, &user.CreateTime, &user.UpdateTime)
	if err != nil {
		fmt.Println(err)
	}
	return user
}
