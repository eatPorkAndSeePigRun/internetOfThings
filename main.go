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
	//go handleSensorData()
	go handleFront()
	fmt.Println("hello world!")
	for true {
		time.Sleep(1)
	}
}

func handleFront() {
	config.SqlOpen()
	http.HandleFunc("/Login", controller.HandleLogin)
	http.HandleFunc("/ChangePW", controller.HandleChangePW)
	http.HandleFunc("/user", controller.HandleUser)
	http.HandleFunc("/users", controller.HandleUsers)
	http.ListenAndServe(":8080", nil)
	config.SqlClose()
}

func handleSensorData() {
	util.HandleSensorData("127.0.0.1:9090")
}
