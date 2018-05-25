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
	http.HandleFunc("/user", controller.HandleUser)
	http.HandleFunc("/users",controller.HandleUsers)
	http.ListenAndServe(":9090", nil)
	config.SqlClose()
}
