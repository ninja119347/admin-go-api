// 访问接口路由配置
package router

import (
	"admin-go-api/api/controller"
	"admin-go-api/common/config"
	"admin-go-api/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

// 初始化路由
func InitRouter() *gin.Engine {
	// 初始化Gin路由器实例。
	router := gin.New()

	// 使用内置的Recovery中间件自动从任何panic中恢复，并向客户端返回500错误，
	// 确保服务器保持运行状态。
	router.Use(gin.Recovery())

	// 应用CORS中间件允许跨域请求，增强API从不同域或端口的可访问性。
	router.Use(middleware.Cors())

	// 从指定目录提供静态文件服务，允许直接访问上传的图片或其他静态内容。
	// 目录路径可以通过应用程序的配置进行配置。
	router.StaticFS(config.Config.ImageSettings.UploadDir, http.Dir(config.Config.ImageSettings.UploadDir))
	//log.Log().Info("uploadDir: ", config.Config.ImageSettings.UploadDir)
	// 集成自定义日志中间件以记录HTTP请求和响应，
	// 帮助监控和调试应用程序。
	router.Use(middleware.Logger())
	register(router)
	return router
}

// 路由注册
func register(router *gin.Engine) {
	// 需要 JWT 验证的路由组
	auth := router.Group("/authJwt")
	auth.Use(middleware.AuthMiddleware())
	auth.POST("/api/m/core/email", controller.Send)
	// 其他需要 JWT 验证的接口
	//auth.POST("/register", controller.AddUser)
	//auth.DELETE("/delete/user/:id", controller.DeleteUser)

	//router.POST("/register", controller.AddUser)
	router.GET("/api/captcha", controller.Captcha)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST("/api/login", controller.Login)
	//router.PUT("/api/updateUser", controller.UpdateUser)
	router.GET("/ping1", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//router.POST("/api/register", controller.AddUser)

}
