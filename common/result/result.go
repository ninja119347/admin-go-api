// 通用访问结构
package result

import (
	"admin-go-api/pkg/log"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 消息结构体
type Result struct {
	Code int         `json:"code"` //状态码
	Msg  string      `json:"msg"`  //提示信息
	Data interface{} `json:"data"` //返回的数据
}

// 返回成功
func Success(c *gin.Context, data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	res := Result{}
	res.Code = int(ApiCode.SUCCESS)
	res.Msg = ApiCode.GetMessage(ApiCode.SUCCESS)
	res.Data = data
	log.Log().Info("返回数据: ", res)
	c.JSON(http.StatusOK, res)
}

// 返回失败
func Failed(c *gin.Context, code uint, msg string) {
	res := Result{}
	res.Code = int(code)
	res.Msg = msg
	res.Data = gin.H{}
	c.JSON(http.StatusOK, res)
}
