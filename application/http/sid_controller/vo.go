package sid_controller

import "github.com/LeeZXin/z-base/common"

type GenerateIdsReqVO struct {
	BatchNum int `json:"batchNum"`
}

type GenerateIdsRespVO struct {
	common.BaseResp
	Ids []int64 `json:"ids"`
}
