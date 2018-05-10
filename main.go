package main

import (
	"fmt"
	"internetOfThings/config"
	_ "github.com/Go-SQL-Driver/MySQL"
	"net/http"
	"internetOfThings/controller"
)

func main() {
	config.SqlOpen()
	fmt.Println("hello world!")
	http.HandleFunc("/login", controller.HandleLogin)
	http.ListenAndServe(":9090", nil)
	config.SqlClose()
}
