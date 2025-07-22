package fs

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

func AddTableOneMsg(rf map[string]interface{}) error {
	// 生成 Token
	token, err2 := util.GetFsToken()
	if err2 != nil {
		return err2
	}
	// 创建 Client
	client := lark.NewClient(global.APPID, global.APP_SECRET)
	// 创建请求对象
	req := larkbitable.NewCreateAppTableRecordReqBuilder().
		AppToken(global.APP_TOKEN).
		TableId(global.TABLE_ID).
		UserIdType(`open_id`).
		AppTableRecord(larkbitable.NewAppTableRecordBuilder().
			Fields(rf).
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

func AddTableMultiMsg(rfs []map[string]interface{}) error {
	// 创建 Client
	client := lark.NewClient(global.APPID, global.APP_SECRET)
	// 创建请求对象
	records := createAppTableRecord(rfs)
	req := larkbitable.NewBatchCreateAppTableRecordReqBuilder().
		AppToken(global.APP_TOKEN).
		TableId(global.TABLE_ID).
		UserIdType(`open_id`).
		Body(larkbitable.NewBatchCreateAppTableRecordReqBodyBuilder().
			Records(records).Build()).
		Build()

	// 发起请求
	resp, err := client.Bitable.V1.AppTableRecord.BatchCreate(context.Background(), req)
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

func createAppTableRecord(rfs []map[string]interface{}) []*larkbitable.AppTableRecord {
	appTableRecords := make([]*larkbitable.AppTableRecord, 0)
	var appTableRecord *larkbitable.AppTableRecord
	for _, rf := range rfs {
		appTableRecord = larkbitable.NewAppTableRecordBuilder().
			Fields(rf).
			Build()
		appTableRecords = append(appTableRecords, appTableRecord)
	}
	return appTableRecords
}
