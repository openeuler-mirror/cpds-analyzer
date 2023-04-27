package monitor

import (
	"cpds/cpds-analyzer/pkg/prometheus"
	"encoding/json"
	"errors"
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

func GetMonitorDataFromDetector(url string) ([]prometheus.Metric, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response map[string]interface{}
	if err := jsoniter.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}
	if response["status"] != float64(200) {
		return nil, errors.New("cannot get monitor data from detector")
	}

	var n []prometheus.Metric
	for _, data := range response["data"].([]interface{}) {
		var ns prometheus.Metric
		dataBytes, err := jsoniter.Marshal(data)
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal(dataBytes, &ns); err != nil {
			return nil, err
		}
		n = append(n, ns)
	}
	return n, nil
}
