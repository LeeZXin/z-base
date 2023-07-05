package sid_controller

import (
	"github.com/LeeZXin/z-base/common"
	"github.com/LeeZXin/z-base/service/sid_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 雪花id服务 controller

// GenerateIds 按批生成雪花id
func GenerateIds(c *gin.Context) {
	var reqVO GenerateIdsReqVO
	err := c.ShouldBind(&reqVO)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.DefaultFailBindArgResp)
		return
	}
	reqDTO := sid_service.GenerateIdsReqDTO{
		BatchNum: reqVO.BatchNum,
	}
	respDTO := sid_service.ServiceImpl.GenerateIds(reqDTO)
	respVO := GenerateIdsRespVO{
		BaseResp: respDTO.BaseResp,
		Ids:      respDTO.Result,
	}
	c.JSON(http.StatusOK, respVO)
}
