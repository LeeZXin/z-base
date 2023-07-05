package sid_service

import (
	"context"
	"github.com/LeeZXin/z-base/repo/redis"
	"github.com/LeeZXin/zsf/logger"
	"github.com/LeeZXin/zsf/property"
)

// 雪花id service

var (
	ServiceImpl IdGenerateService

	NodeIdScript = `
		local i = tonumber(redis.call('INCR', KEYS[1]))
		if  i > 1024 then
			redis.call('SET', KEYS[1], 0)
			return 0
		else
			return i
		end
	`
)

func init() {
	nodeId := getNodeId()
	ServiceImpl = NewIdGenerateService(nodeId)
}

func getNodeId() int64 {
	if property.GetBool("snowflake.nodeIdFromRedis") {
		result, err := redis.Client.Eval(context.Background(), NodeIdScript, []string{"snowflake:nodeId"}).Result()
		if err != nil {
			logger.Logger.Panic(err.Error())
		}
		id, ok := result.(int64)
		if ok {
			return id
		}
		logger.Logger.Panic("snowflake.nodeIdFromRedis result is not int64")
	}
	return property.GetInt64("snowflake.nodeId")
}

type IdGenerateService interface {
	// GenerateIds 生成id
	GenerateIds(GenerateIdsReqDTO) GenerateIdsRespDTO
}
