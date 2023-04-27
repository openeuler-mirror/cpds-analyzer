package monitor

type clusterMonitorQueryParams struct {
	StartTime  int64 `json:"start_time"`
	EndTime    int64 `json:"end_time"`
	StepSecond int64 `json:"step"`
}

type nodeMonitorDataQueryParams struct {
	Instance   string `json:"instance"`
	StartTime  int64  `json:"start_time"`
	EndTime    int64  `json:"end_time"`
	StepSecond int64  `json:"step"`
}
