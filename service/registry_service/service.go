package registry_service

import "context"

var (
	ServiceImpl = NewRegistryService()
)

type RegistryService interface {
	RegisterService(context.Context, RegisterServiceReqDTO)
	DeregisterService(context.Context, DeregisterServiceReqDTO) error
	PassTTL(context.Context, PassTtlReqDTO) error
	GetServiceInfoList(ctx context.Context, serviceName string) []ServiceInfo
}
