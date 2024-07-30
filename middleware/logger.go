// 日志log中间件
package middleware

import (
	"admin-go-api/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"time"
)

func Logger() gin.HandlerFunc {
	logger := log.Log()
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime) / time.Microsecond
		regMethod := c.Request.Method
		regUri := c.Request.RequestURI
		header := c.Request.Header
		proto := c.Request.Proto
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		err := c.Err()
		body, _ := io.ReadAll(c.Request.Body)
		//log.Log().Info("body: ", body)
		logger.WithFields(logrus.Fields{
			"status_code":      statusCode,
			"latency_time(μs)": latencyTime,
			"client_ip":        clientIP,
			"req_method":       regMethod,
			"req_uri":          regUri,
			"req_proto":        proto,
			"req_header":       header,
			"req_body":         string(body),
			"err":              err,
		}).Info()

	}
}
