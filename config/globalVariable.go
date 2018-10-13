package config

import "database/sql"

var Db *sql.DB

const (
	SUCCESS = 2000
	ERROR   = 1
)
