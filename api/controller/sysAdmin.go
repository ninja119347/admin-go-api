// 用户控制层
package controller

import (
	"admin-go-api/api/dto"
	"admin-go-api/api/service"
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

// @Summary 新增用户接口
// @Description 新增用户接口
// @Produce json
// @Param data body dto.CreateUserDto true "data"
// @Success 200 {object} result.Result
// @Router /api/register [post]
func AddUser(c *gin.Context) {
	var dto dto.CreateUserDto
	_ = c.BindJSON(&dto)
	service.SysAdminService().CreateUser(c, dto)
}
