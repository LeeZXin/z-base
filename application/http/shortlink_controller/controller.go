package shortlink_controller

import (
	"github.com/LeeZXin/z-base/common"
	"github.com/LeeZXin/z-base/service/shortlink_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 短链服务controller

// CreateShortLink 生成短链
func CreateShortLink(c *gin.Context) {
	var reqVO CreateShortLinkReqVO
	err := c.ShouldBind(&reqVO)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.DefaultFailBindArgResp)
		return
	}
	reqDTO := shortlink_service.CreateShortLinkReqDTO{
		LongLink: reqVO.LongLink,
	}
	respDTO := shortlink_service.ServiceImpl.CreateShortLink(c.Request.Context(), reqDTO)
	respVO := CreateShortLinkRespVO{
		BaseResp:  respDTO.BaseResp,
		ShortLink: respDTO.ShortLink,
	}
	c.JSON(http.StatusOK, respVO)
}

// GetLongLinkByShortLink 获取长链
func GetLongLinkByShortLink(c *gin.Context) {
	var reqVO GetLongLinkReqVO
	err := c.ShouldBind(&reqVO)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.DefaultFailBindArgResp)
		return
	}
	reqDTO := shortlink_service.GetLongLinkReqDTO{
		ShortLink: reqVO.ShortLink,
	}
	respDTO := shortlink_service.ServiceImpl.GetLongLink(c.Request.Context(), reqDTO)
	respVO := GetLongLinkRespVO{
		BaseResp: respDTO.BaseResp,
		LongLink: respDTO.LongLink,
	}
	c.JSON(http.StatusOK, respVO)
}
