/* 
 *  Copyright 2023 CPDS Author
 *  
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *  
 *       https://www.apache.org/licenses/LICENSE-2.0
 *  
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

package monitor

import (
	"cpds/cpds-analyzer/internal/pkg/detector/monitor"
	"cpds/cpds-analyzer/pkg/prometheus"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	jsoniter "github.com/json-iterator/go"
)

type Operator interface {
	GetMonitorTargets() (*MonitorTargets, error)

	GetNodeInfo(instance string) ([]NodeInfo, error)

	GetNodeStatus(instance string) ([]NodeStatus, error)

	GetNodeResources(instance string, startTime time.Time, endTime time.Time, step int64) ([]prometheus.Metric, error)

	GetNodeContainerStatus(instance string) ([]prometheus.Metric, error)

	GetClusterResource(startTime time.Time, endTime time.Time, step int64) ([]prometheus.Metric, error)

	GetClusterContainerStatus(startTime time.Time, endTime time.Time, step int64) ([]prometheus.Metric, error)
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

func (o *operator) GetMonitorTargets() (*MonitorTargets, error) {
	url := fmt.Sprintf("http://%s:%d/api/v1/monitor/targets", o.detectorConfig.host, o.detectorConfig.port)
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

	var mt MonitorTargets
	dataBytes, err := jsoniter.Marshal(response["data"])
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(dataBytes, &mt); err != nil {
		return nil, err
	}

	return &mt, nil
}

func (o *operator) GetNodeInfo(instance string) ([]NodeInfo, error) {
	var url string
	if instance == "" {
		url = fmt.Sprintf("http://%s:%d/api/v1/monitor/node_info", o.detectorConfig.host, o.detectorConfig.port)
	} else {
		url = fmt.Sprintf("http://%s:%d/api/v1/monitor/node_info?instance=%s", o.detectorConfig.host, o.detectorConfig.port, instance)
	}

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

	var n []NodeInfo
	for _, data := range response["data"].([]interface{}) {
		var ni NodeInfo
		dataBytes, err := jsoniter.Marshal(data)
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal(dataBytes, &ni); err != nil {
			return nil, err
		}
		n = append(n, ni)
	}

	return n, nil
}

func (o *operator) GetNodeStatus(instance string) ([]NodeStatus, error) {
	var url string
	if instance == "" {
		url = fmt.Sprintf("http://%s:%d/api/v1/monitor/node_status", o.detectorConfig.host, o.detectorConfig.port)
	} else {
		url = fmt.Sprintf("http://%s:%d/api/v1/monitor/node_status?instance=%s", o.detectorConfig.host, o.detectorConfig.port, instance)
	}

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

	var n []NodeStatus
	for _, data := range response["data"].([]interface{}) {
		var ns NodeStatus
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

func (o *operator) GetNodeResources(instance string, startTime time.Time, endTime time.Time, step int64) ([]prometheus.Metric, error) {
	url := fmt.Sprintf(
		"http://%s:%d/api/v1/monitor/node_resources?instance=%s&start_time=%s&end_time=%s&step=%d",
		o.detectorConfig.host,
		o.detectorConfig.port,
		instance,
		strconv.FormatInt(startTime.Unix(), 10),
		strconv.FormatInt(endTime.Unix(), 10),
		step,
	)
	metrics, err := monitor.GetMonitorDataFromDetector(url)
	if err != nil {
		return nil, err
	}

	return metrics, nil
}

func (o *operator) GetNodeContainerStatus(instance string) ([]prometheus.Metric, error) {
	url := fmt.Sprintf("http://%s:%d/api/v1/monitor/node_container_status?instance=%s", o.detectorConfig.host, o.detectorConfig.port, instance)
	metrics, err := monitor.GetMonitorDataFromDetector(url)
	if err != nil {
		return nil, err
	}

	return metrics, nil
}

func (o *operator) GetClusterResource(startTime time.Time, endTime time.Time, step int64) ([]prometheus.Metric, error) {
	url := fmt.Sprintf(
		"http://%s:%d/api/v1/monitor/cluster_resources?start_time=%s&end_time=%s&step=%d",
		o.detectorConfig.host,
		o.detectorConfig.port,
		strconv.FormatInt(startTime.Unix(), 10),
		strconv.FormatInt(endTime.Unix(), 10),
		step,
	)
	metrics, err := monitor.GetMonitorDataFromDetector(url)
	if err != nil {
		return nil, err
	}

	return metrics, nil
}

func (o *operator) GetClusterContainerStatus(startTime time.Time, endTime time.Time, step int64) ([]prometheus.Metric, error) {
	url := fmt.Sprintf(
		"http://%s:%d/api/v1/monitor/cluster_container_status?start_time=%s&end_time=%s&step=%d",
		o.detectorConfig.host,
		o.detectorConfig.port,
		strconv.FormatInt(startTime.Unix(), 10),
		strconv.FormatInt(endTime.Unix(), 10),
		step,
	)
	metrics, err := monitor.GetMonitorDataFromDetector(url)
	if err != nil {
		return nil, err
	}

	return metrics, nil
}
