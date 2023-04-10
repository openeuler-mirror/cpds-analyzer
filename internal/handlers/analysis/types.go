package analysis

import (
	"cpds/cpds-analyzer/internal/models/analysis"
	"cpds/cpds-analyzer/pkg/prometheus"
)

type getResultOptions struct {
	filter    string
	sortField string
	sortOrder string
	pageNo    int
	pageSize  int
}

type getResultResponse struct {
	Records   []analysis.Analysis `json:"records"`
	PageTotal int                 `json:"page_total"`
	PageNo    int                 `json:"page_no"`
	PageSize  int                 `json:"page_size"`
}

type deleteResultRequest struct {
	ID uint `json:"id"`
}

type getRawDataRequset struct {
	ID uint `json:"id"`
}

type getRawDataResponse struct {
	Records *prometheus.Metric `json:"records"`
}
