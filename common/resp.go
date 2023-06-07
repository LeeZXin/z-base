package common

const (
	SuccessCode = 0
	// InvalidParamsCode 参数错误
	InvalidParamsCode = 100001
	// DataNotExistsCode 数据不存在
	DataNotExistsCode = 100002
	// ExecuteFailCode 执行错误
	ExecuteFailCode = 100003
	// SystemErrCode 系统错误
	SystemErrCode = 999999
)

var (
	DefaultSuccessResp     = NewBaseResp(SuccessCode, "success")
	DefaultFailBindArgResp = NewBaseResp(InvalidParamsCode, "args error")
)

// BaseResp 通用基础resp
type BaseResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (b *BaseResp) IsSuccess() bool {
	return b.Code == SuccessCode
}

func NewBaseResp(code int, message string) BaseResp {
	return BaseResp{
		Code:    code,
		Message: message,
	}
}
