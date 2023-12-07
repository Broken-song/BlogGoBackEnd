package common

import (
	"bloggobackend/app/common/request"
	"bloggobackend/app/common/response"
	"bloggobackend/app/services"
	"github.com/gin-gonic/gin"
)

func DataGet(c *gin.Context) {
	var form request.DataGet
	if err := c.ShouldBind(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}
	data := services.DataService.GetDataByPage(form)
	response.Success(c, data)
}
