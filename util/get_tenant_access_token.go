package util

import (
	"bytes"
	"encoding/json"
	"feishu-monitor/global"
	"io"
	"net/http"
)

const (
	url = "https://open.feishu.cn/open-apis/auth/v3/tenant_access_token/internal"
)

type GetTATokenRes struct {
	Code                int    `json:"code"`
	Msg                 string `json:"msg"`
	Tenant_Access_Token string `json:"tenant_access_token"`
	Expire              int64  `json:"expire"`
}

type TaTokenReq struct {
	AppId     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
}

func GetFsToken() (*GetTATokenRes, error) {
	var reqToken TaTokenReq
	reqToken.AppId = global.APPID
	reqToken.AppSecret = global.APP_SECRET
	jsonStr, _ := json.Marshal(reqToken)
	body := bytes.NewBuffer(jsonStr)
	resp, err := http.Post(url, global.CONTENT_TYPE, body)
	if err != nil {
		return nil, err
	}
	all, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var res GetTATokenRes
	err = json.Unmarshal(all, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
