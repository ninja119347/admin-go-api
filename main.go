// 启动程序
package main

import (
	"admin-go-api/common/config"
	_ "admin-go-api/docs"
	"admin-go-api/pkg/db"
	"admin-go-api/pkg/log"
	"admin-go-api/pkg/redis"
	"admin-go-api/router"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// @title admin-go-api
// @version 1.0
// @description 后台管理系统API接口文档
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	log := log.Log()
	gin.SetMode(config.Config.Server.Model)
	router := router.InitRouter()
	srv := &http.Server{
		Addr:    config.Config.Server.Address,
		Handler: router,
	}
	// 启动服务(固定写法)
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Info("listen: %s\n", err)
		}
		log.Info("listen on " + config.Config.Server.Address)
	}()
	quit := make(chan os.Signal)
	//监听消息
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Info("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Info("Server Shutdown:", err)
	}
	log.Info("Server exiting")
}

// 初始化连接
func init() {
	//mysql
	db.SetupDBLink()
	//redis
	redis.SetupRedisDB()
}
