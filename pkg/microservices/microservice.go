package microservices

type MicroService struct {
	Timestamp         string  `json:"timestamp"`
	MsName            string  `json:"msname"`
	MsInstanceId      string  `json:"msinstanceid"`
	NodeId            string  `json:"nodeid"`
	CPUUtilization    float64 `json:"cpuutilization"`
	MemoryUtilization float64 `json:"memoryutilization"`
}
