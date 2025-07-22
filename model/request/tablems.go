package request

type Fields struct {
	Date              int64    `json:"日期"`
	DiskUsageRoot     float64  `json:"磁盘使用率（根分区）"`
	DiskMaxUsageRoot  float64  `json:"磁盘使用率峰值（根分区）"`
	DiskUsageOther    float64  `json:"磁盘使用率（data分区）"`
	DiskMaxUsageOther float64  `json:"磁盘使用率峰值（data分区）"`
	CpuUsage          float64  `json:"CPU使用率"`
	CpuMaxUsage       float64  `json:"CPU使用率峰值"`
	MemoryUsage       float64  `json:"内存使用率"`
	MemMaxUsage       float64  `json:"内存峰值"`
	Priority          []string `json:"低资源利用率"`
	ServerName        string   `json:"服务器名称"`
	ServerIP          string   `json:"服务器IP地址"`
}
