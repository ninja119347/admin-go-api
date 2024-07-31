// 状态码
package result

// codes定义状态
type Codes struct {
	SUCCESS            uint
	FAILED             uint
	Message            map[uint]string
	NOAUTH             uint
	AUTHFORM           uint
	MissingLoginParams uint
	LoginCodeExpire    uint
	CAPTCHANOTTRUE     uint
	PASSWORDNOTTRUE    uint
	STATUSISENABLE     uint
	USERDISABLED       uint
	DatabaseError      uint
	INVALIDTOKEN       uint
}

// ApiCode 状态码
var ApiCode = &Codes{
	SUCCESS:            200,
	FAILED:             501,
	NOAUTH:             403,
	AUTHFORM:           405,
	MissingLoginParams: 407,
	LoginCodeExpire:    408,
	CAPTCHANOTTRUE:     409,
	PASSWORDNOTTRUE:    410,
	STATUSISENABLE:     411,
	USERDISABLED:       412,
	DatabaseError:      413,
	INVALIDTOKEN:       414,
}

// 状态信息
func init() {
	ApiCode.Message = map[uint]string{
		ApiCode.SUCCESS:            "成功",
		ApiCode.FAILED:             "失败",
		ApiCode.NOAUTH:             "请求头中token为空",
		ApiCode.AUTHFORM:           "请求头中token格式有误",
		ApiCode.MissingLoginParams: "缺少登录参数",
		ApiCode.LoginCodeExpire:    "验证码已过期",
		ApiCode.CAPTCHANOTTRUE:     "验证码错误",
		ApiCode.PASSWORDNOTTRUE:    "密码错误",
		ApiCode.STATUSISENABLE:     "状态启用",
		ApiCode.USERDISABLED:       "用户已禁用",
		ApiCode.DatabaseError:      "数据库错误",
		ApiCode.INVALIDTOKEN:       "无效的token",
	}
}

// 供外部调用
func (c *Codes) GetMessage(code uint) string {
	msg, ok := c.Message[code]
	if ok {
		return msg
	}
	return c.Message[c.FAILED]
}
