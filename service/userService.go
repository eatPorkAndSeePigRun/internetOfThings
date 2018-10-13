package service

import (
	"internetOfThings/dao"
	"internetOfThings/config"
)

type UserService struct {
}

func (us UserService) Login(username string, password string) config.ServerResponse {
	var um dao.UserMapper
	resultCount := um.CheckUsername(username);
	if resultCount == 0 {
		return config.CreateByErrorMessage("账号不存在")
	}
	user := um.SelectLogin(username, password)
	if (user.Username == "") {
		return config.CreateByErrorMessage("密码错误")
	} else {
		temp := struct {
			Username string `json:"username"`
		}{user.Username}
		return config.CreateBySuccessMessage3("OK", temp)
	}
}

func (us UserService) UserInformation(id string) config.ServerResponse {
	var um dao.UserMapper
	user := um.SelectById(id)
	if user.Id == "" {
		return config.CreateByErrorMessage("找不到当前用户")
	} else {
		return config.CreateBySuccessMessage2(user)
	}
}

func (us UserService) UsersInformation() config.ServerResponse {
	var um dao.UserMapper
	user := um.SelectAll()
	return config.CreateBySuccessMessage2(user)
}

func (us UserService) ChangePW(username string, newpassword string) config.ServerResponse {
	var um dao.UserMapper
	resultCount := um.CheckUsername(username);
	if resultCount == 0 {
		return config.CreateByErrorMessage("没有该用户")
	}
	um.ChangePW(username, newpassword)
	temp := struct {
		Username string `json:"username"`
	}{username}
	return config.CreateBySuccessMessage3("密码更改成功", temp)

}
