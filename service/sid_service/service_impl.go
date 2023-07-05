package sid_service

import (
	"github.com/LeeZXin/z-base/common"
	"github.com/bwmarrin/snowflake"
	"sync"
)

type IdGenerateServiceImpl struct {
	node *snowflake.Node
	mu   sync.Mutex
}

// NewIdGenerateService 初始化雪花id node
func NewIdGenerateService(nodeId int64) IdGenerateService {
	if nodeId < 0 {
		nodeId = 0
	}
	nodeId = nodeId % 1024
	n, _ := snowflake.NewNode(nodeId)
	return &IdGenerateServiceImpl{
		node: n,
		mu:   sync.Mutex{},
	}
}

// GenerateIds 生成雪花id
func (s *IdGenerateServiceImpl) GenerateIds(reqDTO GenerateIdsReqDTO) (resp GenerateIdsRespDTO) {
	if reqDTO.BatchNum < 0 {
		resp = GenerateIdsRespDTO{
			BaseResp: common.NewBaseResp(common.InvalidParamsCode, "batchNum should greater than 0"),
		}
		return
	}
	if reqDTO.BatchNum > 5000 {
		resp = GenerateIdsRespDTO{
			BaseResp: common.NewBaseResp(common.InvalidParamsCode, "batchNum should less than 5000"),
		}
		return
	}
	if s.node == nil {
		resp = GenerateIdsRespDTO{
			BaseResp: common.NewBaseResp(common.SystemErrCode, "nil node"),
		}
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	result := make([]int64, 0, reqDTO.BatchNum)
	for i := 0; i < reqDTO.BatchNum; i++ {
		result = append(result, s.node.Generate().Int64())
	}
	resp = GenerateIdsRespDTO{
		BaseResp: common.DefaultSuccessResp,
		Result:   result,
	}
	return
}
