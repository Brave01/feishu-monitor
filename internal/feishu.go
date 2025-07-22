package internal

import (
	"feishu-monitor/internal/service/fs"
	"feishu-monitor/internal/service/prometheus"
	"feishu-monitor/model/response"
	"feishu-monitor/util"
)

func AddFsOneRecord() {}

func AddFsMultiRecords() error {
	// 获取prometheus的数据，步长为1个小时
	allPrometheus, err := prometheus.GetAllPrometheus()
	if err != nil {
		return err
	}
	// 收集到的数据按照 ip 整合为一个map
	toFs := prometheus.PrometheusToFs(allPrometheus)
	// TODO 这里后面再优化
	// 将map转换为飞书消息格式的结构体
	rfs := util.MapToStruct(toFs)
	// 飞书发送消息的格式为map类型，这里再转换为map
	fsMaps := make([]map[string]interface{}, 0)
	for _, rf := range rfs {
		toMap := util.StructToMap(rf)
		fsMaps = append(fsMaps, toMap)
	}
	err = fs.AddTableMultiMsg(fsMaps)
	if err != nil {
		return err
	}
	return nil
}

func DelFsAllRecord() error {
	// 初始化
	mms := response.TableMsFs{}
	mms.Data.Total = 1
	recordId := make([]string, 0)
	for mms.Data.Total != 0 {
		// 获取多维表格的信息，获取RecordId
		ms, err := fs.GetTableMsg()
		if err != nil {
			return err
		}
		mms.Data.Total = ms.Data.Total
		// 获取所有recordId
		for _, item := range ms.Data.Items {
			recordId = append(recordId, item.RecordId)
		}
		// 删除
		err = fs.DelTableMsg(recordId)
		if err != nil {
			return err
		}
	}

	return nil
}
