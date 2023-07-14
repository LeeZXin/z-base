package main

import (
	"github.com/LeeZXin/z-base/application/http/shortlink_controller"
	"github.com/LeeZXin/z-base/application/http/sid_controller"
	"github.com/LeeZXin/zsf/property"
	"github.com/LeeZXin/zsf/quit"
	"github.com/LeeZXin/zsf/sa_registry/server"
	"github.com/LeeZXin/zsf/zsf"
	"github.com/gin-gonic/gin"
)

func main() {
	// 单机版注册中心
	registryServer := server.NewRegistryServer(0, property.GetString("saRegistry.token"), nil)
	zsf.RegisterHttpRouter(func(e *gin.Engine) {
		sidGroup := e.Group("/sid")
		{
			// 生成雪花id
			sidGroup.POST("/generateIds", sid_controller.GenerateIds)
		}
		shortLinkGroup := e.Group("/shortLink")
		{
			// 生成短链
			shortLinkGroup.POST("/createShortLink", shortlink_controller.CreateShortLink)
			// 获取长链
			shortLinkGroup.POST("/getLongLink", shortlink_controller.GetLongLinkByShortLink)
		}
		registryServer.HttpRouter(e)
	})
	registryServer.Start(false)
	quit.AddShutdownHook(func() {
		registryServer.Stop()
	})
	zsf.Run()
}
