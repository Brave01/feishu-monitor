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

func UpdateTableMsg(recordID string) error {
	// 生成 Token
	token, err2 := util.GetFsToken()
	if err2 != nil {
		return err2
	}
	// 创建 Client
	client := lark.NewClient(global.APPID, global.APP_SECRET)
	// 创建请求对象
	req := larkbitable.NewUpdateAppTableRecordReqBuilder().
		AppToken(global.APP_TOKEN).
		TableId(global.TABLE_ID).
		RecordId(recordID).
		AppTableRecord(larkbitable.NewAppTableRecordBuilder().
			Fields(map[string]interface{}{`日期`: 1752646469000}).
			Build()).
		Build()

	// 发起请求
	resp, err := client.Bitable.V1.AppTableRecord.Update(context.Background(), req, larkcore.WithTenantAccessToken(token.Tenant_Access_Token))

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
