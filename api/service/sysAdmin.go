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
}

type SysAdminServiceImpl struct{}

var sysAdminServiceImpl = SysAdminServiceImpl{}

// 实现用户登录接口
func (s SysAdminServiceImpl) Login(c *gin.Context, dto dto.LoginDto) {
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
	if sysAdmin.Password != util.EncryptionMd5(dto.Password) {
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
		result.Failed(c, uint(result.ApiCode.MissingLoginParams), "参数校验失败")
		return
	}
	// 调用DAO层保存用户
	if err := dao.CreateUser(dto); err != nil {
		result.Failed(c, uint(result.ApiCode.DatabaseError), "数据库错误")
		return
	}
	result.Success(c, map[string]interface{}{"data": dto})
}

func SysAdminService() ISysAdminService {
	return &sysAdminServiceImpl
}

//type admin_info struct {
//	Username string `json:"username"`
//	Password string `json:"password"`
//}
