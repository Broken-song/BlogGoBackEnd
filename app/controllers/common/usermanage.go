package common

import (
	"bloggobackend/app/common/request"
	"bloggobackend/app/common/response"
	"bloggobackend/app/services"
	"bloggobackend/global"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func UserManage(c *gin.Context) {
	var form request.UserManage
	if err := c.ShouldBind(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}
	data := services.UserManageService.GetUserByPage(form)
	c.JSON(200, data)
}

func ChangePassword(c *gin.Context) {
	var form request.ProfileSet
	if err := c.ShouldBind(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}
	err, _ := services.UserService.ChangeUserPassword(form)
	response.Success(c, err)
}

func UserDelete(c *gin.Context) {
	var form request.User
	if err := c.ShouldBind(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}
	tokenStr := c.Request.Header.Get("Authorization")
	if tokenStr == "" {
		response.TokenFail(c)
		c.Abort()
		return
	}
	tokenStr = tokenStr[len(services.TokenType)+1:]
	token, err := jwt.ParseWithClaims(tokenStr, &services.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.App.Config.Jwt.Secret), nil
	})
	if err != nil || services.JwtService.IsInBlacklist(tokenStr) {
		response.TokenFail(c)
		c.Abort()
		return
	}
	form.Token = token
	data, _ := services.UserService.DeleteUser(form)
	response.Success(c, data)
}
