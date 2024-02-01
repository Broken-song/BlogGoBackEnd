package common

import (
	"bloggobackend/app/common/request"
	"bloggobackend/app/common/response"
	"bloggobackend/app/services"
	"github.com/gin-gonic/gin"
)

func ProfileSet(c *gin.Context) {
	var form request.ProfileSet
	if err := c.ShouldBind(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}
	err, _ := services.UserService.SetUserInfo(form)
	response.Success(c, err)
}
