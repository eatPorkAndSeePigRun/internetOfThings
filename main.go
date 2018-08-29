package main

import (
	"fmt"
	"internetOfThings/config"
	_ "github.com/Go-SQL-Driver/MySQL"
	"net/http"
	"internetOfThings/controller"
	"time"
	"internetOfThings/util"
)

func main() {
	go handleSensorData()
	go handleFront()
	fmt.Println("hello world!")
	for true  {
		time.Sleep(1)
	}
}

func handleFront()  {
	config.SqlOpen()
	http.HandleFunc("/login", controller.HandleLogin)
	http.HandleFunc("/user", controller.HandleUser)
	http.HandleFunc("/users",controller.HandleUsers)
	http.ListenAndServe(":9090", nil)
	config.SqlClose()
}

func handleSensorData()  {
	util.HandleSensorData("127.0.0.1:7070")
}
