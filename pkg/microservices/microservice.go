package microservices

type MicroService struct {
	Timestamp         string  `json:"timestamp"`
	MsName            string  `json:"msname"`
	MsInstanceId      string  `json:"msinstanceid"`
	NodeId            string  `json:"nodeid"`
	CPUUtilization    float64 `json:"cpuutilization"`
	MemoryUtilization float64 `json:"memoryutilization"`
}

// 可能在实际的应用场景中，trace不需要下探到实例级别的深度
// 这里的调用图是服务级别的
type CallGraph struct {
	Timestamp string `json:"timestamp"`
	TraceId   string `json:"traceid"`
	RpcId     string `json:"rpcid"`
	RpcType   string `json:"rpctype"`
	Interface string `json:"interface"`
	// Upstream microservices name
	UM string `json:"um"`
	// Downstream microservices name
	DM string `json:"dm"`
	RT string `json:"rt"`
}

type CallInfo struct {
	Timestamp    string  `json:"timestamp"`
	MsName       string  `json:"msname"`
	MsInstanceId string  `json:"msinstanceid"`
	ResponseTime float64 `json:"reponsetime,omitempty"`
	MSR          float64 `json:"msr,omitempty"`
	Role         string  `json:"role"`
	Type         string  `json:"type"`
}
