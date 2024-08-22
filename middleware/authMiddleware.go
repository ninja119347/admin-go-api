// 鉴权中间件
package middleware

import (
	"admin-go-api/common/constant"
	"admin-go-api/common/result"
	"admin-go-api/pkg/jwt"
	"github.com/gin-gonic/gin"
	"strings"
)

func AuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			result.Failed(c, result.ApiCode.FAILED, result.ApiCode.GetMessage(result.ApiCode.NOAUTH))
			c.Abort()
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			result.Failed(c, result.ApiCode.FAILED, result.ApiCode.GetMessage(result.ApiCode.AUTHFORM))
			c.Abort()
			return
		}
		// 验证 token
		claims, err := jwt.ValidateToken(parts[1])
		if err != nil {
			result.Failed(c, result.ApiCode.FAILED, result.ApiCode.GetMessage(result.ApiCode.INVALIDTOKEN))
			c.Abort()
			return
		}
		// 将用户信息存储在上下文中
		c.Set(constant.ContexkeyUserObj, claims)
		c.Next()
	}
}
