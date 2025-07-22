package prometheus

type PrometheusMsg struct {
	// IP地址
	Instance string
	// 主机类型
	NodeType string
	// CPU 平均占用率
	CpuAvgUsage float64
	// 内存 平均占用率
	MemoryAvgUsage float64
	// 磁盘 平均占用率
	DiskAvgUsageRoot  float64
	DiskAvgUsageOther float64
	// CPU 峰值
	CpuMaxUsage float64
	// 内存 峰值
	MemoryMaxUsage float64
	// 磁盘 峰值
	DiskMaxUsageRoot  float64
	DiskMaxUsageOther float64
	// 挂载根分区
	MountRoot string
	// 挂载其他分区
	MountOther string
	// 低资源利用率
	Priority []string
}
