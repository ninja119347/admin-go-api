// 用户控制层
package controller

import (
	"admin-go-api/api/dto"
	"admin-go-api/api/service"
	"admin-go-api/common/result"
	"admin-go-api/pkg/log"
	"github.com/gin-gonic/gin"
	"strconv"
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

// @Summary 删除用户接口
// @Description 删除用户接口
// @Produce json
// @Param id path int true "用户ID"
// @Success 200 {object} result.Result
// @Router /api/user/{id} [delete]
func DeleteUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	log.Log().Infof("id: %d, idParam: %v", id, idParam)
	if err != nil {
		result.Failed(c, uint(result.ApiCode.InvalidID), result.ApiCode.GetMessage(result.ApiCode.InvalidID))
		return
	}
	service.SysAdminService().DeleteUser(c, uint(id))
}

// @Summary 查询用户接口
// @Description 查询用户接口
// @Produce json
// @Param username query string true "用户名"
// @Success 200 {object} result.Result
// @Router /api/user [get]
func SearchUser(c *gin.Context) {
	username := c.Param("username")
	service.SysAdminService().SearchUser(c, username)
}

// @Summary 更新用户接口
// @Description 更新用户接口
// @Produce json
// @Param id path int true "用户ID"
// @Param data body dto.UpdateUserDto true "data"
// @Success 200 {object} result.Result
// @Router /api/user/{id} [put]
func UpdateUser(c *gin.Context) {
	var dto dto.UpdateUserDto
	_ = c.BindJSON(&dto)
	service.SysAdminService().UpdateUser(c, dto)
}

// @Summary 查询所有用户接口
// @Description 查询所有用户接口
// @Produce json
// @Success 200 {object} result.Result
// @Router /api/user [get]
func SearchUserAll(c *gin.Context) {
	service.SysAdminService().SearchUserAll(c)
}
