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

type Rule struct {
	ID                     uint    `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	Name                   string  `json:"name" gorm:"unique;not null"`
	Expression             string  `json:"expression" gorm:"not null;type:varchar(512)"`
	SubhealthConditionType string  `json:"subhealth_condition_type"`
	SubhealthThresholds    float64 `json:"subhealth_thresholds"`
	FaultConditionType     string  `json:"fault_condition_type"`
	FaultThresholds        float64 `json:"fault_thresholds"`
	Severity               string  `json:"severity" gorm:"not null"`
	Duration               string  `json:"duration" gorm:"not null"`
	CreateTime             int64   `json:"create_time" gorm:"not null"`
	UpdateTime             int64   `json:"update_time" gorm:"not null"`
}

type Rules struct {
	ID                     uint    `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	Name                   string  `json:"name" gorm:"unique;not null"`
	Expression             string  `json:"expression" gorm:"not null;type:varchar(512)"`
	SubhealthConditionType string  `json:"subhealth_condition_type"`
	SubhealthThresholds    string `json:"subhealth_thresholds"`
	FaultConditionType     string  `json:"fault_condition_type"`
	FaultThresholds        string `json:"fault_thresholds"`
	Severity               string  `json:"severity" gorm:"not null"`
	Duration               string  `json:"duration" gorm:"not null"`
	CreateTime             int64   `json:"create_time" gorm:"not null"`
	UpdateTime             int64   `json:"update_time" gorm:"not null"`
}
