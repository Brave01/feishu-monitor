package util

import (
	"errors"
	"feishu-monitor/model/response"
	"github.com/jinzhu/copier"
	larkbitable "github.com/larksuite/oapi-sdk-go/v3/service/bitable/v1"
)

func FsToStruct(resp *larkbitable.SearchAppTableRecordResp) (tableMs *response.TableMsFs, err error) {
	tableMs = &response.TableMsFs{}
	err = copier.Copy(tableMs, resp)
	if err != nil {
		return tableMs, err
	}
	return tableMs, nil
}

func FsMetaToStruct(resp *larkbitable.DisplayApp) (metaMs *response.FeiShuApp, err error) {
	if resp == nil {
		return nil, errors.New("resp is nil")
	}
	metaMs = &response.FeiShuApp{}
	if resp.AppToken != nil {
		metaMs.AppToken = *resp.AppToken
	}
	if resp.Name != nil {
		metaMs.Name = *resp.Name
	}
	if resp.Revision != nil {
		metaMs.Revision = *resp.Revision
	}
	if resp.IsAdvanced != nil {
		metaMs.IsAdvanced = *resp.IsAdvanced
	}
	if resp.TimeZone != nil {
		metaMs.TimeZone = *resp.TimeZone
	}
	if resp.FormulaType != nil {
		metaMs.FormulaType = *resp.FormulaType
	}
	if resp.AdvanceVersion != nil {
		metaMs.AdvanceVersion = *resp.AdvanceVersion
	}
	return metaMs, nil
}
