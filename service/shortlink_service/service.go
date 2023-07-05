package shortlink_service

import "context"

// 短链service

var (
	ServiceImpl LinkService
)

func init() {
	ServiceImpl = NewLinkService()
}

type LinkService interface {
	// CreateShortLink 生成短链
	CreateShortLink(context.Context, CreateShortLinkReqDTO) CreateShortLinkRespDTO
	// GetLongLink 获取长链
	GetLongLink(context.Context, GetLongLinkReqDTO) GetLongLinkRespDTO
}
