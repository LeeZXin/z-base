package main

import (
	"github.com/LeeZXin/z-base/application/http/shortlink"
	"github.com/LeeZXin/z-base/application/http/sid"
	"github.com/LeeZXin/zsf/zsf"
	"github.com/gin-gonic/gin"
)

func main() {
	zsf.RegisterHttpRouter(func(e *gin.Engine) {
		sidGroup := e.Group("/sid")
		{
			// 生成雪花id
			sidGroup.POST("/generateIds", sid.GenerateIds)
		}
		shortLinkGroup := e.Group("/shortLink")
		{
			// 生成短链
			shortLinkGroup.POST("/createShortLink", shortlink.CreateShortLink)
			// 获取长链
			shortLinkGroup.POST("/getLongLink", shortlink.GetLongLinkByShortLink)
		}
	})
	zsf.Run()
}
