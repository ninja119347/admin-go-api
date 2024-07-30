// jwt工具类（生成token以及获取当前登录用户id及用户信息）
package jwt

import (
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
const TokenExpireDuration = time.Hour * 2

// token密钥
var Secret = []byte("admin-go-api")

var (
	ErrAbsent  = "token absent" // token不存在
	ErrInValid = "token invalid"
)

// 根据用户信息生成token
func GenerateTokenByAdmin(admin entity.SysAdmin) (string, error) {
	var JwtAdmin entity.JwtAdmin
	JwtAdmin.ID = admin.ID
	JwtAdmin.Username = admin.Username
	JwtAdmin.Nickname = admin.Nickname
	JwtAdmin.Icon = admin.Icon
	JwtAdmin.Email = admin.Email
	JwtAdmin.Phone = admin.Phone
	JwtAdmin.Note = admin.Note
	claims := userStdClaims{
		JwtAdmin,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "admin",                                    //签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(Secret)
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

// 获取当前登录用户id
func GetAdminId(c *gin.Context) (uint, error) {
	u, exit := c.Get(constant.ContexkeyUserObj)
	if !exit {
		return 0, errors.New("can't get user id")
	}
	admin, ok := u.(*entity.JwtAdmin)
	if ok {
		return admin.ID, nil
	}
	return 0, errors.New("can't convert to id struct")
}

// 返回用户名
func GetAdminName(c *gin.Context) (string, error) {
	u, exit := c.Get(constant.ContexkeyUserObj)
	if !exit {
		return "0", errors.New("can't get user name")
	}
	admin, ok := u.(*entity.JwtAdmin)
	if ok {
		return admin.Username, nil
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
