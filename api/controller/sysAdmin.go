// 用户控制层
package controller

import (
	"admin-go-api/api/dto"
	"admin-go-api/api/service"
	"admin-go-api/pkg/log"
	"github.com/gin-gonic/gin"
)

// @Summary 用户登录接口
// @Description 用户登录接口
// @Produce json
// @Param data body dto.LoginDto true "data"
// @Success 200 {object} result.Result
// @Router /api/login [post]
func Login(c *gin.Context) {
	var dto dto.LoginDto
	//绑定参数将HTTP request中的json参数绑定到dto中
	_ = c.BindJSON(&dto)
	service.SysAdminService().Login(c, dto)
}

// @Summary 发送邮件接口
// @Description 发送邮件接口
// @Produce json
// @Param data body dto.SendDto true "data"
// @Success 200 {object} result.Result
// @Router /api/m/core/email [post]
func Send(c *gin.Context) {
	var dto dto.SendDto
	//绑定参数将HTTP request中的json参数绑定到dto中
	_ = c.BindJSON(&dto)
	log.Log().Info("dto: ", dto)
	service.SysAdminService().Send(c, dto)
}
