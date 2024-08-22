// 系统常量
package constant

import "time"

const (

	//存登录用户
	ContexkeyUserObj = "autheUserObj"

	//验证码
	LOGIN_CODE = "login_code:"
	// token过期时间
	TokenExpireDuration = time.Minute * 8
	// token最大刷新时间
	TokenMaxRefreshTime = time.Hour * 2

	//用户状态
	USER_STATUS_DISABLE = 2
	USER_STATUS_ENABLE  = 1
)
