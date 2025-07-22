package service

import (
	"context"
	"errors"
	"feishu-monitor/global"
	"feishu-monitor/model/response"
	"feishu-monitor/util"
	"fmt"
	lark "github.com/larksuite/oapi-sdk-go/v3"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	larkbitable "github.com/larksuite/oapi-sdk-go/v3/service/bitable/v1"
)

func GetTableMsg() (tableMs *response.TableMsFs, err error) {
	// 生成 Token
	token, err2 := util.GetFsToken()
	if err2 != nil {
		return nil, err2
	}
	// 创建 Client
	client := lark.NewClient(global.APPID, global.APP_SECRET)
	// 创建请求对象
	req := larkbitable.NewSearchAppTableRecordReqBuilder().
		AppToken(global.APP_TOKEN).
		TableId(global.TABLE_ID).
		UserIdType(`open_id`).
		PageSize(20).
		Body(larkbitable.NewSearchAppTableRecordReqBodyBuilder().
			Build()).
		Build()

	// 发起请求
	resp, err := client.Bitable.V1.AppTableRecord.Search(context.Background(),
		req,
		larkcore.WithUserAccessToken(token.Tenant_Access_Token))

	// 处理错误
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// 服务端错误处理
	if !resp.Success() {
		return nil, errors.New(fmt.Sprintf("logId: %s, error response: \n%s", resp.RequestId(), larkcore.Prettify(resp.CodeError)))
	}
	// 结构体转换
	toStruct, err := util.FsToStruct(resp)
	if err != nil {
		return nil, err
	} else {
		// 业务处理
		return &toStruct, nil
	}
}
