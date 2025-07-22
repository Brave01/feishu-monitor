package prometheus

import (
	"context"
	"feishu-monitor/global"
	"feishu-monitor/initiate"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
	"log"
	"sync"
	"time"
)

func getPrometheus(quires map[string]string, start, end time.Time) (map[string]model.Value, error) {
	// 步长
	step := time.Hour
	// 初始化 prometheus 查询接口
	v1Api, err := initiate.InitPrometheusClient()
	if err != nil {
		return nil, err
	}
	// 调用查询
	// 启用协程池
	var wg sync.WaitGroup
	rs := make(map[string]model.Value, 10)
	// 读写锁
	var mu sync.Mutex
	for name, query := range quires {
		wg.Add(1)
		go func(name, query string) {
			defer wg.Done()
			queryResult, _, err := v1Api.QueryRange(context.Background(), query, v1.Range{
				Start: start,
				End:   end,
				Step:  step,
			})
			if err != nil {
				log.Fatalf("query %s failed: %v", name, err)
			}
			mu.Lock()
			rs[name] = queryResult
			mu.Unlock()
		}(name, query)
	}
	wg.Wait()

	return rs, nil
}

func GetAllPrometheus() (map[string]model.Value, error) {
	// 上月的第一天到最后一天
	firstOfLastMouth := time.Date(time.Now().Year(), time.Now().Month()-1, 1, 0, 0, 0, 0, time.Local)
	lastOfLastMouth := firstOfLastMouth.AddDate(0, 1, -1)

	queries := map[string]string{
		global.MEM_USAGE: `100 * (node_memory_MemTotal_bytes - node_memory_MemAvailable_bytes) / node_memory_MemTotal_bytes`,
		global.DISKUSAGE: `100 - ((node_filesystem_avail_bytes{job="node",mountpoint=~".*",fstype=~"ext4|xfs"} * 100) / node_filesystem_size_bytes{job="node",mountpoint=~".*",fstype=~"ext4|xfs"})`,
		global.CPU_USAGE: `100 - ((avg by (instance,job,env)(irate(node_cpu_seconds_total{mode="idle"}[30s]))) * 100)`,
	}
	rs, err := getPrometheus(queries, firstOfLastMouth, lastOfLastMouth)
	if err != nil {
		return nil, err
	}
	return rs, nil
}
