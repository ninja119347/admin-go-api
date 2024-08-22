// 用户服务层
package service

import (
	"admin-go-api/api/dao"
	"admin-go-api/api/dto"
	"admin-go-api/common/constant"
	"admin-go-api/common/result"
	"admin-go-api/pkg/jwt"
	"admin-go-api/pkg/log"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"time"
)

// 定义接口
type ISysAdminService interface {
	Login(c *gin.Context, dto dto.LoginDto)
	Send(c *gin.Context, dto dto.SendDto)
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
	log.Log().Info("shibai")
	//查看时间是否超时
	verifyTime := dto.IsUtraDataExpired()
	if verifyTime {
		fmt.Errorf("登录时间超时: %v", err)
		result.Failed(c, uint(result.ApiCode.LoginOutOfTime), result.ApiCode.GetMessage(result.ApiCode.LoginOutOfTime))
		return
	}
	//校验appName
	appName, _ := dao.CheckAppNameExists(dto.AppName)
	if !appName {
		result.Failed(c, uint(result.ApiCode.APPNAMEERROR), result.ApiCode.GetMessage(result.ApiCode.APPNAMEERROR))
		return
	}

	//生成token
	tokenString, expireTime, _ := jwt.GenerateTokenByAdmin(dto)
	result.Success(c, map[string]interface{}{"access_token": tokenString, "expire_at": expireTime, "max_refresh": constant.TokenMaxRefreshTime / time.Second, "timeout:": constant.TokenExpireDuration / time.Second})
	log.Log().Info("登录成功", dto, "token:", tokenString)
	//返回结果
	//dao.SysAdminDetail(dto)
}

// 实现发送邮件接口
func (s SysAdminServiceImpl) Send(c *gin.Context, dto dto.SendDto) {
	//参数校验
	err := validator.New().Struct(dto)
	if err != nil {
		fmt.Errorf("参数校验失败: %v", err)
		result.Failed(c, uint(result.ApiCode.ParamsFormError), result.ApiCode.GetMessage(result.ApiCode.ParamsFormError))
		return
	}
	//发送邮件
	//dao.SendEmail(dto)
	result.Success(c, map[string]interface{}{"code": "451196"})

}

func SysAdminService() ISysAdminService {
	return &sysAdminServiceImpl
}
