package services

import (
	"bloggobackend/app/common/request"
	"bloggobackend/app/models"
	"bloggobackend/global"
	"bloggobackend/utils"
	"errors"
	"strconv"
)

type userService struct {
}

var UserService = new(userService)

// Register 注册
func (userService *userService) Register(params request.Register) (err error, user models.User) {
	var result = global.App.DB.Where("email = ?", params.Email).Select("id").First(&models.User{})
	if result.RowsAffected != 0 {
		err = errors.New("邮箱已存在")
		return
	}
	user = models.User{Account: params.Account, Email: params.Email, Password: utils.BcryptMake([]byte(params.Password))}
	err = global.App.DB.Create(&user).Error
	return
}

// Login 登录
func (userService *userService) Login(params request.Login) (err error, user *models.User) {
	err = global.App.DB.Where("email = ?", params.Email).First(&user).Error
	if err != nil || !utils.BcryptMakeCheck([]byte(params.Password), user.Password) {
		err = errors.New("用户不存在或密码错误")
	}
	return
}
func (userService *userService) GetUserInfo(id string) (err error, user models.User) {
	intId, err := strconv.Atoi(id)
	err = global.App.DB.First(&user, intId).Error
	if err != nil {
		err = errors.New("数据不存在")
	}
	return
}
