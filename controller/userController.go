package controller

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"internetOfThings/service"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		//1.解码json
		result, _ := ioutil.ReadAll(r.Body)
		r.Body.Close()
		var f interface{}
		err := json.Unmarshal([]byte(result), &f)
		if err != nil {
			fmt.Println("json err:", err)
		}
		m := f.(map[string]interface{})
		id := m["id"]
		password := m["password"]
		//2.验证用户和密码是否正确
		var us service.UserService
		serverResponse:= us.Login(id.(string), password.(string))
		//3.编码json
		b, err := json.Marshal(serverResponse)
		if err != nil {
			fmt.Println("json err:", err)
		}
		fmt.Fprint(w, string(b))
	}
}

func HandleUser(w http.ResponseWriter, r *http.Request) {
	if r.Method=="POST"{
		//1.解码json
		result, _ := ioutil.ReadAll(r.Body)
		r.Body.Close()
		var f interface{}
		err := json.Unmarshal([]byte(result), &f)
		if err != nil {
			fmt.Println("json err:", err)
		}
		m := f.(map[string]interface{})
		id := m["id"]
		//2.查询是否存在该id的用户
		var us service.UserService
		serverResponse:= us.UserInformation(id.(string))
		//3.编码json
		b, err := json.Marshal(serverResponse)
		if err != nil {
			fmt.Println("json err:", err)
		}
		fmt.Fprint(w, string(b))
	}
}

func HandleUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method=="GET"{
		//1.查询所有用户
		var us service.UserService
		serverResponse:= us.UsersInformation()
		//2.编码json
		b, err := json.Marshal(serverResponse)
		if err != nil {
			fmt.Println("json err:", err)
		}
		fmt.Fprint(w, string(b))
	}
}
