package shortlink

import (
	"context"
	"github.com/LeeZXin/z-base/repo/store"
	"github.com/LeeZXin/zsf/logger"
	"sync"
)

type LinkStoreImpl struct {
}

// InsertLink 插入记录
func (*LinkStoreImpl) InsertLink(ctx context.Context, reqDTO InsertLinkReqDTO) (respDTO InsertLinkRespDTO) {
	model := Short4LongLinkModel{
		ShortLink: reqDTO.ShortLink,
		LongLink:  reqDTO.LongLink,
		IsDelete:  0,
	}
	insert, err := store.Orm.Insert(model)
	if err != nil {
		logger.Logger.WithContext(ctx).Error(err.Error())
	}
	respDTO = InsertLinkRespDTO{
		Success: insert > 0 && err == nil,
	}
	return
}

// GetLongLinkByShortLink 获取长链
func (*LinkStoreImpl) GetLongLinkByShortLink(ctx context.Context, reqDTO GetLongLinkByShortLinkReqDTO) (respDTO GetLongLinkByShortLinkRespDTO) {
	query := store.Orm.Where("short_link = ?", reqDTO.ShortLink)
	query.And("is_delete = 0")
	var model Short4LongLinkModel
	ok, err := query.Get(&model)
	if err != nil {
		logger.Logger.WithContext(ctx).Info(err.Error())
	}
	if ok {
		respDTO = GetLongLinkByShortLinkRespDTO{
			Exists:   true,
			LongLink: model.LongLink,
		}
	} else {
		respDTO = GetLongLinkByShortLinkRespDTO{
			Exists: false,
		}
	}
	return
}

// MemLinkStoreImpl 内存存储 测试用
type MemLinkStoreImpl struct {
	store sync.Map
}

// InsertLink 插入记录
func (s *MemLinkStoreImpl) InsertLink(ctx context.Context, reqDTO InsertLinkReqDTO) (respDTO InsertLinkRespDTO) {
	_, loaded := s.store.LoadOrStore(reqDTO.ShortLink, reqDTO.LongLink)
	if loaded {
		respDTO = InsertLinkRespDTO{Success: false}
		return
	}
	respDTO = InsertLinkRespDTO{Success: true}
	return
}

// GetLongLinkByShortLink 获取长链
func (s *MemLinkStoreImpl) GetLongLinkByShortLink(ctx context.Context, reqDTO GetLongLinkByShortLinkReqDTO) (respDTO GetLongLinkByShortLinkRespDTO) {
	value, ok := s.store.Load(reqDTO.ShortLink)
	if ok {
		respDTO = GetLongLinkByShortLinkRespDTO{
			Exists:   true,
			LongLink: value.(string),
		}
	} else {
		respDTO = GetLongLinkByShortLinkRespDTO{
			Exists: false,
		}
	}
	return
}
