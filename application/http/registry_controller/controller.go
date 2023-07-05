package registry_controller

import (
	"github.com/LeeZXin/z-base/common"
	"github.com/LeeZXin/z-base/service/registry_service"
	"github.com/LeeZXin/zsf/validator"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterService(c *gin.Context) {
	var reqVO RegisterServiceReqVO
	err := c.ShouldBind(&reqVO)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.DefaultFailBindArgResp)
		return
	}
	err = validator.Validate(reqVO)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.NewBaseResp(common.InvalidParamsCode, err.Error()))
		return
	}
	registry_service.ServiceImpl.RegisterService(c.Request.Context(), registry_service.RegisterServiceReqDTO{
		ServiceName:   reqVO.ServiceName,
		Ip:            reqVO.Ip,
		Port:          reqVO.Port,
		InstanceId:    reqVO.InstanceId,
		Weight:        reqVO.Weight,
		Version:       reqVO.Version,
		LeaseDuration: reqVO.LeaseDuration,
	})
	c.JSON(http.StatusOK, common.DefaultSuccessResp)
}

func DeregisterService(c *gin.Context) {
	var reqVO DeregisterServiceReqVO
	err := c.ShouldBind(&reqVO)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.DefaultFailBindArgResp)
		return
	}
	err = validator.Validate(reqVO)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.NewBaseResp(common.InvalidParamsCode, err.Error()))
		return
	}
	err = registry_service.ServiceImpl.DeregisterService(c.Request.Context(), registry_service.DeregisterServiceReqDTO{
		ServiceName: reqVO.ServiceName,
		InstanceId:  reqVO.InstanceId,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, common.NewBaseResp(common.DataNotExistsCode, err.Error()))
		return
	}
	c.JSON(http.StatusOK, common.DefaultSuccessResp)
}

func PassTTL(c *gin.Context) {
	var reqVO PassTtlReqVO
	err := c.ShouldBind(&reqVO)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.DefaultFailBindArgResp)
		return
	}
	err = validator.Validate(reqVO)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.NewBaseResp(common.InvalidParamsCode, err.Error()))
		return
	}
	err = registry_service.ServiceImpl.PassTTL(c.Request.Context(), registry_service.PassTtlReqDTO{
		ServiceName: reqVO.ServiceName,
		InstanceId:  reqVO.InstanceId,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, common.NewBaseResp(common.DataNotExistsCode, err.Error()))
		return
	}
	c.JSON(http.StatusOK, common.DefaultSuccessResp)
}

func GetServiceInfoList(c *gin.Context) {
	var reqVO GetServiceInfoListReqVO
	err := c.ShouldBind(&reqVO)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.DefaultFailBindArgResp)
		return
	}
	err = validator.Validate(reqVO)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.NewBaseResp(common.InvalidParamsCode, err.Error()))
		return
	}
	serviceInfoList := registry_service.ServiceImpl.GetServiceInfoList(c.Request.Context(), reqVO.ServiceName)
	infoVoList := make([]ServiceInfoVO, 0, len(serviceInfoList))
	for _, info := range serviceInfoList {
		infoVoList = append(infoVoList, ServiceInfoVO{
			ServiceName:   info.ServiceName,
			Ip:            info.Ip,
			Port:          info.Port,
			InstanceId:    info.InstanceId,
			Weight:        info.Weight,
			Version:       info.Version,
			LeaseDuration: int(info.LeaseDuration.Seconds()),
		})
	}
	c.JSON(http.StatusOK, GetServiceInfoListRespVO{
		BaseResp:    common.DefaultSuccessResp,
		ServiceList: infoVoList,
	})
}
