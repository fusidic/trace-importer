package nodes

// Node means machine in alibaba trace instead of node in neo4j structure.
type Node struct {
	Timestamp         string  `json:"timestamp"`
	NodeId            string  `json:"nodeid"`
	CPUUtilization    float64 `json:"cpuutilization"`
	MemoryUtilization float64 `json:"memoryutilization"`
}

