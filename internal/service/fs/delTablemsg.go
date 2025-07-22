package fs

import (
	"context"
	"errors"
	"feishu-monitor/global"
	"fmt"
	lark "github.com/larksuite/oapi-sdk-go/v3"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	larkbitable "github.com/larksuite/oapi-sdk-go/v3/service/bitable/v1"
)

func DelTableMsg(recordID []string) error {
	// 创建 Client
	client := lark.NewClient(global.APPID, global.APP_SECRET)
	// 创建请求对象
	req := larkbitable.NewBatchDeleteAppTableRecordReqBuilder().
		AppToken(global.APP_TOKEN).
		TableId(global.TABLE_ID).
		Body(larkbitable.NewBatchDeleteAppTableRecordReqBodyBuilder().
			Records(recordID).
			Build()).
		Build()

	// 发起请求
	resp, err := client.Bitable.V1.AppTableRecord.BatchDelete(context.Background(), req)
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
