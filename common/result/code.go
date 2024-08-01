// 状态码
package result

// codes定义状态
type Codes struct {
	SUCCESS         uint
	FAILED          uint
	Message         map[uint]string
	NOAUTH          uint
	AUTHFORM        uint
	ParamsFormError uint
	LoginCodeExpire uint
	CAPTCHANOTTRUE  uint
	PASSWORDNOTTRUE uint
	STATUSISENABLE  uint
	USERDISABLED    uint
	DatabaseError   uint
	INVALIDTOKEN    uint
	InvalidID       uint
	UserNotExist    uint
	DeleteUserError uint
	UpdateUserError uint
}

// ApiCode 状态码
var ApiCode = &Codes{
	SUCCESS:         200,
	FAILED:          501,
	NOAUTH:          403,
	AUTHFORM:        405,
	ParamsFormError: 407,
	LoginCodeExpire: 408,
	CAPTCHANOTTRUE:  409,
	PASSWORDNOTTRUE: 410,
	STATUSISENABLE:  411,
	USERDISABLED:    412,
	DatabaseError:   413,
	INVALIDTOKEN:    414,
	InvalidID:       415,
	UserNotExist:    416,
	DeleteUserError: 417,
	UpdateUserError: 418,
}

func init() {
	ApiCode.Message = map[uint]string{
		ApiCode.SUCCESS:         "成功",
		ApiCode.FAILED:          "失败",
		ApiCode.NOAUTH:          "请求头中token为空",
		ApiCode.AUTHFORM:        "请求头中token格式有误",
		ApiCode.ParamsFormError: "参数格式出错",
		ApiCode.LoginCodeExpire: "验证码已过期",
		ApiCode.CAPTCHANOTTRUE:  "验证码错误",
		ApiCode.PASSWORDNOTTRUE: "密码错误",
		ApiCode.STATUSISENABLE:  "状态启用",
		ApiCode.USERDISABLED:    "用户已禁用",
		ApiCode.DatabaseError:   "数据库错误",
		ApiCode.INVALIDTOKEN:    "无效的token",
		ApiCode.InvalidID:       "无效的ID",
		ApiCode.UserNotExist:    "用户不存在",
		ApiCode.DeleteUserError: "删除用户失败",
		ApiCode.UpdateUserError: "更新用户失败",
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
