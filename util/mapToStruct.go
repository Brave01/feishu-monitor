package util

import (
	"feishu-monitor/model/prometheus"
	"feishu-monitor/model/request"
	"time"
)

func MapToStruct(m map[string]*prometheus.PrometheusMsg) (rfs []request.Fields) {
	var fs request.Fields
	for _, v := range m {
		fs.ServerIP = v.Instance
		fs.ServerName = v.NodeType
		fs.DiskUsageOther = v.DiskAvgUsageOther
		fs.DiskMaxUsageOther = v.DiskMaxUsageOther
		fs.DiskUsageRoot = v.DiskAvgUsageRoot
		fs.DiskMaxUsageRoot = v.DiskMaxUsageRoot
		fs.MemMaxUsage = v.MemoryMaxUsage
		fs.MemoryUsage = v.MemoryAvgUsage
		fs.CpuUsage = v.CpuAvgUsage
		fs.CpuMaxUsage = v.CpuMaxUsage
		fs.Priority = v.Priority
		fs.Date = time.Now().UnixMilli()
		rfs = append(rfs, fs)
	}

	return rfs
}
