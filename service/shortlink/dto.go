package shortlink

import (
	"github.com/LeeZXin/z-base/common"
)

type CreateShortLinkReqDTO struct {
	LongLink string `json:"longLink"`
}

type CreateShortLinkRespDTO struct {
	common.BaseResp
	ShortLink string `json:"shortLink"`
}

type GetLongLinkReqDTO struct {
	ShortLink string `json:"shortLink"`
}

type GetLongLinkRespDTO struct {
	common.BaseResp
	LongLink string `json:"longLink"`
}
