package service

import (
	"internetOfThings/dao"
	"internetOfThings/config"
)

type UserService struct {
}

func (us UserService) Login(id string,password string) config.ServerResponse{
	var um dao.UserMapper
	resultCount := um.CheckUsername(id);
	if resultCount==0 {
		return config.CreateByErrorMessage("账号不存在")
	}
	user:= um.SelectLogin(id,password)
	if(user.Id == ""){
		return config.CreateByErrorMessage("密码错误")
	}else{
		return config.CreateBySuccessMessage1("登录成功")
	}
}

func (us UserService) UserInformation(id string) config.ServerResponse {
	var um dao.UserMapper
	user:=um.SelectById(id)
	if user.Id == "" {
		return config.CreateByErrorMessage("找不到当前用户")
	}else {
		return config.CreateBySuccessMessage2(user)
	}
}

func (us UserService) UsersInformation() config.ServerResponse {
	var um dao.UserMapper
	user:=um.SelectAll()
	return config.CreateBySuccessMessage2(user)
}

