package registry_controller

import (
	"github.com/LeeZXin/z-base/common"
)

type RegisterServiceReqVO struct {
	ServiceName   string `json:"serviceName" validate:"nonzero"`
	Ip            string `json:"ip" validate:"nonzero"`
	Port          int    `json:"port" validate:"nonzero"`
	InstanceId    string `json:"instanceId" validate:"nonzero"`
	Weight        int    `json:"weight" validate:"nonzero"`
	Version       string `json:"version" validate:"nonzero"`
	LeaseDuration int    `json:"leaseDuration" validate:"nonzero"`
}

type DeregisterServiceReqVO struct {
	ServiceName string `json:"serviceName"  validate:"nonzero"`
	InstanceId  string `json:"instanceId"  validate:"nonzero"`
}

type PassTtlReqVO struct {
	ServiceName string `json:"serviceName"  validate:"nonzero"`
	InstanceId  string `json:"instanceId"  validate:"nonzero"`
}

type GetServiceInfoListReqVO struct {
	ServiceName string `json:"serviceName"  validate:"nonzero"`
}

type GetServiceInfoListRespVO struct {
	common.BaseResp
	ServiceList []ServiceInfoVO `json:"serviceList"`
}

type ServiceInfoVO struct {
	ServiceName   string `json:"serviceName"`
	Ip            string `json:"ip"`
	Port          int    `json:"port"`
	InstanceId    string `json:"instanceId"`
	Weight        int    `json:"weight"`
	Version       string `json:"version"`
	LeaseDuration int    `json:"leaseDuration"`
}
