package service

import (
	"context"
	"errors"
	"feishu-monitor/global"
	"feishu-monitor/util"
	"fmt"
	lark "github.com/larksuite/oapi-sdk-go/v3"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	larkbitable "github.com/larksuite/oapi-sdk-go/v3/service/bitable/v1"
)

func AddTableMsg() error {
	// 生成 Token
	token, err2 := util.GetFsToken()
	if err2 != nil {
		return err2
	}
	// 测试
	//var f request.Fields
	//f.Date = 1752646469000
	//f.CPUUsage = 0.1
	//f.DiskUsage = 0.2
	//f.MemoryUsage = 0.3
	//f.ManDays = 3
	//f.Priority = "不重要不紧急"
	//f.ServerName = "wwwww"
	//fsMap := util.StructToMap(f)
	//fmt.Println(fsMap)
	// 创建 Client
	client := lark.NewClient(global.APPID, global.APP_SECRET)
	// 创建请求对象
	req := larkbitable.NewCreateAppTableRecordReqBuilder().
		AppToken(global.APP_TOKEN).
		TableId(global.TABLE_ID).
		UserIdType(`open_id`).
		AppTableRecord(larkbitable.NewAppTableRecordBuilder().
			Fields(map[string]interface{}{}).
			Build()).
		Build()

	// 发起请求
	resp, err := client.Bitable.V1.AppTableRecord.Create(context.Background(), req, larkcore.WithTenantAccessToken(token.Tenant_Access_Token))

	// 处理错误
	if err != nil {
		return err
	}

	// 服务端错误处理
	if !resp.Success() {
		return errors.New(fmt.Sprintf("logId: %s, error response: \n%s", resp.RequestId(), larkcore.Prettify(resp.CodeError)))
	}

	// 业务处理
	fmt.Println(larkcore.Prettify(resp))
	return nil
}
