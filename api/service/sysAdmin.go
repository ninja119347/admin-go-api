// 用户服务层
package service

import (
	"admin-go-api/api/dao"
	"admin-go-api/api/dto"
	"admin-go-api/common/constant"
	"admin-go-api/common/result"
	"admin-go-api/common/util"
	"admin-go-api/pkg/jwt"
	"admin-go-api/pkg/log"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// 定义接口
type ISysAdminService interface {
	Login(c *gin.Context, dto dto.LoginDto)
	CreateUser(c *gin.Context, dto dto.CreateUserDto)
	DeleteUser(c *gin.Context, id uint)
	SearchUser(c *gin.Context, username string)
	UpdateUser(c *gin.Context, dto dto.UpdateUserDto)
	SearchUserList(c *gin.Context, page, pageSize int)
}

type SysAdminServiceImpl struct{}

var sysAdminServiceImpl = SysAdminServiceImpl{}

// 实现用户登录接口
func (s SysAdminServiceImpl) Login(c *gin.Context, dto dto.LoginDto) {
	//参数校验
	err := validator.New().Struct(dto)
	if err != nil {
		fmt.Errorf("参数校验失败: %v", err)
		result.Failed(c, uint(result.ApiCode.ParamsFormError), result.ApiCode.GetMessage(result.ApiCode.ParamsFormError))
		return
	}
	//查看验证码是否过期
	code := util.RedisStore{}.Get(dto.IdKey, true)
	if len(code) == 0 {
		result.Failed(c, uint(result.ApiCode.LoginCodeExpire), result.ApiCode.GetMessage(result.ApiCode.LoginCodeExpire))
		return
	}
	//校验验证码
	verifyRes := CaptVerify(dto.IdKey, dto.Image)
	if !verifyRes {
		result.Failed(c, uint(result.ApiCode.CAPTCHANOTTRUE), result.ApiCode.GetMessage(result.ApiCode.CAPTCHANOTTRUE))
		return
	}
	//校验用户
	sysAdmin := dao.SysAdminDetail(dto)
	if sysAdmin.Password != util.EncryptionMd5(dto.Password) {
		result.Failed(c, uint(result.ApiCode.PASSWORDNOTTRUE), result.ApiCode.GetMessage(result.ApiCode.PASSWORDNOTTRUE))
		return
	}
	if sysAdmin.Status == constant.USER_STATUS_DISABLE {
		result.Failed(c, uint(result.ApiCode.USERDISABLED), result.ApiCode.GetMessage(result.ApiCode.USERDISABLED))
		return
	}
	//生成token
	tokenString, _ := jwt.GenerateTokenByAdmin(sysAdmin)
	result.Success(c, map[string]interface{}{"token": tokenString, "sysAdmin": sysAdmin})
	log.Log().Info("登录成功", sysAdmin, "token:", tokenString)
	//返回结果
	//dao.SysAdminDetail(dto)
}

// 实现新增用户接口
func (s SysAdminServiceImpl) CreateUser(c *gin.Context, dto dto.CreateUserDto) {
	// 参数校验
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c, uint(result.ApiCode.ParamsFormError), result.ApiCode.GetMessage(result.ApiCode.ParamsFormError))
		return
	}
	// 调用DAO层保存用户
	if err := dao.CreateUser(dto); err != nil {
		result.Failed(c, uint(result.ApiCode.DatabaseError), result.ApiCode.GetMessage(result.ApiCode.DatabaseError))
		return
	}
	result.Success(c, map[string]interface{}{"data": dto})
}

// 实现删除用户接口
func (s SysAdminServiceImpl) DeleteUser(c *gin.Context, id uint) {
	// 调用DAO层删除用户
	if err := dao.DeleteUser(id); err != nil {
		result.Failed(c, uint(result.ApiCode.DeleteUserError), result.ApiCode.GetMessage(result.ApiCode.DeleteUserError))
		return
	}
	result.Success(c, nil)

}

// 实现查询用户接口
func (s SysAdminServiceImpl) SearchUser(c *gin.Context, username string) {
	// 调用DAO层查询用户
	user, err := dao.SearchUser(username)
	if err != nil {
		result.Failed(c, uint(result.ApiCode.UserNotExist), result.ApiCode.GetMessage(result.ApiCode.UserNotExist))
		return
	}
	result.Success(c, map[string]interface{}{"data": user})
}

// 实现更新用户接口
func (s SysAdminServiceImpl) UpdateUser(c *gin.Context, dto dto.UpdateUserDto) {
	// 参数校验
	errParam := validator.New().Struct(dto)
	if errParam != nil {
		result.Failed(c, uint(result.ApiCode.ParamsFormError), result.ApiCode.GetMessage(result.ApiCode.ParamsFormError))
		return
	}
	//查询用户是否存在
	_, err := dao.SearchUserById(dto.ID)
	if err != nil {
		result.Failed(c, uint(result.ApiCode.UserNotExist), result.ApiCode.GetMessage(result.ApiCode.UserNotExist))
		return
	}
	// 调用DAO层更新用户
	if err := dao.UpdateUser(dto); err != nil {
		result.Failed(c, uint(result.ApiCode.UpdateUserError), result.ApiCode.GetMessage(result.ApiCode.UpdateUserError))
		return
	}
	result.Success(c, map[string]interface{}{"data": dto})

}

// 实现查询分页查询用户接口
func (s SysAdminServiceImpl) SearchUserList(c *gin.Context, page, pageSize int) {
	//检验参数
	log.Log().Info("page: ", page, "pageSize: ", pageSize)
	if page < 0 || pageSize < 0 {
		result.Failed(c, uint(result.ApiCode.ParamsFormError), result.ApiCode.GetMessage(result.ApiCode.ParamsFormError))
		return
	}
	if pageSize == 0 {
		pageSize = -1
	}
	if page == 0 {
		page = -1
	}
	// 调用DAO层查询用户
	users, err := dao.SearchUserList(page, pageSize)
	if err != nil {
		result.Failed(c, uint(result.ApiCode.UserNotExist), result.ApiCode.GetMessage(result.ApiCode.UserNotExist))
		return
	}
	result.Success(c, map[string]interface{}{"data": users})
}

func SysAdminService() ISysAdminService {
	return &sysAdminServiceImpl
}
