package main

import (
	"github.com/LeeZXin/z-base/application/http/registry_controller"
	"github.com/LeeZXin/z-base/application/http/shortlink_controller"
	"github.com/LeeZXin/z-base/application/http/sid_controller"
	"github.com/LeeZXin/zsf/zsf"
	"github.com/gin-gonic/gin"
)

func main() {
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
		registryGroup := e.Group("/registry")
		{
			// 注册服务
			registryGroup.POST("/registerService", registry_controller.RegisterService)
			// 注销服务
			registryGroup.POST("/deregisterService", registry_controller.DeregisterService)
			// 续期
			registryGroup.POST("/passTTL", registry_controller.PassTTL)
			// 获取服务列表
			registryGroup.POST("/getServiceInfoList", registry_controller.GetServiceInfoList)
		}
	})
	zsf.Run()
}
