package rules

import (
	"cpds/cpds-analyzer/internal/pkg/detector"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Operator interface {
	GetRules(filter, sortField, sortOrder string, pageNo, pageSize int) ([]Rule, error)

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

func (o *operator) GetRules(filter, sortField, sortOrder string, pageNo, pageSize int) ([]Rule, error) {
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
	return rules, nil
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
