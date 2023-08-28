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

package analysis

import (
	"cpds/cpds-analyzer/internal/models/rules"
	"cpds/cpds-analyzer/pkg/prometheus"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	jsoniter "github.com/json-iterator/go"
	"gorm.io/gorm"
)

type Operator interface {
	GetAnalysisResult(filter, sortField, sortOrder string, pageNo, pageSize int) ([]Analysis, error)

	DeleteAnalysisResultByID(ID uint) error

	GetRawData(ID uint) (*prometheus.Metric, error)

	GetTotalPages(filter string) int
}

type operator struct {
	db             *gorm.DB
	detectorConfig *detectorConf
}

type detectorConf struct {
	host string
	port int
}

func NewOperator(detedetectorHost string, detectorPort int, db *gorm.DB) Operator {
	return &operator{
		db: db.Session(&gorm.Session{}),
		detectorConfig: &detectorConf{
			host: detedetectorHost,
			port: detectorPort,
		},
	}
}

func (o *operator) GetAnalysisResult(filter, sortField, sortOrder string, pageNo, pageSize int) ([]Analysis, error) {
	var query = o.db.Model(&Analysis{})
	if filter != "" {
		query = query.Where("rule_name LIKE ?", "%"+filter+"%")
	}

	query = query.Order(fmt.Sprintf("%s %s", sortField, sortOrder))

	offset := (pageNo - 1) * pageSize
	query = query.Offset(offset).Limit(pageSize)

	var analysis []Analysis
	err := query.Find(&analysis).Error
	if err != nil {
		return nil, err
	}
	return analysis, nil
}

func (o *operator) DeleteAnalysisResultByID(ID uint) error {
	result := o.db.Delete(&Analysis{}, ID)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (o *operator) GetRawData(analysisID uint) (*prometheus.Metric, error) {
	var analysis *Analysis
	result := o.db.Model(&Analysis{}).Where("id = ?", analysisID).Find(&analysis)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}

	var rule *rules.Rule
	result = o.db.Model(&rules.Rule{}).Where("id = ?", analysis.RuleID).Find(&rule)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}

	step := (analysis.UpdateTime - analysis.CreateTime) / 250 // format step
	if step == 0 {
		step = 1
	}

	urlStr := fmt.Sprintf(
		"http://%s:%d/api/v1/prometheus/query_range?query=%s&start_time=%d&end_time=%d&step=%d",
		o.detectorConfig.host,
		o.detectorConfig.port,
		url.QueryEscape(rule.Expression),
		analysis.CreateTime,
		analysis.UpdateTime,
		step,
	)

	resp, err := http.Get(urlStr)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data *detectorRawDataResponse
	if err := jsoniter.NewDecoder(resp.Body).Decode(&data); err != nil || data.Status != 200 {
		return nil, errors.New("cannot get raw data from detector")
	}

	return data.Data, err
}

func (o *operator) GetTotalPages(filter string) int {
	var tableCount int64
	var query = o.db
	query = query.Model(&Analysis{}).Where("rule_name LIKE ?", "%"+filter+"%").Count(&tableCount)
	return int(tableCount)
}
