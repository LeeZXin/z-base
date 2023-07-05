package shortlink_service

import (
	"context"
	"github.com/LeeZXin/z-base/common"
	"github.com/LeeZXin/z-base/repo/store/shortlink_repo"
	"github.com/LeeZXin/z-base/repo/util"
	"github.com/LeeZXin/zsf/logger"
)

type LinkServiceImpl struct {
}

func NewLinkService() LinkService {
	return &LinkServiceImpl{}
}

func (s *LinkServiceImpl) CreateShortLink(ctx context.Context, reqDTO CreateShortLinkReqDTO) (resp CreateShortLinkRespDTO) {
	if reqDTO.LongLink == "" || len(reqDTO.LongLink) > 2000 {
		resp = CreateShortLinkRespDTO{
			BaseResp: common.NewBaseResp(common.InvalidParamsCode, "long link err"),
		}
		return
	}
	str := s.hash(reqDTO.LongLink)
	for i := 0; i < 10; i++ {
		linkReqDTO := shortlink_repo.InsertLinkReqDTO{
			ShortLink: str,
			LongLink:  reqDTO.LongLink,
		}
		linkRespDTO := shortlink_repo.StoreImpl.InsertLink(ctx, linkReqDTO)
		logger.Logger.WithContext(ctx).Infof("CreateShortLink %s %s", reqDTO.LongLink, str)
		if linkRespDTO.Success {
			resp = CreateShortLinkRespDTO{
				BaseResp:  common.DefaultSuccessResp,
				ShortLink: str,
			}
			return
		}
		// 失败默认是重复 然后拼接随机字符串
		str = str + util.RandomStr(2)
	}
	resp = CreateShortLinkRespDTO{
		BaseResp: common.NewBaseResp(common.ExecuteFailCode, "execute failed"),
	}
	return
}

func (*LinkServiceImpl) hash(shortLink string) string {
	hashRet := util.Murmur3([]byte(shortLink))
	return util.To62Str(hashRet)
}

func (*LinkServiceImpl) GetLongLink(ctx context.Context, reqDTO GetLongLinkReqDTO) (resp GetLongLinkRespDTO) {
	if reqDTO.ShortLink == "" {
		resp = GetLongLinkRespDTO{
			BaseResp: common.NewBaseResp(common.InvalidParamsCode, "empty link"),
		}
		return
	}
	linkReqDTO := shortlink_repo.GetLongLinkByShortLinkReqDTO{
		ShortLink: reqDTO.ShortLink,
	}
	ret := shortlink_repo.StoreImpl.GetLongLinkByShortLink(ctx, linkReqDTO)
	if ret.Exists {
		resp = GetLongLinkRespDTO{
			BaseResp: common.DefaultSuccessResp,
			LongLink: ret.LongLink,
		}
		return
	}
	resp = GetLongLinkRespDTO{
		BaseResp: common.NewBaseResp(common.DataNotExistsCode, "not exists"),
	}
	return
}
