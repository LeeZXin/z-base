package registry_service

type DeregisterServiceReqDTO struct {
	ServiceName string `json:"serviceName"`
	InstanceId  string `json:"instanceId"`
}

type PassTtlReqDTO struct {
	ServiceName string `json:"serviceName"`
	InstanceId  string `json:"instanceId"`
}

type RegisterServiceReqDTO struct {
	ServiceName   string `json:"service_name"`
	Ip            string `json:"ip"`
	Port          int    `json:"port"`
	InstanceId    string `json:"instanceId"`
	Weight        int    `json:"weight"`
	Version       string `json:"version"`
	LeaseDuration int    `json:"leaseDuration"`
}
