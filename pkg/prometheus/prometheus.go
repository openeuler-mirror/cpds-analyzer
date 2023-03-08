package prometheus

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/prometheus/client_golang/api"
	apiv1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
)

const MeteringDefaultTimeout = 20 * time.Second

// prometheus implements monitoring interface backed by Prometheus
type prometheus struct {
	client apiv1.API
}

func NewPrometheus(host string, port int) (prometheus, error) {
	cfg := api.Config{
		Address: fmt.Sprintf("%s/%d", host, port),
	}

	client, err := api.NewClient(cfg)
	return prometheus{client: apiv1.NewAPI(client)}, err
}

func (p prometheus) GetSingleMetric(expr string, ts time.Time) Metric {
	var parsedResp Metric

	value, _, err := p.client.Query(context.Background(), expr, ts)
	if err != nil {
		parsedResp.Error = err.Error()
	} else {
		parsedResp.MetricData = parseQueryResp(value, nil)
	}

	return parsedResp
}

func (p prometheus) GetSingleMetricOverTime(expr string, start, end time.Time, step time.Duration) Metric {
	timeRange := apiv1.Range{
		Start: start,
		End:   end,
		Step:  step,
	}

	value, _, err := p.client.QueryRange(context.Background(), expr, timeRange)

	var parsedResp Metric
	if err != nil {
		parsedResp.Error = err.Error()
	} else {
		parsedResp.MetricData = parseQueryRangeResp(value, nil)
	}
	return parsedResp
}

func (p prometheus) GetMultiMetrics(metrics []string, ts time.Time) []Metric {
	var res []Metric
	var mtx sync.Mutex
	var wg sync.WaitGroup

	for _, metric := range metrics {
		wg.Add(1)
		go func(metric string) {
			parsedResp := Metric{MetricName: metric}

			value, _, err := p.client.Query(context.Background(), metric, ts)
			if err != nil {
				parsedResp.Error = err.Error()
			} else {
				parsedResp.MetricData = parseQueryResp(value, nil)
			}

			mtx.Lock()
			res = append(res, parsedResp)
			mtx.Unlock()

			wg.Done()
		}(metric)
	}

	wg.Wait()

	return res
}

func (p prometheus) GetMultiMetricsOverTime(metrics []string, start, end time.Time, step time.Duration) []Metric {
	var res []Metric
	var mtx sync.Mutex
	var wg sync.WaitGroup

	timeRange := apiv1.Range{
		Start: start,
		End:   end,
		Step:  step,
	}

	for _, metric := range metrics {
		wg.Add(1)
		go func(metric string) {
			parsedResp := Metric{MetricName: metric}

			value, _, err := p.client.QueryRange(context.Background(), metric, timeRange)
			if err != nil {
				parsedResp.Error = err.Error()
			} else {
				parsedResp.MetricData = parseQueryRangeResp(value, nil)
			}

			mtx.Lock()
			res = append(res, parsedResp)
			mtx.Unlock()

			wg.Done()
		}(metric)
	}

	wg.Wait()

	return res
}

func parseQueryRangeResp(value model.Value, metricFilter func(metric model.Metric) bool) MetricData {
	res := MetricData{MetricType: MetricTypeMatrix}

	data, _ := value.(model.Matrix)

	for _, v := range data {
		if metricFilter != nil && !metricFilter(v.Metric) {
			continue
		}
		mv := MetricValue{
			Metadata: make(map[string]string),
		}

		for k, v := range v.Metric {
			mv.Metadata[string(k)] = string(v)
		}

		for _, k := range v.Values {
			mv.Series = append(mv.Series, Point{float64(k.Timestamp) / 1000, float64(k.Value)})
		}

		res.MetricValues = append(res.MetricValues, mv)
	}

	return res
}

func parseQueryResp(value model.Value, metricFilter func(metric model.Metric) bool) MetricData {
	res := MetricData{MetricType: MetricTypeVector}

	data, _ := value.(model.Vector)

	for _, v := range data {
		if metricFilter != nil && !metricFilter(v.Metric) {
			continue
		}
		mv := MetricValue{
			Metadata: make(map[string]string),
		}

		for k, v := range v.Metric {
			mv.Metadata[string(k)] = string(v)
		}

		mv.Sample = &Point{float64(v.Timestamp) / 1000, float64(v.Value)}

		res.MetricValues = append(res.MetricValues, mv)
	}

	return res
}
