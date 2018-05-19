package config

import (
	"database/sql"
	_ "github.com/Go-SQL-Driver/MySQL"
)

func SqlOpen() {
	var err error
	Db, err = sql.Open(
		"mysql",
		"root:812872@tcp(127.0.0.1:3306)/dba?charset=utf8&parseTime=true")
	if err != nil {
		panic(err)
	}
}

func SqlClose() {
	Db.Close()
}
