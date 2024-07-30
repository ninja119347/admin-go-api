// 验证码 服务层
package service

import (
	"admin-go-api/common/util"
	"github.com/mojocn/base64Captcha"
	"image/color"
)

var store = util.RedisStore{}

// 生成验证码
func CaptMake() (id, ba64s string) {
	var driver base64Captcha.Driver
	var driverString base64Captcha.DriverString
	//配置验证码信息
	captcharConfig := base64Captcha.DriverString{
		Height:          60,
		Width:           200,
		NoiseCount:      0,
		ShowLineOptions: 2 | 4,
		Length:          6,
		Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm",
		BgColor: &color.RGBA{
			R: 3,
			G: 102,
			B: 214,
			A: 125,
		},
		Fonts: []string{"wqy-microhei.ttc"},
	}
	driverString = captcharConfig
	driver = driverString.ConvertFonts()
	captcha := base64Captcha.NewCaptcha(driver, &store)
	lid, lb64s, _, _ := captcha.Generate()
	return lid, lb64s
}

// 验证验证码
func CaptVerify(id, answer string) bool {
	return store.Verify(id, answer, false)
}
