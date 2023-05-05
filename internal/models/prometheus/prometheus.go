package prometheus

import (
	"cpds/cpds-analyzer/pkg/prometheus"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

type Operator interface {
	Query(expr string, timestamp int64) (*prometheus.MetricData, error)

	QueryRange(expr string, startTime, endTime int64, step int64) (*prometheus.MetricData, error)
}

type operator struct {
	detectorConfig *detectorConfig
}

type detectorConfig struct {
	host string
	port int
}

func NewOperator(detectorHost string, detectorPort int) Operator {
	return &operator{
		detectorConfig: &detectorConfig{
			host: detectorHost,
			port: detectorPort,
		},
	}
}

func (o operator) Query(expr string, timestamp int64) (*prometheus.MetricData, error) {
	url := fmt.Sprintf("http://%s:%d/api/v1/prometheus/query?query=%s&time=%d", o.detectorConfig.host, o.detectorConfig.port, expr, timestamp)
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
		return nil, errors.New("cannot get data from detector")
	}

	var m prometheus.MetricData
	dataBytes, err := jsoniter.Marshal(response["data"])
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(dataBytes, &m); err != nil {
		return nil, err
	}

	return &m, nil
}

func (o operator) QueryRange(expr string, startTime, endTime int64, step int64) (*prometheus.MetricData, error) {
	url := fmt.Sprintf("http://%s:%d/api/v1/prometheus/query_range?query=%s&start_time=%d&end_time=%d&step=%d",
		o.detectorConfig.host,
		o.detectorConfig.port,
		expr,
		startTime,
		endTime,
		step,
	)
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
		return nil, errors.New("cannot get data from detector")
	}

	var m prometheus.MetricData
	dataBytes, err := jsoniter.Marshal(response["data"])
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(dataBytes, &m); err != nil {
		return nil, err
	}

	return &m, nil
}
