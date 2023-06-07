package sid

import (
	"github.com/LeeZXin/z-base/common"
)

type GenerateIdsReqDTO struct {
	BatchNum int `json:"batchNum"`
}

type GenerateIdsRespDTO struct {
	common.BaseResp
	Result []int64
}
