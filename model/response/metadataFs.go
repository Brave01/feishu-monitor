package response

type ResFeiShuResponse struct {
	Code int        `json:"code"`
	Msg  string     `json:"msg"`
	Data FeiShuData `json:"data"`
}

type FeiShuData struct {
	App FeiShuApp `json:"app"`
}

type FeiShuApp struct {
	AppToken       string `json:"app_token"`
	Name           string `json:"name"`
	Revision       int    `json:"revision"`
	IsAdvanced     bool   `json:"is_advanced"`
	TimeZone       string `json:"time_zone"`
	FormulaType    int    `json:"formula_type"`
	AdvanceVersion string `json:"advance_version"`
}
