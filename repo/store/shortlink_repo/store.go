package shortlink_repo

import (
	"context"
	"github.com/LeeZXin/zsf/property"
)

const (
	mysqlStoreType = "mysql"
	memStoreType   = "mem"
)

var (
	StoreImpl   LinkStore
	storeTypeFn = map[string]func() LinkStore{
		mysqlStoreType: func() LinkStore {
			return &LinkStoreImpl{}
		},
		memStoreType: func() LinkStore {
			return &MemLinkStoreImpl{}
		},
	}
)

func init() {
	StoreImpl = NewLinkStore()
}

type LinkStore interface {
	// InsertLink 插入记录
	InsertLink(context.Context, InsertLinkReqDTO) InsertLinkRespDTO
	// GetLongLinkByShortLink 获取长链
	GetLongLinkByShortLink(context.Context, GetLongLinkByShortLinkReqDTO) GetLongLinkByShortLinkRespDTO
}

func NewLinkStore() LinkStore {
	storeType := property.GetString("shortLink.store")
	fn, ok := storeTypeFn[storeType]
	if ok {
		return fn()
	}
	return storeTypeFn[memStoreType]()
}
