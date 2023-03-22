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
	Records   []rules.Rule `json:"records"`
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

type deleteRequest struct {
	ID int `json:"id"`
}
