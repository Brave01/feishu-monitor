package response

type FeiShuApp struct {
	AppToken       string `json:"app_token,omitempty"`       // 多维表格的 app_token;[app_token 参数说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/bitable/notification#8121eebe)
	Name           string `json:"name,omitempty"`            // 多维表格的名字
	Revision       int    `json:"revision,omitempty"`        // 多维表格的版本号（对多维表格进行修改时更新，如新增、删除数据表，修改数据表名等，初始为1，每次更新+1）
	IsAdvanced     bool   `json:"is_advanced,omitempty"`     // 多维表格是否开启了高级权限。取值包括：;- true：表示开启了高级权限;- false：表示关闭了高级权限;;[了解更多：使用多维表格高级权限](https://www.feishu.cn/hc/zh-CN/articles/588604550568)
	TimeZone       string `json:"time_zone,omitempty"`       // 文档时区
	FormulaType    int    `json:"formula_type,omitempty"`    // 文档公式字段类型
	AdvanceVersion string `json:"advance_version,omitempty"` // 文档高级权限版本
}
