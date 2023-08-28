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

package rules

import (
	"cpds/cpds-analyzer/internal/pkg/detector"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type Operator interface {
	GetRules(filter, sortField, sortOrder string, pageNo, pageSize int) ([]Rules, error)

	CreateRule(rule *Rule) error

	UpdateRule(rule *Rule) error

	SendRuleUpdatedRequset() error

	DeleteRuleByID(id int) error

	GetTotalPages(filter string) int
}

type operator struct {
	detectorConfig *detectorConfig
	db             *gorm.DB
}

type detectorConfig struct {
	host string
	port int
}

func NewOperator(detectorHost string, detectorPort int, db *gorm.DB) Operator {
	return &operator{
		db: db.Session(&gorm.Session{}),
		detectorConfig: &detectorConfig{
			host: detectorHost,
			port: detectorPort,
		},
	}
}

func (o *operator) GetRules(filter, sortField, sortOrder string, pageNo, pageSize int) ([]Rules, error) {
	var query = o.db
	if filter != "" {
		query = query.Where("name LIKE ?", "%"+filter+"%")
	}

	query = query.Order(fmt.Sprintf("%s %s", sortField, sortOrder))

	offset := (pageNo - 1) * pageSize
	query = query.Offset(offset).Limit(pageSize)

	var rules []Rule
	err := query.Find(&rules).Error
	if err != nil {
		return nil, err
	}

	ruleData := make([]Rules, 0)
	for _, rule := range rules{
		ruleData = append(ruleData, Rules{
			ID: rule.ID,
			Name: rule.Name,
			CreateTime: rule.CreateTime,
			UpdateTime: rule.UpdateTime,
			Duration: rule.Duration,
			Expression: rule.Expression,
			SubhealthConditionType: rule.SubhealthConditionType,
			SubhealthThresholds:strconv.FormatFloat(rule.SubhealthThresholds, 'f', -1, 64),
			FaultConditionType: rule.FaultConditionType,
			FaultThresholds: strconv.FormatFloat(rule.FaultThresholds, 'f', -1, 64),
			Severity: rule.Severity,
		})
	}
	return ruleData, nil
}

func (o *operator) CreateRule(rule *Rule) error {
	rule.CreateTime = time.Now().Unix()
	rule.UpdateTime = time.Now().Unix()

	if err := o.db.Create(rule).Error; err != nil {
		return err
	}

	return nil
}

func (o *operator) UpdateRule(rule *Rule) error {
	b,_ :=json.Marshal(rule)
	ruleMap :=make(map[string]interface{})
	json.Unmarshal(b, &ruleMap)
	result := o.db.Model(&rule).Updates(ruleMap)
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return errors.New("nothing changed")
	}
	o.db.Model(&rule).Update("update_time", time.Now().Unix())

	return nil
}

func (o *operator) DeleteRuleByID(id int) error {
	result := o.db.Delete(&Rule{}, id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (o *operator) SendRuleUpdatedRequset() error {
	if err := detector.SendRuleUpdatedRequset(o.detectorConfig.host, o.detectorConfig.port); err != nil {
		return err
	}

	return nil
}

func (o *operator) GetTotalPages(filter string) int {
	var tableCount int64
	var query = o.db
	query = query.Model(&Rule{}).Where("name LIKE ?", "%"+filter+"%").Count(&tableCount)
	return int(tableCount)
}
