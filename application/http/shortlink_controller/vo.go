package shortlink_controller

import "github.com/LeeZXin/z-base/common"

type CreateShortLinkReqVO struct {
	LongLink string `json:"longLink"`
}

type CreateShortLinkRespVO struct {
	common.BaseResp
	ShortLink string `json:"shortLink"`
}

type GetLongLinkReqVO struct {
	ShortLink string `json:"shortLink"`
}

type GetLongLinkRespVO struct {
	common.BaseResp
	LongLink string `json:"longLink"`
}
