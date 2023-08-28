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

import "cpds/cpds-analyzer/internal/models/rules"

type getOptions struct {
	filter    string
	sortField string
	sortOrder string
	pageNo    int
	pageSize  int
}

type getResponse struct {
	Records   []rules.Rules `json:"records"`
	PageTotal int          `json:"page_total"`
	PageNo    int          `json:"page_no"`
	PageSize  int          `json:"page_size"`
}

type createRequest struct {
	*rules.Rule
}

type updateRequest struct {
	*rules.Rule
}

type ruleRequest struct {
	*rules.Rules
}

type deleteRequest struct {
	ID int `json:"id"`
}
