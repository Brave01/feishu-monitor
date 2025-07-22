package response

type TableMsFs struct {
	Code int       `json:"code"`
	Msg  string    `json:"msg"`
	Data TableData `json:"data"`
}

type TableData struct {
	Items []DataItem `json:"items"`
	Total int        `json:"total"`
}

type DataItem struct {
	Fields   map[string]interface{} `json:"fields"`
	RecordId string                 `json:"record_id"`
}
