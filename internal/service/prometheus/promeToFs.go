package prometheus

import (
	"feishu-monitor/global"
	"feishu-monitor/model/prometheus"
	"fmt"
	"github.com/prometheus/common/model"
	"slices"
	"strings"
)

func PrometheusToFs(values map[string]model.Value) map[string]*prometheus.PrometheusMsg {
	pms := make(map[string]*prometheus.PrometheusMsg, 100)

	for name, value := range values {
		switch value.Type() {
		case model.ValMatrix:
			metric := value.(model.Matrix)
			for _, stream := range metric {
				// 判断是否存在 IP 地址字段, map的key必须为ip地址，没有就过滤掉
				if ip, ok := stream.Metric["instance"]; ok {
					// 根据instance 判断key值是否存在，不存在则申请一个
					if ps, ok := pms[string(ip)]; !ok {
						ps = &prometheus.PrometheusMsg{}
						pms[string(ip)] = ps

						// 去掉端口号 :9100
						ps.Instance = strings.Split(string(ip), ":")[0]
						if nodeType, ok := stream.Metric["nodetype"]; ok {
							ps.NodeType = string(nodeType)
						}
						// 根据name，判断查询语句为 CPU，内存还是磁盘
						switch name {
						case global.MEM_USAGE:
							m, a := getMaxAvgNum(stream.Values)
							ps.MemoryAvgUsage = a
							ps.MemoryMaxUsage = m
							if a < 0.4 && m < 0.4 {
								ps.Priority = append(ps.Priority, global.UsageLowMem)
							}
						case global.CPU_USAGE:
							m, a := getMaxAvgNum(stream.Values)
							ps.CpuAvgUsage = a
							ps.CpuMaxUsage = m
							if a < 0.4 && m < 0.4 {
								ps.Priority = append(ps.Priority, global.UsageLowCpu)
							}
						case global.DISKUSAGE:
							// TODO 磁盘可能会有多个挂载点，这里只判断了两种
							if string(stream.Metric["mountpoint"]) == "/" {
								ps.MountRoot = string(stream.Metric["mountpoint"])
								m, a := getMaxAvgNum(stream.Values)
								ps.DiskAvgUsageRoot = a
								ps.DiskMaxUsageRoot = m
								if a < 0.4 && m < 0.4 {
									ps.Priority = append(ps.Priority, global.UsageLowRootDisk)
								}
							}
							if string(stream.Metric["mountpoint"]) == "/data" {
								ps.MountOther = string(stream.Metric["mountpoint"])
								m, a := getMaxAvgNum(stream.Values)
								ps.DiskAvgUsageOther = a
								ps.DiskMaxUsageOther = m
								if a < 0.4 && m < 0.4 {
									ps.Priority = append(ps.Priority, global.UsageLowOtherDisk)
								}
							}
						default:
							_ = fmt.Errorf("暂时没有此种类型")
						}
					} else {
						if nodeType, ok := stream.Metric["nodetype"]; ok {
							ps.NodeType = string(nodeType)
						}
						switch name {
						case global.MEM_USAGE:
							m, a := getMaxAvgNum(stream.Values)
							ps.MemoryAvgUsage = a
							ps.MemoryMaxUsage = m
							if a < 0.4 && m < 0.4 {
								ps.Priority = append(ps.Priority, global.UsageLowMem)
							}
						case global.CPU_USAGE:
							m, a := getMaxAvgNum(stream.Values)
							ps.CpuAvgUsage = a
							ps.CpuMaxUsage = m
							if a < 0.4 && m < 0.4 {
								ps.Priority = append(ps.Priority, global.UsageLowCpu)
							}
						case global.DISKUSAGE:
							// TODO 磁盘可能会有多个挂载点，这里只判断了两种
							if string(stream.Metric["mountpoint"]) == "/" {
								ps.MountRoot = string(stream.Metric["mountpoint"])
								m, a := getMaxAvgNum(stream.Values)
								ps.DiskAvgUsageRoot = a
								ps.DiskMaxUsageRoot = m
								if a < 0.4 && m < 0.4 {
									ps.Priority = append(ps.Priority, global.UsageLowRootDisk)
								}
							}
							if string(stream.Metric["mountpoint"]) == "/data" {
								ps.MountOther = string(stream.Metric["mountpoint"])
								m, a := getMaxAvgNum(stream.Values)
								ps.DiskAvgUsageOther = a
								ps.DiskMaxUsageOther = m
								if a < 0.4 && m < 0.4 {
									ps.Priority = append(ps.Priority, global.UsageLowOtherDisk)
								}
							}
						default:
							_ = fmt.Errorf("暂时没有此种类型")
						}
					}

				}
			}

		default:
			_ = fmt.Errorf("unknown type: %T", values)
		}
	}
	return pms
}

// 返回切片的最大值和平均值
func getMaxAvgNum(msm []model.SamplePair) (m, a float64) {
	var sum float64
	dl := make([]float64, 0)
	l := len(msm)
	for _, v := range msm {
		dl = append(dl, float64(v.Value))
		sum += float64(v.Value)
	}
	// 转换为百分位
	return slices.Max(dl) / 100, (sum / float64(l)) / 100
}
