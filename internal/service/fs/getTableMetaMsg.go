package fs

import (
	"context"
	"feishu-monitor/global"
	"feishu-monitor/model/response"
	"feishu-monitor/util"
	"fmt"
	"github.com/larksuite/oapi-sdk-go/v3"
	"github.com/larksuite/oapi-sdk-go/v3/core"
	"github.com/larksuite/oapi-sdk-go/v3/service/bitable/v1"
)

// SDK 使用文档：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/server-side-sdk/golang-sdk-guide/preparations
// 复制该 Demo 后, 需要将 "YOUR_APP_ID", "YOUR_APP_SECRET" 替换为自己应用的 APP_ID, APP_SECRET.
// 以下示例代码默认根据文档示例值填充，如果存在代码问题，请在 API 调试台填上相关必要参数后再复制代码使用
func GetTableMetaMsg() (mms *response.FeiShuApp, err error) {
	// 创建 Client
	client := lark.NewClient(global.APPID, global.APP_SECRET)
	// 创建请求对象
	req := larkbitable.NewGetAppReqBuilder().
		AppToken(`KbmBbNojxaZexLsL539cQpJvnVh`).
		Build()

	// 发起请求
	resp, err := client.Bitable.V1.App.Get(context.Background(), req)

	// 处理错误
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// 服务端错误处理
	if !resp.Success() {
		fmt.Printf("logId: %s, error response: \n%s", resp.RequestId(), larkcore.Prettify(resp.CodeError))
		return nil, err
	}

	// 业务处理
	ms, err := util.FsMetaToStruct(resp.Data.App)
	if err != nil {
		return nil, err
	}

	return ms, nil
}
