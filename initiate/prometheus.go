package initiate

import (
	"feishu-monitor/global"
	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

func InitPrometheusClient() (v1.API, error) {
	client, err := api.NewClient(api.Config{Address: global.PROME_URL})
	newAPI := v1.NewAPI(client)
	if err != nil {
		return nil, err
	}
	return newAPI, nil
}
