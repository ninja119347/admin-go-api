// jwt工具类（生成token以及获取当前登录用户id及用户信息）
package jwt

import (
	"admin-go-api/api/dto"
	"admin-go-api/api/entity"
	"admin-go-api/common/constant"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

type userStdClaims struct {
	entity.JwtAdmin
	jwt.StandardClaims
}

// token过期时间
//const TokenExpireDuration = time.Hour * 2

// token密钥
var Secret = []byte("admin-go-api")

var (
	ErrAbsent  = "token absent" // token不存在
	ErrInValid = "token invalid"
)

// 根据用户信息生成token
func GenerateTokenByAdmin(admin dto.LoginDto) (string, int64, error) {
	var JwtAdmin entity.JwtAdmin
	JwtAdmin.Id = admin.Id
	JwtAdmin.App_name = admin.AppName
	JwtAdmin.Id_type = admin.IdType
	claims := userStdClaims{
		JwtAdmin,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(constant.TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "admin",                                             //签发人
			IssuedAt:  time.Now().Unix(),                                   // 签发时间
			NotBefore: time.Now().Unix(),                                   // 生效时间
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(Secret)
	return tokenString, claims.ExpiresAt, err
}

// 解析JWT(固定写法)
func ValidateToken(tokenString string) (*entity.JwtAdmin, error) {
	if tokenString == "" {
		return nil, errors.New(ErrAbsent)
	}
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	if token == nil {
		return nil, errors.New(ErrInValid)
	}
	claims := userStdClaims{}
	_, err := jwt.ParseWithClaims(tokenString, &userStdClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return Secret, nil
	})
	if err != nil {
		return nil, err
	}
	return &claims.JwtAdmin, nil
}

// 获取当前登录appid
func GetAdminId(c *gin.Context) (string, error) {
	u, exit := c.Get(constant.ContexkeyUserObj)
	if !exit {
		return "", errors.New("can't get user id")
	}
	admin, ok := u.(*entity.JwtAdmin)
	if ok {
		return admin.Id, nil
	}
	return "", errors.New("can't convert to id struct")
}

// 返回app_name
func GetAdminName(c *gin.Context) (string, error) {
	u, exit := c.Get(constant.ContexkeyUserObj)
	if !exit {
		return "0", errors.New("can't get app name")
	}
	admin, ok := u.(*entity.JwtAdmin)
	if ok {
		return admin.App_name, nil
	}
	return "0", errors.New("can't convert to name struct")
}

// 返回admin信息
func GetAdmin(c *gin.Context) (*entity.JwtAdmin, error) {
	u, exit := c.Get(constant.ContexkeyUserObj)
	if !exit {
		return nil, errors.New("can't get user")
	}
	admin, ok := u.(*entity.JwtAdmin)
	if ok {
		return admin, nil
	}
	return nil, errors.New("can't convert to admin struct")
}
