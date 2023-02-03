package mappers

import (
	"DuolinGo/model"
	"DuolinGo/utils"
	"fmt"
)

func Home() interface{} {
	var users []model.User
	result := utils.GetDB().Find(&users)
	err := result.Error
	num := result.RowsAffected
	if err != nil || num == 0 {
		return nil
	}
	return users
}
func Login(user model.User) interface{} {
	result := utils.GetDB().Where("userName = ? and passWord = ?", user.UserName, user.PassWord).Find(&user)
	err := result.Error
	num := result.RowsAffected
	// 获取用户错题数据
	errors := GetUserErrors(user.Uid)
	user.Errors = errors

	if err != nil || num == 0 {
		fmt.Println("User Not Found", err)
		return nil
	} else {
		fmt.Println("Login Success")
		return user
	}
}

func Register(user model.User) bool {

	if Login(user) == nil {
		result := utils.GetDB().Create(&user)
		num := result.RowsAffected
		err := result.Error
		if err != nil || num == 0 {
			fmt.Println("Register Fail", err)
			return false
		} else {
			fmt.Println("Register Success")
			return true
		}
	} else {
		return false
	}
}
