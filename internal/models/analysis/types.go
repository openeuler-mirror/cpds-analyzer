package analysis

import "cpds/cpds-analyzer/pkg/prometheus"

type Analysis struct {
	ID         uint   `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	RuleID     uint   `json:"rule_id" gorm:"not null"`
	RuleName   string `json:"rule_name" gorm:"not null"`
	Status     string `json:"status" gorm:"not null"`
	Count      uint   `json:"count" gorm:"not null"`
	CreateTime int64  `json:"create_time" gorm:"not null"`
	UpdateTime int64  `json:"update_time" gorm:"not null"`
}

type detectorRawDataResponse struct {
	Status    int                `json:"status"`
	Code      int                `json:"code"`
	Message   string             `json:"message"`
	Data      *prometheus.Metric `json:"data"`
	Timestamp int64              `json:"timestamp"`
}
