package config

import "database/sql"

var Db *sql.DB

const (
	SUCCESS = 0
	ERROR   = 1
)
