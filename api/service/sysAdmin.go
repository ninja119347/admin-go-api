// 用户服务层
package service

import (
	"admin-go-api/api/dao"
	"admin-go-api/api/entity"
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
	Login(c *gin.Context, dto entity.LoginDto)
}

type SysAdminServiceImpl struct{}

var sysAdminService = SysAdminServiceImpl{}

// 实现用户登录接口
func (s SysAdminServiceImpl) Login(c *gin.Context, dto entity.LoginDto) {
	//参数校验
	err := validator.New().Struct(dto)
	if err != nil {
		fmt.Errorf("参数校验失败: %v", err)
		result.Failed(c, uint(result.ApiCode.MissingLoginParams), result.ApiCode.GetMessage(result.ApiCode.MissingLoginParams))
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
	if sysAdmin.Password != dto.Password {
		log.Log().Errorf("密码错误, 密码: %s, 输入密码：%s", sysAdmin.Password, dto.Password)
		result.Failed(c, uint(result.ApiCode.PASSWORDNOTTRUE), result.ApiCode.GetMessage(result.ApiCode.PASSWORDNOTTRUE))
		return
	}
	if sysAdmin.Status == constant.USER_STATUS_DISABLE {
		result.Failed(c, uint(result.ApiCode.USERDISABLED), result.ApiCode.GetMessage(result.ApiCode.USERDISABLED))
		return
	}
	//生成token
	tokenString, _ := jwt.GenerateTokenByAdmin(sysAdmin)
	//admin_In := admin_info{
	//	Username: sysAdmin.Username,
	//	Password: sysAdmin.Password,
	//}
	result.Success(c, map[string]interface{}{"token": tokenString, "sysAdmin": sysAdmin})
	log.Log().Info("登录成功", sysAdmin)
	//返回结果
	//dao.SysAdminDetail(dto)
}
func SysAdminService() ISysAdminService {
	return &sysAdminService
}

type admin_info struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
