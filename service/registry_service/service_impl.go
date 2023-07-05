package registry_service

import (
	"context"
	"errors"
	"github.com/LeeZXin/zsf/logger"
	"hash/crc32"
	"sort"
	"strings"
	"sync"
	"time"
)

// 默认分64segment
const (
	segmentSize = 64
)

var (
	registryObj = newRegistry()
)

func init() {
	// 每十秒钟清理过期数据
	go func() {
		for {
			registryObj.ClearExpire()
			time.Sleep(10 * time.Second)
		}
	}()
}

type RegistryServiceImpl struct {
}

func NewRegistryService() RegistryService {
	return &RegistryServiceImpl{}
}

// RegisterService 注册service
func (*RegistryServiceImpl) RegisterService(ctx context.Context, reqDTO RegisterServiceReqDTO) {
	logger.Logger.WithContext(ctx).Info("RegisterService: ", reqDTO)
	registryObj.Register(reqDTO)
}

// PassTTL 续期
func (*RegistryServiceImpl) PassTTL(ctx context.Context, reqDTO PassTtlReqDTO) error {
	logger.Logger.WithContext(ctx).Info("PassTTL: ", reqDTO)
	return registryObj.PassTTL(reqDTO)
}

// DeregisterService 注销
func (*RegistryServiceImpl) DeregisterService(ctx context.Context, reqDTO DeregisterServiceReqDTO) error {
	logger.Logger.WithContext(ctx).Info("DeregisterService: ", reqDTO)
	return registryObj.Deregister(reqDTO)
}

// GetServiceInfoList 获取服务列表
func (*RegistryServiceImpl) GetServiceInfoList(ctx context.Context, serviceName string) []ServiceInfo {
	return registryObj.GetServiceInfoList(serviceName)
}

// ServiceInfo 注册服务信息
type ServiceInfo struct {
	ServiceName   string        `json:"serviceName,omitempty"`
	Ip            string        `json:"ip,omitempty"`
	Port          int           `json:"port,omitempty"`
	InstanceId    string        `json:"instanceId,omitempty"`
	Weight        int           `json:"weight,omitempty"`
	Version       string        `json:"version,omitempty"`
	LeaseDuration time.Duration `json:"leaseDuration,omitempty"`
	ExpireTime    time.Time     `json:"expireTime"`
}

type registry struct {
	segments []*segment
}

func newRegistry() *registry {
	segments := make([]*segment, 0, segmentSize)
	for i := 0; i < segmentSize; i++ {
		segments = append(segments, &segment{
			serviceCache: make(map[string]map[string]*ServiceInfo, 8),
			cacheMu:      sync.Mutex{},
		})
	}
	return &registry{
		segments: segments,
	}
}

// Register 注册service
func (r *registry) Register(reqDTO RegisterServiceReqDTO) {
	r.getSegment(reqDTO.ServiceName).Register(reqDTO)
}

// PassTTL 续期
func (r *registry) PassTTL(reqDTO PassTtlReqDTO) error {
	return r.getSegment(reqDTO.ServiceName).PassTTL(reqDTO)
}

// Deregister 注销
func (r *registry) Deregister(reqDTO DeregisterServiceReqDTO) error {
	return r.getSegment(reqDTO.ServiceName).Deregister(reqDTO)
}

// GetServiceInfoList 获取服务列表
func (r *registry) GetServiceInfoList(serviceName string) []ServiceInfo {
	return r.getSegment(serviceName).GetServiceInfoList(serviceName)
}

// ClearExpire 清理过期数据
func (r *registry) ClearExpire() {
	for _, seg := range r.segments {
		seg.ClearExpire()
	}
}

func (r *registry) getSegment(serviceName string) *segment {
	hashRet := crc32.ChecksumIEEE([]byte(serviceName))
	return r.segments[int(hashRet)&0x3f]
}

type segment struct {
	serviceCache map[string]map[string]*ServiceInfo
	cacheMu      sync.Mutex
}

func (r *segment) Register(reqDTO RegisterServiceReqDTO) {
	r.cacheMu.Lock()
	defer r.cacheMu.Unlock()
	duration := time.Duration(reqDTO.LeaseDuration) * time.Second
	info := &ServiceInfo{
		ServiceName:   reqDTO.ServiceName,
		Ip:            reqDTO.Ip,
		Port:          reqDTO.Port,
		InstanceId:    reqDTO.InstanceId,
		Weight:        reqDTO.Weight,
		Version:       reqDTO.Version,
		LeaseDuration: duration,
		ExpireTime:    time.Now().Add(duration),
	}
	infoMap, b := r.serviceCache[info.ServiceName]
	if b {
		infoMap[info.InstanceId] = info
	} else {
		infoMap = make(map[string]*ServiceInfo, 8)
		infoMap[info.InstanceId] = info
		r.serviceCache[info.ServiceName] = infoMap
	}
}

func (r *segment) PassTTL(reqDTO PassTtlReqDTO) error {
	r.cacheMu.Lock()
	defer r.cacheMu.Unlock()
	infoMap, b := r.serviceCache[reqDTO.ServiceName]
	if !b {
		return errors.New("service name not found")
	}
	info, b := infoMap[reqDTO.InstanceId]
	if !b || info.ExpireTime.Before(time.Now()) {
		return errors.New("instanceId not found")
	}
	info.ExpireTime = time.Now().Add(info.LeaseDuration)
	return nil
}

func (r *segment) Deregister(reqDTO DeregisterServiceReqDTO) error {
	r.cacheMu.Lock()
	defer r.cacheMu.Unlock()
	infoMap, b := r.serviceCache[reqDTO.ServiceName]
	if !b {
		return errors.New("service name not found")
	}
	_, b = infoMap[reqDTO.InstanceId]
	if !b {
		return errors.New("instanceId not found")
	}
	delete(infoMap, reqDTO.InstanceId)
	return nil
}

func (r *segment) GetServiceInfoList(serviceName string) []ServiceInfo {
	r.cacheMu.Lock()
	defer r.cacheMu.Unlock()
	infoMap, b := r.serviceCache[serviceName]
	if !b {
		return []ServiceInfo{}
	}
	ret := make([]ServiceInfo, 0, len(infoMap))
	now := time.Now()
	for _, info := range infoMap {
		i := info
		if i.ExpireTime.Before(now) {
			continue
		}
		ret = append(ret, ServiceInfo{
			ServiceName:   i.ServiceName,
			Ip:            i.Ip,
			Port:          i.Port,
			InstanceId:    i.InstanceId,
			Weight:        i.Weight,
			Version:       i.Version,
			LeaseDuration: i.LeaseDuration,
			ExpireTime:    i.ExpireTime,
		})
	}
	sort.Slice(ret, func(i, j int) bool {
		return strings.Compare(ret[i].InstanceId, ret[j].InstanceId) < 0
	})
	return ret
}

func (r *segment) ClearExpire() {
	r.cacheMu.Lock()
	defer r.cacheMu.Unlock()
	now := time.Now()
	for serviceName, infoMap := range r.serviceCache {
		for instanceId, info := range infoMap {
			if info.ExpireTime.Before(now) {
				delete(infoMap, instanceId)
			}
		}
		if len(infoMap) == 0 {
			delete(r.serviceCache, serviceName)
		}
	}
}
